/**
 * @param {number} num
 * @return {number}
 */
var smallestNumber = function(num) {
    let strNum = num.toString()
    let strArr = []
    let isNegative = false
    let result = ""
    if(strNum[0] === "-"){
        isNegative = true;
        for(let i = 1; i < strNum.length; i++){
            strArr.push(strNum[i])
        }
    } else {
          for(let i = 0; i < strNum.length; i++){
            strArr.push(strNum[i])
        }
    }
    if(isNegative){
        strArr.sort((a,b)=> b-a)
        result+="-"
        for(let i of strArr){
            result+=i.toString()
        }
    } else {
        strArr.sort((a,b)=> a-b)
        if(strArr[0] === "0" && strArr.length > 1){
            let firstPositiveIntIndex = strArr.findIndex(i => parseInt(i) > 0)
            strArr[0] = strArr[firstPositiveIntIndex]
            strArr[firstPositiveIntIndex] = "0"
        }

        for(let i of strArr){
            result+=i.toString()
        }
    }
   
    return parseInt(result)
};
