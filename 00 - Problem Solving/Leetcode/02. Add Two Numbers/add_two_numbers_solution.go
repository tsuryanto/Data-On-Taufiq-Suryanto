package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var curry, sum, val int
	newListNode := ListNode{}

	isFirst := true
	for {
		// fmt.Println("L1:", l1, " L2:", l2)
		if l1 == nil && l2 == nil {
			if curry != 0 {
				AddNode(&newListNode, curry)
			}
			break
		} else {
			sum = curry + replaceNil(l1) + replaceNil(l2)
		}

		if sum >= 10 {
			curry = sum / 10
			val = sum % 10
		} else {
			curry = 0
			val = sum
		}

		if isFirst {
			newListNode.Val = val
			newListNode.Next = nil
			isFirst = false
		} else {
			AddNode(&newListNode, val)
		}
		sum = 0
		if l1 != nil {
			if l1.Next == nil {
				l1 = nil
			} else {
				l1 = l1.Next
			}
		}
		if l2 != nil {
			if l2.Next == nil {
				l2 = nil
			} else {
				l2 = l2.Next
			}
		}
	}

	return &newListNode
}

func replaceNil(n *ListNode) int {
	if n != nil {
		return n.Val
	} else {
		return 0
	}
}

func AddNode(n *ListNode, data int) {
	newNode := ListNode{data, nil}
	iter := n
	for iter.Next != nil {
		iter = iter.Next
	}
	iter.Next = &newNode
}
