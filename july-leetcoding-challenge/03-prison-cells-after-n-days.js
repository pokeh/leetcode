"use strict";

/**
 * @param {number[]} cells
 * @param {number} N
 * @return {number[]}
 */
// TLE
var prisonAfterNDays = function (cells, N) {
  const res = Array(cells.length).fill(0);

  for (let i = 0; i < N; i++) {
    for (let j = 1; j < cells.length - 1; j++) {
      res[j] = cells[j - 1] === cells[j + 1] ? 1 : 0;
    }

    cells = [...res];
  }

  return res;
};

console.log(prisonAfterNDays([0, 1, 1, 0], 3));
