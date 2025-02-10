/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 {return head}
    slow, fast, tmp := head, head, head
    l := 0
    for tmp != nil {
        tmp = tmp.Next    
        l++
    }
    if l == 0 || l == 1 {return head}
    if k > l {k = k%l}
    for i:=0; i < k; i++ {
        fast = fast.Next
    }
    if fast == nil {
        return head
    }
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    fast.Next = head
    newHead := slow.Next
    slow.Next = nil

    return newHead
}