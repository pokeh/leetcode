"use strict";

/*
// Brute force: TLE
var threeSum = function (nums) {
  let res = [[]];

  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      for (let k = j + 1; k < nums.length; k++) {
        if (nums[i] + nums[j] + nums[k] == 0) {
          res.push([nums[i], nums[j], nums[k]]);
        }
      }
    }
  }
  res.shift(); // take out the initial empty array

  res.map((arr) => arr.sort());

  // get unique set
  res = new Set(res.map(JSON.stringify));
  res = Array.from(res).map(JSON.parse);

  return res;
};
*/

/*
// Still TLE
var threeSum = function (nums) {
  let res = new Set();

  let numCounts = new Map();
  for (let i = 0; i < nums.length; i++) {
    numCounts[nums[i]] = numCounts[nums[i]] ? numCounts[nums[i]] + 1 : 1;
  }

  var isResult = function (i, j) {
    const third = -(nums[i] + nums[j]);
    const thirdCnt = numCounts[third];

    if (nums[i] == nums[j]) {
      return thirdCnt > 2;
    } else {
      if (third == nums[i] || third == nums[j]) {
        return thirdCnt > 1;
      } else {
        return thirdCnt > 0;
      }
    }
  };

  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (isResult(i, j)) {
        res.add(
          JSON.stringify([nums[i], nums[j], -(nums[i] + nums[j])].sort())
        );
      }
    }
  }

  return Array.from(res).map(JSON.parse);
};
*/

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function (nums) {
  if (nums.filter((i) => i != 0).length == 0) {
    return nums.length >= 3 ? [[0, 0, 0]] : [];
  }

  let res = new Set();

  // sort from smallest to largest
  nums.sort((a, b) => a - b);

  for (let i = 0; i < nums.length; i++) {
    let j = i + 1;
    let k = nums.length - 1;

    while (j < k) {
      const sum = nums[i] + nums[j] + nums[k];
      if (sum == 0) {
        res.add(JSON.stringify([nums[i], nums[j], nums[k]]));
        j++;
        k--;
      } else if (sum > 0) {
        k--;
      } else {
        j++;
      }
    }
  }

  return Array.from(res).map(JSON.parse);
};

console.log(threeSum([-1, 0, 1, 2, -1, -4]));
