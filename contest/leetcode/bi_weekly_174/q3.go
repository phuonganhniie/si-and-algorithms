package biweekly174

const MOD = 1_000_000_007

func countPartitions(nums []int, target1 int, target2 int) int {
	n := len(nums)

	// prefix[i] = XOR of nums[0..i-1]
	prefix := make([]int, n+1)
	prefix[0] = 0
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] ^ nums[i]
	}

	// sum_t1[p] = sum of ways ending at positions with prefix p, needing target1 next
	// sum_t2[p] = sum of ways ending at positions with prefix p, needing target2 next
	sumT1 := make(map[int]int)
	sumT2 := make(map[int]int)

	// Initially at position 0, we need target1 (no blocks yet)
	sumT1[0] = 1

	// needT1[j] = ways to partition [0..j-1] such that next block needs XOR = target1
	// needT2[j] = ways to partition [0..j-1] such that next block needs XOR = target2
	var needT1, needT2 int

	for j := 1; j <= n; j++ {
		p := prefix[j]

		// To end a block with XOR = target1, we need prefix[i] where prefix[j] ^ prefix[i] = target1
		// So prefix[i] = prefix[j] ^ target1
		needT2 = sumT1[p^target1]

		// To end a block with XOR = target2, we need prefix[i] where prefix[j] ^ prefix[i] = target2
		// So prefix[i] = prefix[j] ^ target2
		needT1 = sumT2[p^target2]

		// Update sums for future positions
		sumT1[p] = (sumT1[p] + needT1) % MOD
		sumT2[p] = (sumT2[p] + needT2) % MOD
	}

	// Answer: partitions ending with target1 block (needT2) + partitions ending with target2 block (needT1)
	return (needT1 + needT2) % MOD
}
