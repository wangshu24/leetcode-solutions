/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {return}
    arr1 := []int{}
    var temp = head

    for temp != nil {
        arr1 = append(arr1, temp.Val)
        temp = temp.Next
    }

    var arr2 = arr1[len(arr1)/2:len(arr1)]
    arr1 = arr1[0:len(arr1)/2]

    fmt.Println(arr1, arr2)
    l, r := 1, 0
    var res = head
    for l < len(arr1) || r < len(arr2) {
        if r < len(arr2) {
            fmt.Println(arr2[len(arr2)-1-r])
            res.Next = &ListNode{Val : arr2[len(arr2)-1-r], Next : nil}
            r++
            res = res.Next
        }
        if l < len(arr1) {
            res.Next = &ListNode{ Val: arr1[l], Next : nil}
            l++
            res = res.Next
        }
    }

    head = head.Next
}