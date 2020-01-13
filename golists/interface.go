package golists

type List interface {
	// 追加
	Add(item interface{})
	// 按下标插入
	AddByIndex(idx int, item interface{})
	// 获取
	Get(idx int) interface{}
	// 删除
	Remove(idx int) bool
	// 总数
	Size() int
	// 判空
	IsEmpty() bool
	// 判非空
	NotEmpty() bool
	// 迭代方法
	Foreach(iterator func(item interface{}))
}
