"use strict";

/**
 * @param {number} x
 * @param {number} y
 * @return {number}
 */
var hammingDistance = function (x, y) {
  let xor = x ^ y;
  let count = 0;

  while (xor > 0) {
    // count number of 1's in the XORed value
    if (xor & 1) {
      count++;
    }
    xor >>= 1;
  }

  return count;
};

console.log(hammingDistance(1, 4));
