/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    let prefix = "";
    const sorted = strs.sort();
    const longestStr = sorted[sorted.length-1]
    for(let i=0; i < longestStr.length; i++){
        if(strs.every((word) => word[i] === longestStr[i])){
            prefix+=longestStr[i];
        } else {break}
    }
    return prefix
};