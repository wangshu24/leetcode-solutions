/**
 * @param {string} s1
 * @param {string} s2
 * @return {string[]}
 */
var uncommonFromSentences = function(s1, s2) {
    let s1Arr = s1.split(" ");
    let s2Arr = s2.split(" ");
    let s1Obj = {}
    let s2Obj = {}
    let result = []
    for(let i of s1Arr){
        s1Obj[i] = (s1Obj[i] || 0) + 1 
    }
    for(let i of s2Arr){
        s2Obj[i] = (s2Obj[i] || 0) + 1 
    }

    for(let i in s1Obj){
        if(!s2Obj[i] && s1Obj[i] < 2) {result.push(i)}
    }
     for(let i in s2Obj){
        if(!s1Obj[i] && s2Obj[i] < 2) {result.push(i)}
    }

    return result
};