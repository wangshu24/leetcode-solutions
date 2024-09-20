/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function(nums, val) {
    for(let i = 0; i < nums.length; i++){
        if(nums[i] === val){
            nums[i] = undefined
        }
    }
    nums.sort((a,b) => a === undefined ? 1 : 0)
    let result = nums.filter(e => e !== undefined).length
    return result
};