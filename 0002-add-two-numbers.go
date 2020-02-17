package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// case that either is zero
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	return helper(l1, l2, 0)
}

func helper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	var val1, val2 int
	if l1 != nil {
		val1 = l1.Val
	}
	if l2 != nil {
		val2 = l2.Val
	}

	newVal := val1 + val2 + carry
	newCarry := 0

	if newVal >= 10 {
		newCarry = newVal / 10
		newVal = newVal % 10
	}

	var next1, next2 *ListNode
	if l1 != nil {
		next1 = l1.Next
	}
	if l2 != nil {
		next2 = l2.Next
	}

	if next1 == nil && next2 == nil && newCarry == 0 {
		return &ListNode{
			Val:  newVal,
			Next: nil,
		}
	}

	return &ListNode{
		Val:  newVal,
		Next: helper(next1, next2, newCarry),
	}
}
