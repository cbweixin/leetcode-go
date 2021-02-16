package read_write_lock_v2

import (
	"errors"
	"io"
	"os"
	"sync"
)

// Data 代表数据的类型。
type Data []byte

// DataFile 代表数据文件的接口类型。
type DataFile interface {
	// Read 会读取一个数据块。
	Read() (rsn int64, d Data, err error)
	// Write 会写入一个数据块。
	Write(d Data) (wsn int64, err error)
	// RSN 会获取最后读取的数据块的序列号。
	RSN() int64
	// WSN 会获取最后写入的数据块的序列号。
	WSN() int64
	// DataLen 会获取数据块的长度。
	DataLen() uint32
	// Close 会关闭数据文件。
	Close() error
}

// myDataFile 代表数据文件的实现类型。
type myDataFile struct {
	f       *os.File     // 文件。
	fmutex  sync.RWMutex // 被用于文件的读写锁。
	rcond   *sync.Cond   //读操作需要用到的条件变量
	woffset int64        // 写操作需要用到的偏移量。
	roffset int64        // 读操作需要用到的偏移量。
	wmutex  sync.Mutex   // 写操作需要用到的互斥锁。
	rmutex  sync.Mutex   // 读操作需要用到的互斥锁。
	dataLen uint32       // 数据块长度。
}

// NewDataFile 会新建一个数据文件的实例。
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	// 我们一直在说，条件变量rcond是与读写锁fmutex的“读锁”关联的。这是怎样做到的呢？读者还记得我们在上一节提到读写锁的RLocker方法吗？
	// 它会返回当前读写锁中的“读锁”。这个结果值同时也是sync.Locker接口的实现。因此，我们可以把它作为参数值传给sync.NewCond函数。
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量。
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	//读取一个数据块。
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()
	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				// 添加这条语句的意义在于：当发现由文件内容读取造成的EOF错误时，要让当前Goroutine暂时放弃fmutex的“读锁”并等待通知的到来。
				// 放弃fmutex的“读锁”也就意味着Write方法中的数据块写操作不会受到它的阻碍了。在写操作完成之后，我们应该及时向条件
				// 变量rcond发送通知以唤醒为此而等待的Goroutine。请注意，在某个Goroutine被唤醒之后，应该再次检查需要被满足的条件。
				// 在这里，这个需要被满足的条件是在进行文件内容读取时不会造成EOF错误。如果该条件被满足，那么就可以进行后续的操作了。
				// 否则，应该再次放弃“读锁”并等待通知。这也是我们依然保留for循环的原因。

				// 这里有两点需要特别注意。
				//
				//一定要在调用rcond的Wait方法之前锁定与之关联的那个“读锁”，否则就会造成对Wait方法的调用永远无法返回。这种情况会导致流程
				// 执行的停滞，甚至整个程序的死锁！导致这种结果的原因与条件变量和读写锁的内部实现方式有关（结果也许并不应该是这样，
				// 作者已经向Go语言官方提交了一个issue；Go语言官方已经接受了这个issue，并承诺将会在Go 1.4版本中改进它）。假设，
				// 与条件变量rcond关联的是某个读写锁的“写锁”或普通的互斥锁，那么对rcond.Wait方法的调用将会引发一个运行时恐慌。原因是，
				// 该方法会先对与之关联的锁进行解锁，而试图解锁未被锁定的锁就会引发一个运行时恐慌。
				//一定不要忘记在读操作完成之前解锁与条件变量rcond关联的那个“读锁”，否则对读写锁的写锁定操作将会阻塞相关的Goroutine。
				// 其根本原因是，条件变量rcond的Wait方法在返回之前会重新锁定与之关联的那个“读锁”。因此，在结束这个从文件中读取一个
				// 数据块的流程之前，我们应该调用fmutex字段的RLock方法。那条defer语句就起到了这个作用。
				df.rcond.Wait()
				continue
			}
			return
		}
		d = bytes
		return
	}
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	// 读取并更新写偏移量。
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	//写入一个数据块。
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	// 与Wait方法不同，我们在调用条件变量的Signal方法和Broadcast方法之前无需锁定与之关联的锁。随之，相应的解锁操作也是不需要的。在这个
	// Write方法中的锁定操作和解锁的操作针对的并不是df.rcond.Signal()语句。
	df.rcond.Signal()
	return
}

func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *myDataFile) Close() error {
	if df.f == nil {
		return nil
	}
	return df.f.Close()
}

func main() {

}
