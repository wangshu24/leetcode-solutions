/**
 * @param {number[]} nums
 * @param {number} k
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var rotate = function(nums, k) {
    const arrayLength = nums.length;
    const copied = nums.map(x => x)
    for(let i = 0; i< arrayLength;i++){
        let newIndex;  
        if((i+k) >= arrayLength) {
            newIndex = (i+k)%arrayLength
        } else { newIndex = i+k}
        nums[newIndex] = copied[i]
        
    }
};