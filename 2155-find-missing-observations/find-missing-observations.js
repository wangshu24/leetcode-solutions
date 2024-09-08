/**
 * @param {number[]} rolls
 * @param {number} mean
 * @param {number} n
 * @return {number[]}
 */
var missingRolls = function (rolls, mean, n) {
  let result = [];
  let sumRemain =
    (rolls.length + n) * mean -
    rolls.reduce((accumulator, currentValue) => accumulator + currentValue, 0);
  let averageRemainRoll = sumRemain / n;
  if (averageRemainRoll > 6  || averageRemainRoll < 1) {
    return [];
  }

  if (sumRemain % n === 0) {
    for (let i = 0; i < n; i++) {
      result.push(sumRemain / n);
    }
  } else {
    let partialSumRemain = 0;
    for (let i = 0; i < n - 1; i++) {
      result.push(Math.floor(sumRemain / n));
      partialSumRemain += Math.floor(sumRemain / n)
    }
    result.push(sumRemain - partialSumRemain);
  }
  if(result[result.length-1] > 6){
    const largeValue = result.pop();
    result.push(Math.floor(sumRemain / n));
    let sumReduced = result.reduce((accumulator, currentValue) => accumulator + currentValue, 0);
    let distributor = sumRemain - sumReduced;
    console.log(distributor)
   for(let i = 0; i < distributor; i++){
    if(result[i] < 6) { 
        result[i] = result[i] + 1;
        } 
   }
  }
  return result;
};