package read_write_lock_v3

import (
	"errors"
	"os"
	"sync"
	"sync/atomic"
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

func (df *myDataFile) read() (rsn int64, d Data, err error) {
	var offset int64
	for {
		offset = atomic.LoadInt64(&df.roffset)
	}
}
