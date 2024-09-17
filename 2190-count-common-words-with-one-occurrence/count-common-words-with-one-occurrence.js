/**
 * @param {string[]} words1
 * @param {string[]} words2
 * @return {number}
 */
var countWords = function(words1, words2) {
    let word1Obj = {}
    let word2Obj = {}
    let result = 0   
    for(let i of words1){
        word1Obj[i] = (word1Obj[i] || 0) + 1
    }
    
    for(let i of words2){
        word2Obj[i] = (word2Obj[i] || 0) + 1
    }
   
    if(words1.length > words2.length){
        for(let i in word1Obj){       
            if(word1Obj[i] === 1 && word2Obj[i] === 1) {result++}
        }
    } else {
        for(let i in word2Obj){
            if(word1Obj[i] === 1 && word2Obj[i] === 1) {result++}
        }
    }
    
    return result
};