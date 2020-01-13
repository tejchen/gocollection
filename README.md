# gocollection
用Go语言封装的常见结（lun）构（zi）

## 目前支持的集合
### List
实现类：**增强版 LinkedList**

接口定义：[List 接口](https://github.com/tejchen/gocollection/blob/master/golists/interface.go)

集合介绍:

如果是 Java 背景的小伙伴，应该对 LinkedList 和 ArrayList 两个集合很熟悉了。

Go’s LinkedList 吸收了这两个结构的优缺点，也做了一些改良，下面稍微介绍以及对比一下：

**Java‘s LinkedList：**
1. 底层是双向链表
2. 插入/删除性能优秀，但是随机读性能较差

**Java‘s ArrayList：**
1. 底层是数组，外层加了索引
2. 随机读性能优秀，随机插入/删除性能较差

**gocollection’s LinkedList**
1. 支持 Java’s List 大部分 API
2. 增加 Go 特色 API
3. 底层是双向链表，拥有良好的插入删除性能
4. 增加了简单的跳表索引，固定每100个元素增加一个索引节点，增加随机读性能

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