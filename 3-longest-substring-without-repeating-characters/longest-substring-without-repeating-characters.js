/**
 * @param {string} s
 * @return {number}
 */
function compareFn(a, b) {
  if (a.length > b.length) {
    return -1;
  } else if (a.length <  b.length) {
    return 1;
  }
  // a must be equal to b
  return 0;
}

var lengthOfLongestSubstring = function(s) {
    let final = "";
    let stringArr = [];
    for(let i = 0; i < s.length; i++){
        stringArr.push(s[i])
    }
    let streakArr = [];
    let finalStreak = stringArr.reduce((accumulator, currentValue) => {
  if (accumulator.includes(currentValue)) {
    streakArr.push(accumulator);
    if(accumulator.substring(accumulator.length-1).includes(currentValue)) {return currentValue}
    return accumulator.includes(currentValue) ? accumulator.substring(accumulator.indexOf(currentValue)+1) + currentValue : accumulator.substring(accumulator.length-1) + currentValue;
  } else {
    return accumulator + currentValue;
  }
}, '');
    streakArr.push(finalStreak);
    console.log(streakArr)
    final = streakArr.sort(compareFn)[0]

    return final.length;
};