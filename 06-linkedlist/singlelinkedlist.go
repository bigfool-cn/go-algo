package _6_linkedlist

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(value interface{}) *ListNode {
	return &ListNode{nil, value}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// InsertAfter 插入到给定节点后面
func (this *LinkedList) InsertAfter(p *ListNode, value interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(value)
	oldNextNode := p.next
	newNode.next = oldNextNode
	p.next = newNode
	this.length++
	return true
}

// InterBefore 插入到给定节点之前
func (this *LinkedList) InterBefore(p *ListNode, value interface{}) bool {
	if nil == p || this.head == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(value)
	newNode.next = p
	pre.next = newNode
	this.length++
	return true
}

// InsrtToHead 插入到链头之后
func (this *LinkedList) InsertToHead(value interface{}) bool {
	return this.InsertAfter(this.head, value)
}

// InsertToTail 插入到链尾
func (this *LinkedList) InsertToTail(value interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, value)
}

// FindByIndex 通过给定索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var idx uint
	for idx = 0; idx < index; idx++ {
		cur = cur.next
	}
	return cur
}

// DeleteNode 删除给定节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	pre.next = p.next
	p = nil
	this.length--
	return true
}

// Print 打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.value)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}

	fmt.Println(format)
}
