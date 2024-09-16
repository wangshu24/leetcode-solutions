/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findDuplicates = function(nums) {
    let freq = {}
    let result = []
    for(let i of nums){
        freq[i] = (freq[i] || 0) + 1;
        
    }
    for(let x in freq){
        if(freq[x]>1) {result.push(x)}
    }

    return result;
};