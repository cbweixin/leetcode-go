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
	for {
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
	return df, nil
}

func main() {

}
