/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {

    let map = new Map()
    for(let num of nums){
        if(map.get(num)!==undefined){return true}
        map.set(num, (map.get(num) || 0) + 1);
    }
    return false
    };