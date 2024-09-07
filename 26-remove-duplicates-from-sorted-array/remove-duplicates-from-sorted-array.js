/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    let result = [];
    let duplicate = [];
    nums.forEach(n => {
        !result.includes(n) && result.push(n)
    })
    for(let i = 0; i<result.length; i++){
        nums[i] = result[i]
    }
    for(let i = result.length; i < nums.length; i++) {
        nums[i] = "_"
    }
    return result.length;
};
