/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
    let result = 0
    for(let i = 0; i< nums.length; i++){
        if(nums[i] === val){
            nums[i] = undefined
            result++
            }
    }
    nums.sort((a,b) => a === undefined ? 1 : 0)

    return nums.length - result
};