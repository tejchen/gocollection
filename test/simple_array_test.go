package test

import (
	"github.com/tejchen/gocollection/goarrays"
	"testing"
)

func TestBaseService_SimpleArray(t *testing.T) {
	result := true
	// assert
	array := goarrays.NewSimpleArray()
	result = result && array.IsEmpty()
	result = result && !array.IsNotEmpty()
	// append
	for i := 0; i < 2000; i++ {
		array.Append(i)
	}
	result = result && array.Len() == 2000
	// append all
	array1 := goarrays.NewSimpleArray()
	for i := 2000; i < 4000; i++ {
		array1.Append(i)
	}
	array.AppendAll(array1)
	result = result && array.Len() == 4000
	// replace
	array.Replace(3000, 30000)
	result = result && array.Len() == 4000
	result = result && array.Get(3000) == 30000
	// remove
	array.Remove(0)
	result = result && array.Len() == 3999
	result = result && array.Get(0) == 1
	array.Remove(3998)
	result = result && array.Len() == 3998
	result = result && array.Get(3997) == 3998
	array.Remove(2000)
	result = result && array.Len() == 3997
	result = result && array.Get(2000) == 2002
	// assert
	result = result && !array.IsEmpty()
	result = result && array.IsNotEmpty()
	t.Logf("TestBaseService_SimpleArray:%v", result)
}

func TestForeach_SimpleArray(t *testing.T) {
	result := true
	array := goarrays.NewSimpleArray()
	for i := 0; i < 2000; i++ {
		array.Append(i)
	}
	newArray := goarrays.NewSimpleArray()
	for i := 0; i < 2000; i++ {
		newArray.Append(i)
	}
	array.Foreach(func(item interface{}) {
		result = result && newArray.Interface()[item.(int)] == item.(int)
	})
	t.Logf("TestForeach_SimpleArray:%v", result)
}
