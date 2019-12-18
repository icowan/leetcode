/**
 * @Time : 18/12/2019 2:19 PM
 * @Author : solacowa@gmail.com
 * @File : reverse
 * @Software: GoLand
 */

package linked_list

// 单向链表
type SingleNode struct {
	Value int
	Next  *SingleNode
}

// 双向链表
type DoubleNode struct {
	Value int
	Next  *DoubleNode
	Last  *DoubleNode
}

// 反转单向列表
func ReverseSingleNode(node *SingleNode) *SingleNode {
	var pre, next *SingleNode
	for node != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}

	// 1, 2, 3, 4, 5
	// next = 2, node.Next = nil, pre = 1, node = 2
	// next = 3, node.Next = 1, pre = 2, node = 3
	// next = 4, node.Next = 2, pre = 3, node = 4
	// next = 5, node.Next = 3, pre = 4, node = 5
	// next = nil, node.Next = 4, pre = 5, node = nil

	return pre
}

func ReverseDoubleNode(node *DoubleNode) *DoubleNode {
	var pre, next *DoubleNode
	for node != nil {
		next = node.Next
		node.Next = pre
		node.Last = next
		pre = node
		node = next
	}
	return pre
}
