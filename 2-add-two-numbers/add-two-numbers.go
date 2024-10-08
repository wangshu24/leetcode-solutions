/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    temp := result
    carryOver := 0
    for l1!=nil  || l2 !=nil || carryOver > 0 {
        sum:=carryOver
        if l1!=nil{
            sum+=l1.Val
            l1 = l1.Next  
        }
        if l2!=nil{
            sum+=l2.Val
            l2 = l2.Next  
        }
        temp.Next = &ListNode{ Val : sum%10, Next : nil}
        carryOver = sum/10
        temp = temp.Next
    }

    return result.Next
}