package _6_linkedlist

import "testing"

func TestLinkedList_InterBefore(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 10; i++ {
		linkedList.InsertToTail(i + 1)
	}
	linkedList.Print()
	linkedList.InterBefore(linkedList.FindByIndex(1), 11)
	linkedList.Print()
}

func TestLinkedList_InsertToHead(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 10; i++ {
		linkedList.InsertToHead(i + 1)
	}
	linkedList.Print()
}

func TestLinkedList_InsertToTail(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 10; i++ {
		linkedList.InsertToTail(i + 1)
	}
	linkedList.Print()
}

func TestLinkedList_FindByIndex(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 10; i++ {
		linkedList.InsertToTail(i + 1)
	}
	t.Log(linkedList.FindByIndex(2))
	t.Log(linkedList.FindByIndex(5))
}

func TestLinkedList_DeleteNode(t *testing.T) {
	linkedList := NewLinkedList()
	for i := 0; i < 10; i++ {
		linkedList.InsertToTail(i + 1)
	}
	linkedList.Print()
	linkedList.DeleteNode(linkedList.FindByIndex(1))
	linkedList.Print()
}
