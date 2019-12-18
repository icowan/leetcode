/**
 * @Time : 18/12/2019 2:21 PM
 * @Author : solacowa@gmail.com
 * @File : revese_test
 * @Software: GoLand
 */

package linked_list

import "testing"

func TestReverseSingleNode(t *testing.T) {
	node := ReverseSingleNode(&SingleNode{
		Value: 1,
		Next: &SingleNode{
			Value: 2,
			Next: &SingleNode{
				Value: 3,
				Next: &SingleNode{
					Value: 4,
					Next: &SingleNode{
						Value: 5,
						Next:  nil,
					},
				},
			},
		},
	})

	for node != nil {
		t.Log("val", node.Value)
		node = node.Next
	}
}

func TestReverseDoubleNode(t *testing.T) {
	node := ReverseDoubleNode(&DoubleNode{
		Value: 1,
		Last: &DoubleNode{
			Value: 2,
			Next: &DoubleNode{
				Value: 3,
				Next:  nil,
				Last:  nil,
			},
			Last: &DoubleNode{
				Value: 4,
				Next:  nil,
				Last:  nil,
			},
		},
		Next: &DoubleNode{
			Value: 5,
			Next: &DoubleNode{
				Value: 6,
				Next:  nil,
				Last:  nil,
			},
			Last: &DoubleNode{
				Value: 7,
				Next:  nil,
				Last:  nil,
			},
		},
	})

	for node != nil {
		t.Log("pre", node.Last.Value)
		t.Log("next", node.Next.Value)
		t.Log("val", node.Value)
		node = node.Next
	}
}
