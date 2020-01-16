package goarrays

import "fmt"

type SimpleArray struct {
	data []interface{}
}

func (sa *SimpleArray) Append(item interface{}) {
	sa.data = append(sa.data, item)
}

func (sa *SimpleArray) Interface() []interface{} {
	return sa.data
}

func (sa *SimpleArray) AppendAll(items Array) {
	sa.data = append(sa.data, items.Interface()...)
}

func (sa *SimpleArray) Replace(idx int, item interface{}) {
	sa.checkIndex(idx)
	(sa.data)[idx] = item
}

func (sa *SimpleArray) Remove(idx int) {
	sa.checkIndex(idx)
	if idx == 0 {
		sa.data = append((sa.data)[1:])
		return
	}
	if idx+1 == sa.Len() {
		sa.data = append((sa.data)[:sa.Len()-1])
		return
	}
	sa.data = append(sa.data[:idx-1], sa.data[idx:]...)
}

func (sa *SimpleArray) Len() int {
	return len(sa.data)
}

func (sa *SimpleArray) Get(idx int) interface{} {
	sa.checkIndex(idx)
	return (sa.data)[idx]
}

func (sa *SimpleArray) Foreach(iterator func(interface{})) {
	for i := 0; i < sa.Len(); i++ {
		iterator((sa.data)[i])
	}
}

func (sa *SimpleArray) IsEmpty() bool {
	return len(sa.data) == 0
}

func (sa *SimpleArray) IsNotEmpty() bool {
	return len(sa.data) != 0
}

func (sa *SimpleArray) checkIndex(idx int) {
	if idx >= len(sa.data) || idx < 0 {
		panic(fmt.Sprintf("index out!len:%v, index: %v", len(sa.data), idx))
	}
}
