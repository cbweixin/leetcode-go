package main

import (
	"errors"
	"io"
	"os"
	"sync"
)

type Data []byte

// 而名称wsn和rsn分别是Writing Serial Number和Reading Serial Number的缩写形式。它们分别代表了最后被写入的数据块的序列号和最后被读取的
// 数据块的序列号。这里所说的序列号相当于一个计数值，它会从1开始。因此，我们可以通过调用Rsn方法和Wsn方法得到当前已被读取和写入的数据块的数量。
type DataFile interface {
	// Read would read a data block
	Read() (rsn int64, d Data, err error)
	// write a data block
	Write(d Data) (wsn int64, err error)
	// acquire the last read block
	RSN() int64
	// acquire the last wrote block
	WSN() int64
	// get the data block length
	DataLen() uint32
	// close data block
	Close() error
}

// 根据上面对需求的简单分析和这个DataFile接口类型声明，我们就可以来编写真正的实现了。我们将这个实现类型命名为myDataFile。
type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex // read-write lock
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
	// 把变量df的值作为NewDataFile函数的第一个结果值体现了我们的设计意图。但要想使*myDataFile类型真正成为DataFile类型的一个实现类型，
	// 我们还需要为*myDataFile类型编写出已在DataFile接口类型中声明的所有方法。其中最重要的当属Read方法和Write方法。
	df := &myDataFile{f: f, dataLen: dataLen}
	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读偏移量。
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	// 另一方面，在读取一个数据块的时候，我们适时的进行了fmutex字段的读锁定和读解锁操作。fmutex字段的这两个操作可以保证我们在这里读取
	// 到的是完整的数据块。不过，这个完整的数据块却并不一定是正确的。为什么会这样说呢？
	//
	//请想象这样一个场景。在我们的程序中，有3个Goroutine来并发的执行某个*myDataFile类型值的Read方法，并有2个Goroutine来并发的执行
	// 该值的Write方法。通过前3个Goroutine的运行，数据文件中的数据块被依次的读取了出来。但是，由于进行写操作的Goroutine比进行读操作
	// 的Goroutine少，所以过不了多久读偏移量roffset的值就会等于甚至大于写偏移量woffset的值。也就是说，读操作很快就会没有数据可读了。
	// 这种情况会使上面的df.f.ReadAt方法返回的第二个结果值为代表错误的非nil且会与io.EOF相等的值。实际上，我们不应该把这样的值看成错误
	// 的代表，而应该把它看成一种边界情况。但不幸的是，我们在这个版本的Read方法中并没有对这种边界情况做出正确的处理。该方法在遇到这种
	// 情况时会直接把错误值返回给它的调用方。该调用方会得到读取出错的数据块的序列号，但却无法再次尝试读取这个数据块。由于其他正在或后续
	// 执行的Read方法会继续增加读偏移量roffset的值，所以当该调用方再次调用这个Read方法的时候只可能读取到在此数据块后面的其他数据块。
	// 注意，执行Read方法时遇到上述情况的次数越多，被漏读的数据块也就会越多。

	/*	rsn = offset / int64(df.dataLen)
		df.fmutex.RLock()
		defer df.fmutex.RUnlock()
		bytes := make([]byte, df.dataLen)
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			return
		}
		d = bytes
		return

	*/

	// below is version 2, which is correct
	//读取一个数据块。

	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)

	for {
		// 第二个版本的Read方法使用for语句是为了达到这样一个目的：在其中的df.f.ReadAt方法返回io.EOF错误的时候继续尝试获取同一个数据块，
		// 直到获取成功为止。注意，如果在该for代码块被执行期间一直让读写锁fmutex处于读锁定状态，那么针对它的写锁定操作将永远不会成功，
		// 且相应的Goroutine也会被一直阻塞。因为它们是互斥的。所以，我们不得不在该for语句块中的每条return语句和continue语句的前面都加入
		// 一个针对该读写锁的读解锁操作，并在每次迭代开始时都对fmutex进行一次读锁定。显然，这样的代码看起来很丑陋。冗余的代码会使代码的维护
		// 成本和出错几率大大增加。并且，当for代码块中的代码引发了运行时恐慌的时候，我们是很难及时的对读写锁fmutex进行读解锁的。
		// 即便可以这样做，那也会使Read方法的实现更加丑陋。我们因为要处理一种边界情况而去掉了defer df.fmutex.RUnlock()语句。
		// 这种做法利弊参半。
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.fmutex.RUnlock()
				continue
			}
			df.fmutex.RUnlock()
			return
		}
		d = bytes
		df.fmutex.RUnlock()
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
	// 这里需要注意的是，当参数d的值的长度大于数据块的最大长度的时候，我们会先进行截短处理再将数据写入文件。如果没有这个截短处理，
	// 我们在后面计算的已读数据块的序列号和已写数据块的序列号就会不正确。
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
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
