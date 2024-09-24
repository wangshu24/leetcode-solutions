/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    const stringArr = s.trim().split(' ')
    return stringArr[stringArr.length-1].length
};