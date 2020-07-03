/*
var arrangeCoins = function (n) {
  for (i = 1; ; i++) {
    n -= i;
    if (n < 0) {
      return i - 1;
    }
  }
};
*/

/**
 * @param {number} n
 * @return {number}
 */
var arrangeCoins = function (n) {
  let cnt = 1;
  while (n >= 0) {
    n -= cnt;
    cnt++;
  }
  return cnt - 2;
};
