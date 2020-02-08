package test

import (
	"fmt"
	"github.com/tejchen/gocollection/golists"
	"github.com/tejchen/gocollection/test/base"
	"testing"
	"time"
)

//写性能测试
func TestLinkedListWritePerformanceService(t *testing.T) {
	listStart := time.Now().UnixNano()
	linkedList := golists.NewLinkedList()
	for i := 0; i < base.PerformanceTestRange; i++ {
		linkedList.Add(i)
	}
	fmt.Println(fmt.Sprintf("linkedList time: %v", time.Now().UnixNano()-listStart))

	// 数组对比
	start := time.Now().UnixNano()
	array := make([]interface{}, 0)
	for i := 0; i < base.PerformanceTestRange; i++ {
		array = append(array, i)
	}
	fmt.Println(fmt.Sprintf("array time:      %v", time.Now().UnixNano()-start))
}

//随机写性能测试
func TestLinkedListRandomWritePerformanceService(t *testing.T) {
	// 初始化
	linkedList := golists.NewLinkedList()
	for i := 0; i < base.DefaultRange; i++ {
		linkedList.Add(i)
	}
	array := make([]interface{}, 0)
	for i := 0; i < base.DefaultRange; i++ {
		array = append(array, i)
	}

	// 取样本
	testSample := base.GetRangeTestSample(base.DefaultRange, 100)

	// list 测试
	listStart := time.Now().UnixNano()
	for i := range testSample {
		linkedList.Add(testSample[i])
	}
	fmt.Println(fmt.Sprintf("linkedList time: %v", time.Now().UnixNano()-listStart))

	arrayStart := time.Now().UnixNano()
	for i := range testSample {
		temp1 := array[:testSample[i]]
		temp2 := array[testSample[i]:]
		array = append(append(temp1, testSample[i]), temp2...)
	}
	fmt.Println(fmt.Sprintf("array time:      %v", time.Now().UnixNano()-arrayStart))
}

//大数据量随机读测试
func TestBigLinkedListReadPerformanceService(t *testing.T) {
	linkedList := golists.NewLinkedList()
	for i := 0; i < base.PerformanceTestRange; i++ {
		linkedList.Add(i)
	}

	// 数组对比
	array := make([]interface{}, 0)
	for i := 0; i < base.PerformanceTestRange; i++ {
		array = append(array, i)
	}

	testSample := base.GetRangeTestSample(base.PerformanceTestRange, 10)
	fmt.Println("init test sample success")

	//array
	listStart := time.Now().UnixNano()
	listTemp := make([]interface{}, len(testSample))
	for idx := range testSample {
		listTemp[idx] = linkedList.Get(testSample[idx])
	}
	fmt.Println(fmt.Sprintf("linkedList time: %v", time.Now().UnixNano()-listStart))

	//array
	start := time.Now().UnixNano()
	temp := make([]interface{}, len(testSample))
	for idx := range testSample {
		temp[idx] = array[testSample[idx]]
	}
	fmt.Println(fmt.Sprintf("array time:      %v", time.Now().UnixNano()-start))
}

//少数据量随机读测试
func TestLinkedListReadPerformanceService(t *testing.T) {
	linkedList := golists.NewLinkedList()
	for i := 0; i < base.DefaultRange; i++ {
		linkedList.Add(i)
	}

	// 数组对比
	array := make([]interface{}, 0)
	for i := 0; i < base.DefaultRange; i++ {
		array = append(array, i)
	}

	testSample := base.GetRangeTestSample(base.DefaultRange, 100)
	fmt.Println("init test sample success")

	//array
	listStart := time.Now().UnixNano()
	listTemp := make([]interface{}, len(testSample))
	for idx := range testSample {
		listTemp[idx] = linkedList.Get(testSample[idx])
	}
	fmt.Println(fmt.Sprintf("linkedList time: %v", time.Now().UnixNano()-listStart))

	//array
	start := time.Now().UnixNano()
	temp := make([]interface{}, len(testSample))
	for idx := range testSample {
		temp[idx] = array[testSample[idx]]
	}
	fmt.Println(fmt.Sprintf("array time:      %v", time.Now().UnixNano()-start))
}
