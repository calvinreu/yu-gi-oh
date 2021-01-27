package graphic

//Element of list only instance*
type Element struct {
	prev  *Element
	next  *Element
	Value *Instance
}

//List double linked list which can only store instances
type List struct {
	first *Element
	//last  *Element
}

//First Element in list
func (list *List) First() *Element {
	return list.first
}

//Push adds an element at the front of the list
func (list *List) Push(value *Instance) *Element {
	if list.first == nil {
		list.first = &Element{nil, nil, value}
	}

	element := Element{nil, list.first, value}
	list.first.prev = &element
	list.first = &element
	return list.first
}

/*func (list *List) Push(value Instance) *Element {
	element := Element{list.last, nil, &value}
	list.first.next = &element
	list.last = &element
	return list.last
}*/

//Prev returns element previous to self
func (element *Element) Prev() *Element {
	return element.prev
}

//Next returns element previous to self
func (element *Element) Next() *Element {
	return element.next
}
