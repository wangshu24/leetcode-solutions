/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }

    if list2 == nil {
        return list1
    }

    if list1.Val > list2.Val {
        list2.Next = mergeTwoLists(list2.Next, list1) 
        return list2
    } else {
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    }

}