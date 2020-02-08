package base

import (
	"math/rand"
	"time"
)

// 功能测试样本数量
var DefaultRange = 10000

// 性能测试样本数量
var PerformanceTestRange = 10000000

func GetRangeTestSample(testRange int, rangeSize int) []int {
	array := make([]int, 0)
	for i := 0; i < testRange; i++ {
		if i%rangeSize == 0 {
			array = append(array, i)
		}
	}
	return array
}

func GetRandomTestSample(cap int) []int {
	array := make([]int, 0)
	for i := 0; i < cap; i++ {
		rand.Seed(time.Now().UnixNano())
		randomInt := rand.Intn(DefaultRange)
		diff := false
		for idx := range array {
			if array[idx] == randomInt {
				diff = true
			}
		}
		if diff {
			i--
			continue
		}
		array = append(array, randomInt)
	}
	return array
}
