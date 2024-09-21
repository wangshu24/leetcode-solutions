/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} val
 * @return {ListNode}
 */
var removeElements = function(head, val) {
    if (!head) return null;
    let temp = new ListNode(0) 
    let modifiable = temp
    temp.next = head
    while(modifiable.next !== null){
        if(modifiable.next.val === val) {modifiable.next = modifiable.next.next}
        else {modifiable = modifiable.next}
    }
    
   return temp.next  
};

