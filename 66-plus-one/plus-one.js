/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
    // let digitStr = ""
    // for(let d of digits){
    //     digitStr += d
    // }
    // console.log(digitStr)
    // let result = []
    // let resultStr = (parseInt(digitStr) + 1).toString()
    // console.log(resultStr)
    // for(let i =0; i < resultStr.length; i++){
    //     result.push(parseInt(resultStr[i]))
    // }
    if( digits[digits.length-1] !== 9){
        digits[digits.length-1] = digits[digits.length-1] + 1
        console.log()
    } else {
        digits[digits.length-1] = 0
        n = 2
        while(digits[digits.length - n] === 9 ){
            digits[digits.length - n] = 0
            n++
        }
        digits.length - n < 0 ? digits.splice(0,0,1) : digits[digits.length - n] = digits[digits.length - n] + 1
    }
    return digits
};