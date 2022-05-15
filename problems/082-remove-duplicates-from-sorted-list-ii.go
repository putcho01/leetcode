package main

// Two Pointers
//https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	//Two Pointers
	var pre *ListNode
	cur := head

	for cur != nil && cur.Next != nil {
		if cur.Val != cur.Next.Val {
			pre = cur
		} else {
			for cur.Next != nil && cur.Next.Val == cur.Val {
				cur = cur.Next
			}
			if pre != nil {
				pre.Next = cur.Next
			} else {
				head = cur.Next
			}
		}
		cur = cur.Next
	}

	return head
}
