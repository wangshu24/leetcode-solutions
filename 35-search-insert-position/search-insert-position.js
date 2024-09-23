/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function(nums, target) {
    if(!nums.find(e => e === target)){ nums.push(target)}
    nums.sort((a,b) => a-b)
    return nums.findIndex(e => e === target) 
};