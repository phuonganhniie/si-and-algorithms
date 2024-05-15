package leetcode

func expandAroundCenter(s string, left, right int) int {
	// Mở rộng khi các ký tự bằng nhau và trong giới hạn
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	// Trả về độ dài của chuỗi con đối xứng
	return right - left - 1
}

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0

	// Duyệt qua mỗi ký tự trong chuỗi
	for i := 0; i < len(s); i++ {
		// Kiểm tra chuỗi con đối xứng lẻ (trung tâm là một ký tự)
		odd := expandAroundCenter(s, i, i)
		// Kiểm tra chuỗi con đối xứng chẵn (trung tâm là hai ký tự)
		even := expandAroundCenter(s, i, i+1)

		// Tìm độ dài lớn nhất giữa hai chuỗi con đối xứng
		maxSubsLen := max(odd, even)

		// Cập nhật chỉ số bắt đầu và kết thúc nếu tìm thấy chuỗi dài hơn
		if maxSubsLen > end-start {
			start = i - (maxSubsLen-1)/2
			end = i + maxSubsLen/2
		}
	}

	// Trả về chuỗi con đối xứng dài nhất
	return s[start : end+1]
}
