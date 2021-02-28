package current_map

// ConcurrentMap 代表并发安全的字典的接口。
type CouncurrentMap interface {
	// Concurrency 会返回并发量。
	Concurrency() int
	// Put 会推送一个键-元素对。
	// 注意！参数element的值不能为nil。
	// 第一个返回值表示是否新增了键-元素对。
	// 若键已存在，新元素值会替换旧的元素值。
	Put(key string, element interface{}) (bool, error)
	// Get 会获取与指定键关联的那个元素。
	// 若返回nil，则说明指定的键不存在。
	Get(key string) interface{}
	// Delete 会删除指定的键-元素对。
	// 若结果值为true则说明键已存在且已删除，否则说明键不存在。
	Delete(key string) bool
	// Len 会返回当前字典中键-元素对的数量。
	Len() uint64
}
