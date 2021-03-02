package current_map

import "unsafe"

// linkedPair 代表单向链接的键-元素对的接口。
type linkedPair interface {
	// Next 用于获得下一个键-元素对。
	// 若返回值为nil，则说明当前已在单链表的末尾。
	Next() Pair
	// SetNext 用于设置下一个键-元素对。
	// 这样就可以形成一个键-元素对的单链表。
	SetNext(nextPair Pair) error
}

// Pair 代表并发安全的键-元素对的接口
type Pair interface {
	// linkedPair 代表单链键-元素对接口。
	linkedPair
	// Key 会返回键的值。
	Key() string
	// Hash 会返回键的哈希值。
	Hash() uint64
	// Element 会返回元素的值。
	Element() interface{}
	// Set 会设置元素的值。
	SetElement(element interface{}) error
	// Copy 会生成一个当前键-元素对的副本并返回。
	Copy() Pair
	// String 会返回当前键-元素对的字符串表示形式。
	String() string
}

// pair 代表键-元素对的类型。
type pair struct {
	key string
	// hash 代表键的哈希值。
	hash    uint64
	element unsafe.Pointer
	next    unsafe.Pointer
}

func newPair(key string, element interface{}) (Pair, error) {
	p := &pair{
		key:  key,
		hash: hash(key),
	}

	if element == nil {
		return nil, newIllegalParameterError("element is nil")
	}
	return p, nil
}

func (p *pair) Key() string {
	return p.key
}

func (p *pair) Hash() uint64 {
	return p.hash
}
