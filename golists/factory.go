package golists

import "github.com/tejchen/gocollection/golists/collection"

func NewLinkedList() collection.List {
	return &collection.LinkedList{}
}

func NewArrayList() collection.List {
	return &collection.ArrayList{}
}
