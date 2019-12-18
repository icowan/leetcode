/**
 * @Time : 18/12/2019 5:50 PM
 * @Author : solacowa@gmail.com
 * @File : list
 * @Software: GoLand
 */

package linked_list

type ListNode struct {
	prev  *ListNode
	next  *ListNode
	value int
}

type List struct {
	head *ListNode
	tail *ListNode
	len  int
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (l *List) Head() *ListNode {
	return l.head
}

func (l *List) Tail() *ListNode {
	return l.tail
}

func (l *List) Len() int {
	return l.len
}

func (l *List) RPush(val int) {
	node := NewListNode(val)
	if l.Len() == 0 {
		l.head = node
		l.tail = node
	} else {
		tail := l.tail
		tail.next = node
		node.prev = tail
		l.tail = node
	}

	l.len = l.len + 1
}

func (l *List) LPop() *ListNode {
	if l.len == 0 {
		return nil
	}

	if l.head.next == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
	}

	l.len = l.len - 1

	return
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		value: val,
	}
}

func (n *ListNode) Prev() *ListNode {
	return n.prev
}

func (n *ListNode) Next() *ListNode {
	return n.next
}

func (n *ListNode) Value() int {
	if n == nil {
		return 0
	}
	return n.value
}
