function kInversePairs(n: number, k: number): number {
  const MOD: number = 1_000_000_007;
  let dp: number[][] = Array.from({ length: n + 1 }, () =>
    Array(k + 1).fill(0)
  );

  dp[0][0] = 1;

  for (let i = 1; i <= n; i++) {
    for (let j = 0; j <= k; j++) {
      let val: number = 0;
      for (let x = 0; x <= Math.min(j, i - 1); x++) {
        val += dp[i - 1][j - x];
        val %= MOD;
      }
      dp[i][j] = val;
    }
  }
  return dp[n][k];
}
