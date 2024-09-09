/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    const numString = x.toString();
    const numLength = numString.length;
    if(numLength === 1) return true;
    let iterator = Math.floor(numLength/2);
    
    let isPalindrome  = true;
    for(let i = 0; i < iterator; i++){
        if(numString[i] !== numString[numLength-1-i]) {isPalindrome = false}
    }
    return isPalindrome;
};