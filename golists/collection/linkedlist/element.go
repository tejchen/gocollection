package linkedlist

type Element struct {
	val        interface{}
	next, prev *Element
}

func (e *Element) AddNext(element *Element) {
	if element == nil {
		e.next = nil
		return
	}
	e.next = element
	element.prev = e
}

func (e *Element) AddPrev(element *Element) {
	if element == nil {
		e.prev = nil
		return
	}
	element.next = e
	e.prev = element
}

func (e *Element) Value() interface{} {
	return e.val
}

func (e *Element) GetNext() *Element {
	return e.next
}

func (e *Element) GetPrev() *Element {
	return e.prev
}

func (e *Element) SetNext(next *Element) {
	e.next = next
}

func (e *Element) SetPrev(prev *Element) {
	e.prev = prev
}

func DefaultElement(item interface{}) *Element {
	newElement := &Element{
		val:  item,
		prev: nil,
		next: nil,
	}
	return newElement
}
