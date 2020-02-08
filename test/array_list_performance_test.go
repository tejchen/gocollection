package test

import (
	"fmt"
	"github.com/tejchen/gocollection/golists"
	"github.com/tejchen/gocollection/test/base"
	"testing"
	"time"
)

func TestWritePerformanceService_SimpleArray(t *testing.T) {
	simpleArray := golists.NewArrayList()
	start := time.Now().UnixNano()
	for i := 0; i < base.PerformanceTestRange; i++ {
		simpleArray.Add(i)
	}
	fmt.Println(fmt.Sprintf("simpleArray time: %v", time.Now().UnixNano()-start))

	// 数组对比
	array := make([]interface{}, 0)
	arrayStart := time.Now().UnixNano()
	for i := 0; i < base.PerformanceTestRange; i++ {
		array = append(array, i)
	}
	fmt.Println(fmt.Sprintf("array time:       %v", time.Now().UnixNano()-arrayStart))
}
