package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	arr := []*ListNode{head}
	cur := head.Next
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	cur = arr[len(arr)-1]
	r := cur
	for i := len(arr) - 2; i >= 0; i-- {
		cur.Next = arr[i]
		cur = cur.Next
	}
	cur.Next = nil
	return r
}
