package linkedlist

import "fmt"

type Info struct {
	Size int
	Head *Element
	Tail *Element

	// 简单跳表索引
	SkipIndex *SimpleSkipIndex
}

// 初始化
func (i *Info) Init(element *Element) bool {
	if i.Head != nil {
		return false
	}
	i.Head = element
	i.Tail = element
	i.Size += 1
	i.SkipIndex = GetDefaultSkipIndex(element)
	return true
}

// 更新尾巴
func (i *Info) UpdateTail(element *Element) {
	i.Tail = element
}

// 更新头部
func (i *Info) UpdateHead(element *Element) {
	i.Head = element
}

// 节点数+1
func (i *Info) IncrementSize() {
	i.Size += 1
}

// 节点数-1
func (i *Info) DecrementSize() {
	i.Size += -1
}

// 重算索引(插入元素时)
func (i *Info) ReCalculateIndexByRandomInsert(insertIndex int, element *Element) {
	if insertIndex >= i.Size {
		panic("ReCalculateIndexWhenRandomInsert:index out!")
	}
	i.SkipIndex.RandomInsert(insertIndex, element, i.Size)
}

// 重算索引(删除元素时)
func (i *Info) ReCalculateIndexByRandomRemove(removeIndex int, next *Element) {
	if removeIndex >= i.Size+1 {
		panic(fmt.Sprintf("ReCalculateIndexWhenRandomInsert:index out! size: %v", removeIndex))
	}
	i.SkipIndex.RandomRemove(removeIndex, next)
}

// 匹配跳表索引
func (i *Info) MatchIndex(index int) (*Element, int) {
	node, leftoverRange := i.SkipIndex.MatchIndex(index)
	return node, leftoverRange
}

// 清空信息
func (i *Info) Clear() {
	i.Head = nil
	i.Tail = nil
	i.SkipIndex = nil
}
