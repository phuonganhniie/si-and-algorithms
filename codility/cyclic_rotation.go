package codility

func cyclicRotation(A []int, K int) []int {
	n := len(A)

	if n == 0 || K == 0 {
		return A
	}

	K = K % n

	rs := A[n-K:]
	remain := A[:n-K]
	rs = append(rs, remain...)
	return rs
}
