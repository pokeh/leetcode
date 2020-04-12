/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    multOfTwo := false
    curr := head
    mid := head

    for curr != nil {
        curr = curr.Next

        if multOfTwo {
            mid = mid.Next
        }
        multOfTwo = !multOfTwo
    }

    return mid
}

func middleNode2(head *ListNode) *ListNode {
    curr := head
    mid := head

    for curr != nil && curr.Next != nil {
        curr = curr.Next.Next
        mid = mid.Next
    }

    return mid
}
