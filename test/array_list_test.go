package test

import (
	"github.com/tejchen/gocollection/golists"
	"github.com/tejchen/gocollection/golists/collection"
	"github.com/tejchen/gocollection/test/base"
	"testing"
	"time"
)

//基础服务测试
func TestArrayListBaseService(t *testing.T) {
	var result = true

	for i := 0; i < 100; i++ {
		// 基础测试
		list := golists.NewArrayList()
		result = result && list.Size() == 0
		result = result && list.IsEmpty() == true
		result = result && list.IsNotEmpty() == false

		// 准备数据
		for i := 0; i < base.DefaultRange; i++ {
			list.Add(i)
		}

		result = result && arrayListTestBase(list.(*collection.ArrayList), true)

		// 顺序读测试
		for i := 0; i < base.DefaultRange; i++ {
			result = result && list.Get(i) == i
		}

		// 随机读测试 1/5
		randomTestSample := base.GetRandomTestSample(base.DefaultRange / 5)
		for i := 0; i < len(randomTestSample); i++ {
			result = result && randomTestSample[i] == list.Get(randomTestSample[i])
		}

		result = result && arrayListTestBase(list.(*collection.ArrayList), true)

		result = result && list.Size() == base.DefaultRange
		result = result && list.IsEmpty() == false
		result = result && list.IsNotEmpty() == true
	}

	t.Logf("arrayListTestBaseService: %v", result)
}

// 随机插入测试
func TestArrayListRandomInsert(t *testing.T) {
	var randomInsertCheck = true

	var timeSum int64 = 0

	for i := 0; i < 100; i++ {
		// 容器
		randomInsertList := golists.NewArrayList()

		// 随机样本数据
		randomTestSample := base.GetRandomTestSample(base.DefaultRange / 5)

		// 初始化数据
		for i := 0; i < base.DefaultRange; i++ {
			randomInsertList.Add(i)
		}

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

		randomInsertCheck = randomInsertCheck && arrayListTestBase(randomInsertList.(*collection.ArrayList), true)
	}
	t.Logf("TestRandomInsert time: %v", timeSum/100)
	t.Logf("TestRandomInsert: %v", randomInsertCheck)
}

// 随机删除测试
func TestArrayListRandomRemove(t *testing.T) {
	var randomRemoveCheck = true
	var timeSum int64 = 0

	for i := 0; i < 100; i++ {
		// 容器
		randomRemoveList := golists.NewArrayList()

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

		randomRemoveCheck = randomRemoveCheck && arrayListTestBase(randomRemoveList.(*collection.ArrayList), true)
	}

	t.Logf("TestRandomRemove time: %v", timeSum/100)
	t.Logf("TestRandomRemove: %v", randomRemoveCheck)
}

// 定点边界测试
func TestArrayListBoundary(t *testing.T) {
	result := true
	// 头插
	headInsert := golists.NewArrayList()
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
	tailInsert := golists.NewArrayList()
	for i := 0; i < 300; i++ {
		tailInsert.Add(i)
	}
	tailInsert.AddByIndex(299, 300)
	tailInsert.AddByIndex(300, 301)
	result = result && tailInsert.Get(299) == 300
	result = result && tailInsert.Get(300) == 301
	result = result && tailInsert.Get(298) == 298
	result = result && tailInsert.Size() == 302
	result = result && arrayListTestBase(tailInsert.(*collection.ArrayList), false)

	// 头删
	headRemove := golists.NewArrayList()
	for i := 0; i < 300; i++ {
		headRemove.Add(i)
	}
	headRemove.Remove(0)
	headRemove.Remove(1)
	result = result && headRemove.Size() == 298
	result = result && headRemove.Get(0) == 1
	result = result && arrayListTestBase(headRemove.(*collection.ArrayList), true)

	// 尾删
	tailRemove := golists.NewArrayList()
	for i := 0; i < 300; i++ {
		tailRemove.Add(i)
	}
	tailRemove.Remove(299)
	tailRemove.Remove(298)
	tailRemove.Remove(300)
	result = result && tailRemove.Size() == 298
	result = result && tailRemove.Get(297) == 297
	result = result && arrayListTestBase(tailRemove.(*collection.ArrayList), true)

	t.Logf("TestBoundary: %v", result)
}

// 迭代测试
func TestArrayListForEach(t *testing.T) {
	list := golists.NewArrayList()
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

// 测试 ArrayList 是否正常
func arrayListTestBase(list *collection.ArrayList, isSort bool) bool {
	if list == nil || list.IsEmpty() {
		return false
	}
	result := true
	if isSort {
		temp := -1
		list.Foreach(func(item interface{}) {
			result = result && item.(int) > temp
			temp = item.(int)
		})
	}
	return result
}
