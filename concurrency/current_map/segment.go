package current_map

import "sync"

// Segment 代表并发安全的散列段的接口。
type Segment interface {
	// Put 会根据参数放入一个键-元素对。
	// 第一个返回值表示是否新增了键-元素对。
	Put(p Pair) (bool, error)
	// Get 会根据给定参数返回对应的键-元素对。
	// 该方法会根据给定的键计算哈希值。
	Get(key string) Pair
	// GetWithHash 会根据给定参数返回对应的键-元素对。
	// 注意！参数keyHash应该是基于参数key计算得出哈希值。
	GetWithHash(key string, keyHash uint64) Pair
	// Delete 会删除指定键的键-元素对。
	// 若返回值为true则说明已删除，否则说明未找到该键。
	Delete(key string) bool
	// Size 用于获取当前段的尺寸（其中包含的散列桶的数量）。
	Size() uint64
}

// segment 代表并发安全的散列段的类型。
type segment struct {
	buckets           []Bucket
	bucketsLen        int
	pairTotal         uint64
	pairRedistributor PairRedistributor
	lock              sync.Mutex
}
