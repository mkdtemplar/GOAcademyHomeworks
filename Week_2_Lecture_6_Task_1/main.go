package main

import "fmt"

type MagicList struct {
	LastItem *Item
}

type Item struct {
	Value    int
	PrevItem *Item
}

func add(l *MagicList, value int) {

	v := &Item{Value: value}
	if l.LastItem == nil {
		l.LastItem = v
	} else {
		current := l.LastItem
		for current.PrevItem != nil {
			current = current.PrevItem
		}

		current.PrevItem = v
	}

}

func addFront(l *MagicList, value int) {

	v := &Item{Value: value}

	if l.LastItem == nil {
		l.LastItem = v
	} else {
		v.PrevItem = l.LastItem
		l.LastItem = v
	}
}

func Display(lll *Item) {
	for lll != nil {
		fmt.Print(lll.Value, "->")
		lll = lll.PrevItem
	}
	fmt.Println("NIL")
}

func (ll *MagicList) Reverse() {

	var lSlice []int
	currentNode := ll.LastItem
	var next *Item
	var previousNode *Item

	for currentNode != nil {
		next, currentNode.PrevItem = currentNode.PrevItem, previousNode
		previousNode, currentNode = currentNode, next
	}
	ll.LastItem = previousNode

	Display(ll.LastItem)

	current := ll.LastItem
	for current != nil {
		lSlice = append(lSlice, current.Value)
		current = current.PrevItem
	}
	fmt.Println("Slice from the reversed LinkedList: ", lSlice)
}
func ReverseLinkedList(head *MagicList) {

}
func Size(s *MagicList) int {

	current := s.LastItem
	var len int

	for current != nil {
		len++
		current = current.PrevItem
	}
	return len
}

func (rf *MagicList) RemoveAtFront() error {
	if IsEmpty(rf) {
		return fmt.Errorf("list is empty")
	}
	rf.LastItem = rf.LastItem.PrevItem
	return nil
}

func (rb *MagicList) RemoveBack() error {
	if IsEmpty(rb) {
		return fmt.Errorf(" List is empty")
	}
	var previous *Item
	current := rb.LastItem
	for current.PrevItem != nil {
		previous = current
		current = current.PrevItem
	}
	if previous != nil {
		previous.PrevItem = nil
	} else {
		rb.LastItem = nil
	}

	return nil
}

func IsEmpty(e *MagicList) bool {
	if e.LastItem == nil {
		return true
	} else {
		return false
	}
}

func DisplayList(dl *MagicList) {
	current := dl.LastItem
	for current != nil {
		fmt.Print(current.Value, "->")
		current = current.PrevItem
	}
	fmt.Println("NIL")
}

func main() {

	l := &MagicList{}
	add(l, 10)
	add(l, 22)
	add(l, 44)
	fmt.Println("Magic list elements pushed at back:")
	DisplayList(l)

	fmt.Println("Adding value in front of first element!")
	addFront(l, 77)
	//l.RemoveAtFront()
	//l.RemoveBack()
	DisplayList(l)

	fmt.Println("Reverse Linked List:")
	l.Reverse()
	fmt.Println("Length: ", Size(l))

}
