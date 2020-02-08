package collection

import "fmt"

type ArrayList struct {
	data []interface{}
}

func (sa *ArrayList) Add(item interface{}) {
	sa.data = append(sa.data, item)
}

func (sa *ArrayList) AddByIndex(idx int, item interface{}) {
	sa.indexRequired(idx)
	temp1 := sa.data[:idx]
	temp2 := sa.data[idx:]
	sa.data = append(append(temp1, item), temp2...)
}

// 可以考虑用copy的方式
func (sa *ArrayList) Remove(idx int) bool {
	if !sa.checkIndex(idx) {
		return false
	}
	if idx == 0 {
		sa.data = sa.data[1:]
		return true
	}
	if idx+1 == sa.Size() {
		sa.data = sa.data[:sa.Size()-1]
		return true
	}
	sa.data = append(sa.data[:idx], sa.data[idx+1:]...)
	return true
}

func (sa *ArrayList) Size() int {
	return len(sa.data)
}

func (sa *ArrayList) Get(idx int) interface{} {
	sa.indexRequired(idx)
	return (sa.data)[idx]
}

func (sa *ArrayList) Foreach(iterator func(item interface{})) {
	for i := 0; i < sa.Size(); i++ {
		iterator((sa.data)[i])
	}
}

func (sa *ArrayList) IsEmpty() bool {
	return len(sa.data) == 0
}

func (sa *ArrayList) IsNotEmpty() bool {
	return len(sa.data) != 0
}

func (sa *ArrayList) indexRequired(idx int) {
	if !sa.checkIndex(idx) {
		panic(fmt.Sprintf("index out!len:%v, index: %v", len(sa.data), idx))
	}
}

func (sa *ArrayList) checkIndex(idx int) bool {
	if idx >= len(sa.data) || idx < 0 {
		return false
	}
	return true
}
