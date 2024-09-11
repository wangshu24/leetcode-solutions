/**
 * @param {number} num
 * @return {string}
 */

var calcIntDigitToRomanDigit = function(num){
    if(num<9){
        if(num>3) {
            if(num%5===0 || num/5 <=1){
                if(num===5){
                    return[5]
                } else {return [1,5]}
            } else {
                let result = [5]
                for(let i = 0; i < num%5; i++){
                    result.push(1);
                }
                return result;
            }
        } else {
            let result = []
            for(let i = 0; i < num; i++){
                    result.push(1);
                }
            return result
        }
    }else{ return [1,10]}
};

var intToRoman = function(num) {
    let digits = num.toString().length;
    const divider = [1000,100,10,1]
    const romanDict = {
        "1" :    "I",	
        "5" :  "V",	
        "10": "X",	
        "50" :  "L",	  
        "100" :  "C",	
        "500" : "D",	
        "1000" : "M"	
    }
    const starter = divider.length - digits;
    let removed = 0;
    let roman = "";
    console.log
    for(let i =starter; i < 4; i++){
        let intInPlace = Math.floor((num-removed)/divider[i]);
        if(i===0) {
            for(let i = 0; i < intInPlace; i++){
                roman+="M"
            }
            removed += intInPlace * divider[i];
            }
        else{ 
            let toConvert = calcIntDigitToRomanDigit(intInPlace);
            toConvert.forEach(e => 
            {
                let intValue = e*divider[i]
                // console.log("at ", i," times: value is ", e)
                roman+=romanDict[intValue.toString()]
                })
                removed += intInPlace * divider[i];
        } 
    }
    return roman;
};

