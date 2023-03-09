package leetcode_876

import "example.com/m/util"

/*
给你单链表的头结点 head ，请你找出并返回链表的中间结点。
如果有两个中间结点，则返回第二个中间结点。

思路：
1. 快慢指针，慢指针指向1，快指针指向2，快指针是慢指针的2倍
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *util.ListNode) *util.ListNode {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head
	for {
		slow = slow.Next
		for i := 0; i < 2 && fast.Next != nil; i++ {
			fast = fast.Next
		}
		if fast.Next == nil {
			return slow
		}
	}
}

func middleNode2(head *util.ListNode) *util.ListNode {
	if head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
