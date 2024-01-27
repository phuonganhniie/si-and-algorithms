/**
 * Calculates the number of different arrays consist of numbers from 1 to n
 * such that there are exactly k inverse pairs.
 * @param n The number of elements in the array.
 * @param k The exact number of inverse pairs.
 * @returns The number of different arrays modulo 1_000_000_007.
 */
function kInversePairs(n: number, k: number): number {
  const MOD: number = 1_000_000_007;
  let dp: number[][] = Array.from({ length: n + 1 }, () =>
    Array(k + 1).fill(0)
  );

  // Base case: one way to form an empty array with 0 inverse pairs
  dp[0][0] = 1;

  // Iterate over array sizes
  for (let i = 1; i <= n; i++) {
    // Iterate over possible inverse pair counts
    for (let j = 0; j <= k; j++) {
      let val: number = 0;

      // Sum over possible positions for the ith element
      for (let x = 0; x <= Math.min(j, i - 1); x++) {
        val += dp[i - 1][j - x];
        val %= MOD;
      }

      dp[i][j] = val;
      console.log(`dp[${i}][${j}] = ${val}`); // Debugging: log current DP value
    }
  }
  return dp[n][k];
}
