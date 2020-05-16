/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	firstEven := head.Next
	lastOdd, lastEven := head, head.Next

	node := head.Next.Next
	isOdd := true

	for node != nil {
		nextNode := node.Next

		if isOdd {
			lastOdd.Next = node
			lastOdd = node
			node.Next = firstEven
		} else {
			lastEven.Next = node
			lastEven = node
			node.Next = nil
		}

		isOdd = !isOdd
		node = nextNode
	}

	lastEven.Next = nil

	return head
}
