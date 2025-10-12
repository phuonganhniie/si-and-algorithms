package weekly471

func sumDivisibleByK(nums []int, k int) int {
    m := make(map[int]int)
    for _, num := range nums {
		m[num]++
	} 

	sum := 0
	for k, v := range m {
		if v%k==0 {
			sum += k*v
		}
	}
	return sum
}