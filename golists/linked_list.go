package golists

import (
	"fmt"
	"github.com/tejchen/gocollection/golists/linkedlist"
)

/**
-----------------------------------
-        以下是List接口实现         -
-----------------------------------
*/

type LinkedListImpl struct {
	// 双向链表
	Head *linkedlist.Element
	// list 信息（头，尾，索引，总数）
	Info linkedlist.Info
}

func (l *LinkedListImpl) Add(item interface{}) {
	element := linkedlist.DefaultElement(item)
	if l.IsEmpty() {
		l.Head = element
		l.Info.Init(element)
		return
	}
	l.Info.Tail.AddNext(element)
	l.Info.IncrementSize()
	l.Info.UpdateTail(element)
	l.Info.ReCalculateIndexByRandomInsert(l.Info.Size-1, element)
}

func (l *LinkedListImpl) AddByIndex(idx int, item interface{}) {
	if idx < 0 || idx > l.Info.Size {
		panic(fmt.Sprintf("AppendByIndex:index out! size: %v, index: %v", l.Info.Size, idx))
	}
	// append
	if idx == l.Info.Size {
		l.Add(item)
		return
	}
	// find
	origin := l.get(idx)
	// insert
	element := linkedlist.DefaultElement(item)
	element.AddPrev(origin.GetPrev())
	element.AddNext(origin)
	// update info
	if element.GetPrev() == nil {
		l.updateHead(element)
	}
	l.Info.IncrementSize()
	l.Info.ReCalculateIndexByRandomInsert(idx, element)
}

func (l *LinkedListImpl) Get(idx int) interface{} {
	node := l.get(idx)
	if node != nil {
		return node.Value()
	}
	return nil
}

func (l *LinkedListImpl) Remove(idx int) bool {
	if idx < 0 || idx >= l.Info.Size {
		return false
	}
	target := l.get(idx)
	prev := target.GetPrev()
	next := target.GetNext()
	// head
	if prev == nil {
		// only head
		if next == nil {
			l.clear()
			return true
		} else {
			// 断开
			target.SetNext(nil)
			// 再初始化
			next.AddPrev(nil)
			l.updateHead(next)
		}
	}
	// tail
	if next == nil && prev != nil {
		target.SetPrev(nil)
		prev.AddNext(nil)
		l.Info.UpdateTail(prev)
	}
	// middle
	if prev != nil && next != nil {
		target.SetPrev(nil)
		target.SetNext(nil)
		prev.AddNext(next)
	}
	l.Info.DecrementSize()
	l.Info.ReCalculateIndexByRandomRemove(idx, next)
	return true
}

func (l *LinkedListImpl) Size() int {
	return l.Info.Size
}

func (l *LinkedListImpl) IsEmpty() bool {
	return l.Head == nil
}

func (l *LinkedListImpl) NotEmpty() bool {
	return !l.IsEmpty()
}

func (l *LinkedListImpl) Foreach(iterator func(item interface{})) {
	temp := l.Head
	for {
		if temp == nil {
			break
		}
		iterator(temp.Value())
		temp = temp.GetNext()
		if temp == nil {
			break
		}
	}
}

/**
-----------------------------------
-          以下是内部方法           -
-----------------------------------
*/
func (l *LinkedListImpl) updateHead(element *linkedlist.Element) {
	l.Head = element
	l.Info.UpdateHead(element)
}

func (l *LinkedListImpl) clear() {
	l.Head = nil
	l.Info.Clear()
	return
}

func (l *LinkedListImpl) get(idx int) *linkedlist.Element {
	if idx < 0 || idx >= l.Info.Size {
		panic(fmt.Sprintf("Get:index out! size:%v, index:%v", l.Info.Size, idx))
	}
	// skip table search
	indexNode, leftoverRange := l.Info.MatchIndex(idx)
	var node = indexNode
	for i := 0; i < leftoverRange; i++ {
		node = node.GetNext()
	}
	return node
}
