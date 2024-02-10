/**
 * 647. Palindromic Substrings
 * https://leetcode.com/problems/palindromic-substrings/description/
 */

function countSubstrings(s: string): number {
  const n: number = s.length;
  let dp: boolean[][] = Array.from({ length: n }, () => Array(n).fill(false));
  let count: number = 0;

  for (let i = 0; i < n; i++) {
    dp[i][i] = true;
    count++;
  }

  for (let i = 0; i < n - 1; i++) {
    if (s.charAt(i) === s.charAt(i + 1)) {
      dp[i][i + 1] = true;
      count++;
    }
  }

  for (let l = 3; l <= n; l++) {
    for (let i = 0; i < n - l + 1; i++) {
      let j = i + l - 1;
      if (s.charAt(i) === s.charAt(j) && dp[i + 1][j - 1]) {
        dp[i][j] = true;
        count++;
      }
    }
  }

  return count;
}
