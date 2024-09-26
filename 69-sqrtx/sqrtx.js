/**
 * @param {number} x
 * @return {number}
 */
var mySqrt = function(x) {
    let result
    //First we handle edge case where x = 0 
    if(x === 0) return 0

    //We setup the 2 tailend range that would contain our potential root 
    //Incrementally guess the root by doing binary search using the averageMean of 2 end averageMean
    //If square of averageMean is larger than x, we can lower the right end, vice versa till we inch closer to the true root
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