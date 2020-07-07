"use strict";

/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function (digits) {
  let idx = digits.length - 1;

  while (digits[idx] === 9) {
    digits[idx] = 0;
    idx -= 1;
  }

  if (idx < 0) {
    digits.unshift(1);
  } else {
    digits[idx]++;
  }

  return digits;
};

console.log(plusOne([1, 2, 3]));
console.log(plusOne([2, 9, 9]));
console.log(plusOne([0]));
console.log(plusOne([9]));
