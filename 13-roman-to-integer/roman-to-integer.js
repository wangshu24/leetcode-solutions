/**
 * @param {string} s
 * @return {number}
 */
var romanToInt = function(s) {
    const values = [{value : 1000, letter: "M"}, {value : 500, letter: "D"},{value: 100, letter : "C"}, {value : 50, letter : "L"}, {value : 10, letter : "X"} , {value : 5, letter: "V"}, {value : 1, letter : "I"}]
    let valueArray = [];
    for(let i = 0; i < s.length; i++){
        let extracted = values.filter(num => num.letter === s[i]);
        valueArray.push(extracted[0].value);
    }
    valueArray.forEach((num, index) => {    
        if(valueArray[index+1]){
          if(num < valueArray[index+1]){valueArray[index] = -num}
        }
    })
    let total = valueArray.reduce(
  (accumulator, currentValue) => accumulator + currentValue,
  0,
);
return total;
};