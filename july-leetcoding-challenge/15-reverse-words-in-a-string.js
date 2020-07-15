/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function (s) {
  return s
    .split(" ")
    .reverse()
    .filter((el) => el)
    .join(" ");
};
