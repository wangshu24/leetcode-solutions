/**
 * @param {string[]} strs
 * @return {string}
 */

var longestCommonPrefix = function(strs) {
    let prefix = ""
    
    if(strs.length === 0){return prefix}

    strs.sort()
    let short = strs[0]
    let long = strs[strs.length-1]

    for(let i=0; i<short.length; i++){
        if(short[i] !== long[i]){
            return prefix
        } else {prefix+=short[i]}
    }

    return prefix
};