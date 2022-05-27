package main

import "strings"

// 你需要设计一个文件系统，你可以创建新的路径并将它们与不同的值关联。
//
// 路径的格式是一个或多个连接在一起的字符串，形式为： / ，后面跟着一个或多个小写英文字母。例如， " /leetcode" 和 "/leetcode/
// problems" 是有效路径，而空字符串 "" 和 "/" 不是。
//
// 实现 FileSystem 类:
//
//
// bool createPath(string path, int value) 创建一个新的 path ，并在可能的情况下关联一个 value ，然后返回
// true 。如果路径已经存在或其父路径不存在，则返回 false 。
// int get(string path) 返回与 path 关联的值，如果路径不存在则返回 -1 。
//
//
//
//
// 示例 1：
//
//
// 输入：
// ["FileSystem","create","get"]
// [[],["/a",1],["/a"]]
// 输出：
// [null,true,1]
// 解释：
// FileSystem fileSystem = new FileSystem();
//
// fileSystem.create("/a", 1); // 返回 true
// fileSystem.get("/a"); // 返回 1
//
//
// 示例 2：
//
//
// 输入：
// ["FileSystem","createPath","createPath","get","createPath","get"]
// [[],["/leet",1],["/leet/code",2],["/leet/code"],["/c/d",1],["/c"]]
// 输出：
// [null,true,true,2,false,-1]
// 解释：
// FileSystem fileSystem = new FileSystem();
//
// fileSystem.createPath("/leet", 1); // 返回 true
// fileSystem.createPath("/leet/code", 2); // 返回 true
// fileSystem.get("/leet/code"); // 返回 2
// fileSystem.createPath("/c/d", 1); // 返回 false 因为父路径 "/c" 不存在。
// fileSystem.get("/c"); // 返回 -1 因为该路径不存在。
//
//
//
//
// 提示：
//
//
// 对两个函数的调用次数加起来小于等于 10⁴
// 2 <= path.length <= 100
// 1 <= value <= 10⁹
//
// Related Topics 设计 字典树 哈希表 字符串 👍 27 👎 0

// 2022-05-27 10:03:52
// leetcode submit region begin(Prohibit modification and deletion)
type FileSystem struct {
	path  map[string]*FileSystem
	value int
}

func Constructor() FileSystem {
	return FileSystem{
		path: make(map[string]*FileSystem),
	}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	if path == "" || path == "/" {
		return false
	}
	paths := strings.Split(path, "/")
	m := this
	for i := 1; i < len(paths)-1; i++ {
		if _, ok := m.path[paths[i]]; !ok {
			return false
		}
		m = m.path[paths[i]]
	}
	if _, ok := m.path[paths[len(paths)-1]]; ok {
		return false
	}
	m.path[paths[len(paths)-1]] = &FileSystem{
		path:  make(map[string]*FileSystem),
		value: value,
	}
	return true
}

func (this *FileSystem) Get(path string) int {
	if path == "" || path == "/" {
		return -1
	}
	paths := strings.Split(path, "/")
	m := this
	for i := 1; i < len(paths)-1; i++ {
		v, ok := m.path[paths[i]]
		if !ok {
			return -1
		}
		m = v
	}
	v, ok := m.path[paths[len(paths)-1]]
	if ok {
		return v.value
	}
	return -1
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */
// leetcode submit region end(Prohibit modification and deletion)

func main() {
	fs := Constructor()
	println(fs.CreatePath("/a", 1))
	println(fs.CreatePath("/a/b", 2))
	println(fs.Get("/a/b"))

}
