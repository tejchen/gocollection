package linkedlist

import "fmt"

var IndexRange = 100

type SimpleSkipIndex []*Element

func GetDefaultSkipIndex(headNode *Element) *SimpleSkipIndex {
	return &SimpleSkipIndex{headNode}
}

// 随机插入节点，该方法会自动调整索引
func (si *SimpleSkipIndex) RandomInsert(insertIdx int, element *Element, newSize int) {
	if insertIdx >= newSize {
		panic("RandomInsert:index out of index!")
	}
	// 计算索引位置
	skipIndexIdx := si.calculateNodeIndex(insertIdx)
	// 插入在索引位置
	if insertIdx%IndexRange == 0 {
		// 新索引节点
		if skipIndexIdx == si.Len() {
			*si = append(*si, element)
			return
		}
		// 替换索引节点
		(*si)[skipIndexIdx] = element
	}
	// 索引左移
	si.calculateOffset(skipIndexIdx, true)
	// 补索引节点
	si.appendIndex(newSize)
}

// 随机删除节点，该方法会自动调整索引
func (si *SimpleSkipIndex) RandomRemove(removeIdx int, next *Element) {
	// 计算索引位置
	skipIndexIdx := si.calculateNodeIndex(removeIdx)
	// 删除的是索引节点
	if removeIdx%IndexRange == 0 {
		// 尾部
		if skipIndexIdx == si.Len()-1 {
			if next != nil {
				(*si)[skipIndexIdx] = next
				return
			}
			*si = append((*si)[:si.Len()-1])
			return
		}
		// 替换索引节点
		(*si)[skipIndexIdx] = next
	}
	// 索引右移
	si.calculateOffset(skipIndexIdx, false)
}

// 匹配索引
func (si *SimpleSkipIndex) MatchIndex(index int) (*Element, int) {
	nodeIndex := index / IndexRange
	leftoverRange := index % IndexRange
	if nodeIndex >= si.Len() {
		panic(fmt.Sprintf("MatchIndex:index out of index! size:%v, index length: %v", index, si.Len()))
	}
	return (*si)[nodeIndex], leftoverRange
}

// 索引长度
func (si *SimpleSkipIndex) Len() int {
	return len(*si)
}

// 计算索引位置
func (si *SimpleSkipIndex) calculateNodeIndex(index int) int {
	nodeIndex := index / IndexRange
	return nodeIndex
}

// 计算索引位置
func (si *SimpleSkipIndex) appendIndex(newSize int) {
	if newSize == IndexRange*si.Len()+1 {
		e := (*si)[si.Len()-1]
		for i := 0; i < IndexRange; i++ {
			e = e.GetNext()
		}
		*si = append(*si, e)
	}
}

// 索引偏移
// incr 为 true ，右边索引左移一位
// incr 为 false ，右边索引右移一位
func (si *SimpleSkipIndex) calculateOffset(indexIdx int, incr bool) {
	if indexIdx >= si.Len() {
		panic("CalculateOffset:invalid indexIdx!")
	}
	// 左偏移
	if incr {
		for i := indexIdx + 1; i < si.Len(); i++ {
			(*si)[i] = (*si)[i].GetPrev()
		}
	}
	// 右偏移
	if !incr {
		for i := indexIdx + 1; i < si.Len(); i++ {
			// 删除该索引节点
			if (*si)[i].GetNext() == nil {
				*si = append((*si)[:si.Len()-1])
				continue
			}
			(*si)[i] = (*si)[i].GetNext()
		}
	}
}
