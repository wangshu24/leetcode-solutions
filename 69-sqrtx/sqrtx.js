/**
 * @param {number} x
 * @return {number}
 */
var mySqrt = function(x) {
    let result
    if(x === 0) return 0
    left = 0
    right = x
    while(left <= right){
      mid = Math.floor((left + right)/2)
      if(mid*mid > x){
        right = mid -1
      } else {
        //include case where mid*mid === x
        result = mid
        left = mid + 1
      }
    }
    return result
};