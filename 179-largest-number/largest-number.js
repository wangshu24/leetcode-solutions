/**
 * @param {number[]} nums
 * @return {string}
 */
var largestNumber = function(nums) {
    if(nums.every((e)=> e === 0)) {return "0"}
    nums.sort((a,b) => compareNumber(a,b))
    let string = ""
    nums.forEach((e)=> string+=e.toString())
    return string
};

var compareNumber = function(num1, num2){
    let strNum2 = num2.toString()
    let strNum1 = num1.toString()
    if(parseInt(strNum2+strNum1) > parseInt(strNum1+strNum2)) { return 1} else {return -1}
    return 0
}