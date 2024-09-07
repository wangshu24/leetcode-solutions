/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    let result = [];
    let duplicate = [];
    nums.forEach(n => {
        //Remove if statement to improve runtime
        !result.includes(n) && result.push(n)
    })
    // for(let i = 0; i<result.length; i++){
    //     nums[i] = result[i]
    // }
    for(let i = 0; i < nums.length; i++) {
        result[i] === 0 || result[i] ? nums[i] = result[i] : nums[i] = "_";
    }
    return result.length;
};
