/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersection = function(nums1, nums2) {
    let result = []
    let nums1Obj = {}
    let nums2Obj = {}
    for(let i of nums1){
       nums1Obj[i] =  (nums1Obj[i] || 0) + 1
    }
    for(let i of nums2){
       nums2Obj[i] =  (nums2Obj[i] || 0) + 1
    }

    if(nums1.length < nums2.length){
        for(let i in nums2Obj){          
            if(nums1Obj[i] && !result.includes(i)) {result.push(parseInt(i))}
        }
    } else {
         for(let i in nums1Obj){               
            if(nums2Obj[i] && !result.includes(i)) {result.push(parseInt(i))}
        }
    }
    return result
};