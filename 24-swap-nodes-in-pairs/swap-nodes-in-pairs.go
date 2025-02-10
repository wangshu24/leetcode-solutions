/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil {return nil}
    if head.Next == nil {return head}

    first := head
    second := first.Next
    third := second.Next
    second.Next = first
    first.Next = swapPairs(third)

    return second
}