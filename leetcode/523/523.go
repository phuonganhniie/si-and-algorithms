package leetcode_523

/*
[Medium] 523. Continuous Subarray Sum
https://leetcode.com/problems/continuous-subarray-sum/description/
Created: 2024-06-10
Done   : 30 mins
Attempt: 4
---------------------NOTE---------------------
Time: O(n)
Space: O(n)
Approach: Phương pháp này dựa trên nguyên lý:
Nếu phần dư của tổng tích lũy gặp lại cùng một giá trị thì nghĩa là có một dãy con mà tổng các phần tử trong đó chia hết cho `k`.
Dựa vào điều này, ta chỉ cần duyệt qua mảng một lần, lưu trữ phần dư và chỉ số, kiểm tra các điều kiện cần thiết để xác định xem có tồn tại dãy con thỏa mãn yêu cầu hay không.
*/
func checkSubarraySum(nums []int, k int) bool {
	remainderMap := make(map[int]int, len(nums))
	remainderMap[0] = -1

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		remainder := sum % k

		if v, exist := remainderMap[remainder]; exist {
			if i-v >= 2 {
				return true
			}
			continue
		}
		remainderMap[remainder] = i
	}
	return false
}
