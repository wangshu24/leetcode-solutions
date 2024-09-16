/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    let duplicate = false;
    let map = new Map()
    for(let num of nums){
        if(map.get(num)===1){duplicate = true}
        map.set(num, (map.get(num) || 0) + 1);
    }
    return duplicate
    };