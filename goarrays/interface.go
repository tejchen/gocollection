package goarrays

/**
封装数组操作，操作语义化
*/
type Array interface {

	// 追加元素
	Append(item interface{})

	// 批量追加元素
	AppendAll(items Array)

	// 替换元素
	Replace(idx int, item interface{})

	// 按下标删除
	Remove(idx int)

	// 获取长度
	Len() int

	// 下标查找
	Get(index int) interface{}

	// 迭代方法
	Foreach(func(item interface{}))

	// 判空
	IsEmpty() bool

	// 判空
	IsNotEmpty() bool

	// 原数组
	Interface() []interface{}

	// todo 排序
	//Sort(func(interface{}, interface{}) int)

	// todo 切割
	//SubArray()
}
