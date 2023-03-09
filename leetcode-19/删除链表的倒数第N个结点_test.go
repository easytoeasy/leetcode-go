package leetcode_19

import "example.com/m/util"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 记录前继节点
func removeNthFromEnd(head *util.ListNode, n int) *util.ListNode {
	if head == nil {
		return head
	}
	var h = &util.ListNode{}
	var pre *util.ListNode
	h.Next = head

	slow, fast := h, h

	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}

	// slow即为待删除节点
	pre.Next = pre.Next.Next
	return h.Next
}

// 无需记录前继节点
func removeNthFromEnd2(head *util.ListNode, n int) *util.ListNode {
	if head == nil {
		return head
	}
	var h = &util.ListNode{}
	h.Next = head

	slow, fast := h, h

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// slow即为待删除节点
	slow.Next = slow.Next.Next
	return h.Next
}

/*
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

思路：
1. 循环一次，从后删除
2. 快慢指针，间隔为n, 每次记录前继节点
3. 快慢指针，间隔为n+1,无需记录前继节点

增加空的头指针，适用于删除头节点
*/
