"use strict";

var prisonAfterNDays = function (cells, N) {
  const res = Array(cells.length).fill(0);

  const history = [];

  for (let i = 1; i <= N; i++) {
    for (let j = 1; j < cells.length - 1; j++) {
      res[j] = cells[j - 1] === cells[j + 1] ? 1 : 0;
    }

    if (history.includes(res.join(""))) {
      const idx = N % history.length;

      if (idx === 0) {
        return history[history.length - 1].split("");
      }

      return history[idx - 1].split("");
    }

    history.push(res.join(""));

    cells = [...res];
  }

  return res;
};

// TLE
/*
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
*/

console.log(prisonAfterNDays([0, 1, 1, 0], 3));
