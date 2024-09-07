/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  let indices = [];
  for (let i = 0; i < nums.length; i++) {
    let newNums = nums.toSpliced(i, 1);
    newNums.forEach(n => {
      if (n + nums[i] === target) {
        !indices.includes(i) && indices.push(i);
        !indices.includes(nums.indexOf(n)) && indices.push(nums.indexOf(n));
      }
    });
  }
  return indices;
};

console.log(twoSum([3, 2, 4], 6));
