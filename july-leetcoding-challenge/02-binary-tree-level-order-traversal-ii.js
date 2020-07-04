/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
/*
var levelOrderBottom = function (root) {
  let res = [];
  dfs(root, 0);
  return res.reverse();

  function dfs(node, depth) {
    if (!node) {
      return;
    }

    if (!res[depth]) {
      res[depth] = [];
    }
    res[depth].push(node.val);

    dfs(node.left, depth + 1);
    dfs(node.right, depth + 1);
  }
};
*/

/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var levelOrderBottom = function (root) {
  const res = [];

  const dfs = (node, depth) => {
    if (!node) {
      return;
    }

    if (!res[depth]) {
      res[depth] = [];
    }
    res[depth].push(node.val);

    dfs(node.left, depth + 1);
    dfs(node.right, depth + 1);
  };

  dfs(root, 0);

  return res.reverse();
};
