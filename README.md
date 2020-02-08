# gocollection
基于Go语言实现常用结构,封装原生数组常见功能

## 目前支持的集合

1. List
    1. ListedList
    2. ArrayList
2. [todo] Set
3. [todo] SortedSet
4. [todo] Map 
5. [todo] SortedMap 
6. [todo] MultiMap


### List

#### LinkedList(增强)

接口定义：[List 接口](https://github.com/tejchen/gocollection/blob/master/golists/collection/interface.go)

集合介绍:
1. 基于双向链表实现，拥有有优秀的增删性能
    1. 在长度为1w的情况下，随机插入100条数据，性能约为go原生数组的20倍
2. 实现了简单跳表索引，保持随机读速度，时间复杂度为O(x/100+x%100)
    1. 在长度为1w的情况下，随机读取1000次，性能约为go原生数组的1/5
    2. 在长度为1000w的情况下，随机读取100w次，性能约为go原生数组的1/5

使用入门：

```go
package main
import "fmt"
import "github.com/tejchen/gocollection/golists"

func main(){
    list := golists.NewLinkedList()
    list.Add("i am 1")
    list.Add("i am 2")
    list.Add("i am 3")
    list.AddByIndex(0, "i am 0")
    list.Get(3) // return "i am 3"
    list.Remove(3) // remove "i am 3"
    list.Foreach(func(item interface{}) {
        i := item.(string)
        fmt.Println(i)
    })
}
```

#### ArrayList

接口定义：[List 接口](https://github.com/tejchen/gocollection/blob/master/golists/collection/interface.go)

集合介绍:

1. 基于数组实现，天生支持索引，拥有优秀的查询性能
    1. 数组为连续内存，可直接根据index计算偏移量定位到内存，复杂度为O(1)  
2. 占用空间小

使用入门：

```go
package main
import "fmt"
import "github.com/tejchen/gocollection/golists"

func main(){
    list := golists.NewArrayList()
    list.Add("i am 1")
    list.Add("i am 2")
    list.Add("i am 3")
    list.AddByIndex(0, "i am 0")
    list.Get(3) // return "i am 3"
    list.Remove(3) // remove "i am 3"
    list.Foreach(func(item interface{}) {
        i := item.(string)
        fmt.Println(i)
    })
}
```