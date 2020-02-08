package test

import (
	"github.com/tejchen/gocollection/golists"
	"github.com/tejchen/gocollection/golists/collection"
	"github.com/tejchen/gocollection/golists/collection/linkedlist"
	"github.com/tejchen/gocollection/test/base"
	"sort"
	"testing"
	"time"
)

//基础服务测试
func TestLinkedListBaseService(t *testing.T) {
	var result = true

	for i := 0; i < 100; i++ {
		// 基础测试
		list := golists.NewLinkedList()
		result = result && list.Size() == 0
		result = result && list.IsEmpty() == true
		result = result && list.IsNotEmpty() == false

		// 准备数据
		for i := 0; i < base.DefaultRange; i++ {
			list.Add(i)
		}

		result = result && linkedListTestBase(list.(*collection.LinkedList), true)

		// 顺序读测试
		for i := 0; i < base.DefaultRange; i++ {
			result = result && list.Get(i) == i
		}

		// 随机读测试 1/5
		randomTestSample := base.GetRandomTestSample(base.DefaultRange / 5)
		for i := 0; i < len(randomTestSample); i++ {
			result = result && randomTestSample[i] == list.Get(randomTestSample[i])
		}

		// 保证底层有序
		var temp = list.(*collection.LinkedList)
		node := temp.Head
		for i := 0; i < base.DefaultRange; i++ {
			result = result && node.Value() == i
			node = node.GetNext()
		}
		result = result && node == nil

		// 保证索引有序
		for i := 0; i < temp.Info.SkipIndex.Len(); i++ {
			skipIndex := temp.Info.SkipIndex.Idx[i]
			indexRange := temp.Info.SkipIndex.IndexRange
			result = result && (skipIndex.Value().(int)%indexRange == 0)
		}

		result = result && linkedListTestBase(list.(*collection.LinkedList), true)

		result = result && list.Size() == base.DefaultRange
		result = result && list.IsEmpty() == false
		result = result && list.IsNotEmpty() == true
	}

	t.Logf("linkedListTestBaseService: %v", result)
}

// 随机插入测试
func TestLinkedListRandomInsert(t *testing.T) {
	var randomInsertCheck = true

	var timeSum int64 = 0

	for i := 0; i < 100; i++ {
		// 容器
		randomInsertList := golists.NewLinkedList()

		// 随机样本数据
		randomTestSample := base.GetRandomTestSample(base.DefaultRange / 5)
		sort.Ints(randomTestSample)

		// 初始化数据
		linkedListInitRandomInsertContainer(randomInsertList, randomTestSample)

		start := time.Now().UnixNano()
		// 随机插入
		for i := 0; i < len(randomTestSample); i++ {
			randomInsertList.AddByIndex(randomTestSample[i], randomTestSample[i])
		}
		timeSum += time.Now().UnixNano() - start

		// 顺序读测试
		for i := 0; i < base.DefaultRange; i++ {
			randomInsertCheck = randomInsertCheck && randomInsertList.Get(i) == i
		}

		// 随机读测试 1/5
		newRandomTestSample := base.GetRandomTestSample(base.DefaultRange / 5)
		for i := 0; i < len(newRandomTestSample); i++ {
			randomInsertCheck = randomInsertCheck && newRandomTestSample[i] == randomInsertList.Get(newRandomTestSample[i])
		}

		// 保证底层有序
		var temp = randomInsertList.(*collection.LinkedList)
		node := temp.Head
		for i := 0; i < base.DefaultRange; i++ {
			randomInsertCheck = randomInsertCheck && node.Value() == i
			node = node.GetNext()
		}

		// 保证索引有序
		for i := 0; i < temp.Info.SkipIndex.Len(); i++ {
			skipIndex := temp.Info.SkipIndex.Idx[i]
			skipIndexRange := temp.Info.SkipIndex.IndexRange
			randomInsertCheck = randomInsertCheck && (skipIndex.Value().(int)%skipIndexRange == 0)
		}

		randomInsertCheck = randomInsertCheck && linkedListTestBase(randomInsertList.(*collection.LinkedList), true)
	}
	t.Logf("TestRandomInsert time: %v", timeSum/100)
	t.Logf("TestRandomInsert: %v", randomInsertCheck)
}

// 随机删除测试
func TestLinkedListRandomRemove(t *testing.T) {
	var randomRemoveCheck = true
	var timeSum int64 = 0

	for i := 0; i < 100; i++ {
		// 容器
		randomRemoveList := golists.NewLinkedList()

		// 随机样本数据 1/5
		randomTestSample := base.GetRandomTestSample(base.DefaultRange / 5)

		// 初始化数据
		for i := 0; i < base.DefaultRange; i++ {
			randomRemoveList.Add(i)
		}

		// 随机删除
		deleteCount := 0
		start := time.Now().UnixNano()
		for i := 0; i < len(randomTestSample); i++ {
			if randomRemoveList.Size() < randomTestSample[i]+1 {
				continue
			}
			randomRemoveList.Remove(randomTestSample[i])
			deleteCount++
		}
		timeSum += time.Now().UnixNano() - start

		// 长度测试
		randomRemoveCheck = randomRemoveCheck && randomRemoveList.Size() == base.DefaultRange-deleteCount

		// 保证底层有序
		var temp = randomRemoveList.(*collection.LinkedList)
		var node = temp.Head
		for i := 0; i < randomRemoveList.Size(); i++ {
			if i == 0 {
				continue
			}
			randomRemoveCheck = randomRemoveCheck && node.GetNext().Value().(int) > node.Value().(int)
		}

		// 保证索引节点对应
		node = temp.Head
		for i := 0; i < temp.Info.SkipIndex.Len(); i++ {
			if i%temp.Info.SkipIndex.IndexRange == 0 {
				skipIndex := temp.Info.SkipIndex.Idx[i/temp.Info.SkipIndex.IndexRange]
				randomRemoveCheck = randomRemoveCheck && skipIndex.Value() == node.Value()
			}
			node = node.GetNext()
		}
		randomRemoveCheck = randomRemoveCheck && linkedListTestBase(randomRemoveList.(*collection.LinkedList), true)
	}

	t.Logf("TestRandomRemove time: %v", timeSum/100)
	t.Logf("TestRandomRemove: %v", randomRemoveCheck)
}

// 定点边界测试
func TestLinkedListBoundary(t *testing.T) {
	result := true
	// 头插
	headInsert := golists.NewLinkedList()
	for i := 0; i < 300; i++ {
		headInsert.Add(i)
	}
	headInsert.AddByIndex(0, 0)
	result = result && headInsert.Size() == 301
	result = result && headInsert.Get(300) == 299
	result = result && headInsert.Get(0) == 0
	result = result && headInsert.Get(1) == 0
	result = result && headInsert.Get(2) == 1

	// 尾插
	tailInsert := golists.NewLinkedList()
	for i := 0; i < 300; i++ {
		tailInsert.Add(i)
	}
	tailInsert.AddByIndex(299, 300)
	tailInsert.AddByIndex(300, 301)
	result = result && tailInsert.Get(299) == 300
	result = result && tailInsert.Get(300) == 301
	result = result && tailInsert.Get(298) == 298
	result = result && tailInsert.Size() == 302
	result = result && linkedListTestBase(tailInsert.(*collection.LinkedList), false)

	// 索引节点插入
	nodeInsert := golists.NewLinkedList()
	for i := 0; i < 300; i++ {
		nodeInsert.Add(i)
	}
	nodeInsert.AddByIndex(0, 300)
	nodeInsert.AddByIndex(100, 301)
	result = result && nodeInsert.Get(301) == 299
	result = result && nodeInsert.Get(0) == 300
	result = result && nodeInsert.Get(1) == 0
	result = result && nodeInsert.Get(100) == 301
	result = result && nodeInsert.Get(101) == 99
	result = result && nodeInsert.Size() == 302
	nodeInsert.Remove(0)
	nodeInsert.Remove(100)
	result = result && linkedListTestBase(nodeInsert.(*collection.LinkedList), false)

	// 头删
	headRemove := golists.NewLinkedList()
	for i := 0; i < 300; i++ {
		headRemove.Add(i)
	}
	headRemove.Remove(0)
	headRemove.Remove(1)
	result = result && headRemove.Size() == 298
	result = result && headRemove.Get(0) == 1
	result = result && linkedListTestBase(headRemove.(*collection.LinkedList), true)

	// 尾删
	tailRemove := golists.NewLinkedList()
	for i := 0; i < 300; i++ {
		tailRemove.Add(i)
	}
	tailRemove.Remove(299)
	tailRemove.Remove(298)
	tailRemove.Remove(300)
	result = result && tailRemove.Size() == 298
	result = result && tailRemove.Get(297) == 297
	result = result && linkedListTestBase(tailRemove.(*collection.LinkedList), true)

	// 索引节点删除
	nodeRemove := golists.NewLinkedList()
	for i := 0; i < 301; i++ {
		nodeRemove.Add(i)
	}
	nodeRemove.Remove(300)
	nodeRemove.Remove(100)
	nodeRemove.Remove(200)
	result = result && nodeRemove.Size() == 298
	result = result && nodeRemove.Get(201) == 203
	result = result && linkedListTestBase(nodeRemove.(*collection.LinkedList), true)

	t.Logf("TestBoundary: %v", result)
}

// 迭代测试
func TestLinkedListForEach(t *testing.T) {
	list := golists.NewLinkedList()
	//init
	DefaultRange := 4000
	for i := 1; i <= DefaultRange; i++ {
		list.Add(i)
	}
	result := make([]int, 0)
	list.Foreach(func(t interface{}) {
		item := t.(int)
		result = append(result, item)
	})
	t.Logf("foreach result: %v", len(result) == DefaultRange)
}

func linkedListInitRandomInsertContainer(randomInsertList collection.List, ignoreSample []int) {
	for i := 0; i < base.DefaultRange; i++ {
		hitRandomIdx := false
		for j := 0; j < len(ignoreSample); j++ {
			if ignoreSample[j] == i {
				hitRandomIdx = true
				continue
			}
		}
		if hitRandomIdx {
			continue
		}
		randomInsertList.Add(i)
	}
}

// 测试 LinkedList 是否正常
// 1. 节点前后相连正常
// 2. 索引位置正常
// 3. isSorted 为 true 时，校验节点递增
func linkedListTestBase(list *collection.LinkedList, isSorted bool) bool {
	if list == nil || list.IsEmpty() {
		return false
	}
	result := true
	result = result && linkedListTestNode(list.Head, isSorted)
	// 保证索引节点对应
	var node = list.Head
	for i := 0; i < list.Info.SkipIndex.Len(); i++ {
		if i%list.Info.SkipIndex.IndexRange == 0 {
			skipIndex := list.Info.SkipIndex.Idx[i/list.Info.SkipIndex.IndexRange]
			result = result && skipIndex.Value() == node.Value()
		}
		node = node.GetNext()
	}
	return result
}

func linkedListTestNode(element *linkedlist.Element, isSorted bool) bool {
	result := true
	if element == nil {
		return result
	}
	if element.GetNext() == nil {
		return result
	}
	// element.next.prev = element
	result = result && element.GetNext().GetPrev().Value().(int) == element.Value().(int)
	if isSorted {
		// element.next > element
		result = result && element.GetNext().Value().(int) > element.Value().(int)
	}
	return result && linkedListTestNode(element.GetNext(), isSorted)
}
