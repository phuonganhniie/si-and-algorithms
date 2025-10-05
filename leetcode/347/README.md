# 347. Top K Frequent Elements

## 📋 Đề bài

Cho một mảng số nguyên `nums` và một số nguyên `k`, trả về `k` phần tử xuất hiện nhiều nhất trong mảng. Kết quả có thể được trả về theo bất kỳ thứ tự nào.

**Ví dụ 1:**
```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

**Ví dụ 2:**
```
Input: nums = [1], k = 1
Output: [1]
```

**Constraints:**
- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `k` luôn hợp lệ: `1 <= k <= số lượng phần tử unique`
- Độ phức tạp thời gian phải tốt hơn O(n log n)

---

## 💡 Ý tưởng giải quyết

### Approach: Bucket Sort (Distribution Sort)

**Tư duy:**
1. **Đếm tần suất** của mỗi phần tử → HashMap
2. **Phân phối** các phần tử vào bucket theo tần suất → Bucket Sort
3. **Lấy K phần tử** từ bucket có tần suất cao nhất

**Tại sao không dùng Sort thông thường?**
- Sort thông thường: O(n log n) ❌
- Bucket Sort: O(n) ✅ - Nhanh hơn!

---

## 🔍 Chi tiết thuật toán

### Bước 1: Đếm tần suất (Frequency Count)
```go
countMap := make(map[int]int)
for _, num := range nums {
    countMap[num]++
}
```

**Ví dụ:**
```
nums = [1,1,1,2,2,3]
→ countMap = {1:3, 2:2, 3:1}
```

---

### Bước 2: Bucket Sort - Phân phối theo tần suất
```go
bucket := make([][]int, len(nums)+1)
for num, count := range countMap {
    bucket[count] = append(bucket[count], num)
}
```

**Giải thích:**
- `bucket[i]` chứa các số xuất hiện `i` lần
- Kích thước bucket = `len(nums)+1` vì tần suất tối đa là `n`

**Hình dung:**
```
bucket[0] = []       ← Không có số nào xuất hiện 0 lần
bucket[1] = [3]      ← Số 3 xuất hiện 1 lần
bucket[2] = [2]      ← Số 2 xuất hiện 2 lần
bucket[3] = [1]      ← Số 1 xuất hiện 3 lần
bucket[4] = []
...
```

**✨ Magic:** Bucket tự động sắp xếp theo tần suất mà không cần compare!

---

### Bước 3: Lấy K phần tử từ cuối về đầu
```go
result := make([]int, 0, k)
for i := len(bucket) - 1; i >= 0 && len(result) < k; i-- {
    if len(bucket[i]) == 0 {
        continue
    }
    
    for _, num := range bucket[i] {
        result = append(result, num)
        if len(result) == k {
            break
        }
    }
}
```

**Giải thích:**
- Duyệt từ `bucket[n]` → `bucket[0]` (tần suất cao → thấp)
- Lấy tất cả số trong bucket cho đến khi đủ K phần tử
- Điều kiện dừng: `len(result) == k`

---

## 📊 Độ phức tạp

| Thành phần | Time Complexity | Space Complexity |
|------------|----------------|------------------|
| **Đếm tần suất** | O(n) | O(n) |
| **Bucket Sort** | O(n) | O(n) |
| **Lấy K phần tử** | O(k) | O(k) |
| **Tổng cộng** | **O(n)** | **O(n)** |

**So sánh với Sort thông thường:**
- Dùng Sort: O(n log n) ❌
- Dùng Bucket Sort: O(n) ✅ - Nhanh hơn!

---

## 🎯 Các case cần lưu ý

### Case 1: Bucket rỗng
```go
if len(bucket[i]) == 0 {
    continue  // Bỏ qua bucket không có phần tử
}
```

### Case 2: Nhiều số cùng tần suất
```go
// Bucket có thể chứa nhiều số
bucket[2] = [5, 7, 9]  // 3 số đều xuất hiện 2 lần
```

### Case 3: Đủ K phần tử trước khi hết bucket
```go
if len(result) == k {
    break  // Dừng ngay khi đủ K
}
```

---

## 🧪 Test Cases

### Test 1: Basic case
```go
Input:  nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Giải thích: Số 1 xuất hiện 3 lần, số 2 xuất hiện 2 lần
```

### Test 2: Single element
```go
Input:  nums = [1], k = 1
Output: [1]
```

### Test 3: All same frequency
```go
Input:  nums = [1,2,3,4], k = 2
Output: [1,2] hoặc bất kỳ 2 số nào
Giải thích: Tất cả đều xuất hiện 1 lần
```

### Test 4: Negative numbers
```go
Input:  nums = [-1,-1,-1,2,2,3], k = 2
Output: [-1,2]
```

---

## 🔄 Cách tiếp cận khác

### Approach 2: Min Heap
```go
// Sử dụng Min Heap với size = k
// Time: O(n log k)
// Space: O(n)
```

**Ưu điểm:**
- Không cần biết range của tần suất
- Tiết kiệm bộ nhớ khi k nhỏ

**Nhược điểm:**
- Chậm hơn Bucket Sort: O(n log k) vs O(n)

---

### Approach 3: Quick Select
```go
// Sử dụng QuickSelect để tìm K phần tử lớn nhất
// Time: O(n) average, O(n²) worst case
// Space: O(n)
```

---

## 📚 Bài tập tương tự

1. **LeetCode 692** - Top K Frequent Words
2. **LeetCode 451** - Sort Characters By Frequency  
3. **LeetCode 973** - K Closest Points to Origin
4. **LeetCode 215** - Kth Largest Element in an Array

**Pattern chung:** Đếm → Sắp xếp thông minh → Lấy K phần tử

---

## 💻 Cách chạy test

```bash
# Chạy test
go test -v

# Chạy test với coverage
go test -v -cover

# Chạy test cụ thể
go test -v -run TestTopKFrequent
```

---

## 🎓 Takeaways

✅ **Bucket Sort** là dạng Distribution Sort - phân phối thay vì so sánh  
✅ Độ phức tạp **O(n)** - nhanh hơn sort thông thường O(n log n)  
✅ Phù hợp khi **range nhỏ** và **biết trước** (ở đây là 0 → n)  
✅ Pattern "Top K" thường dùng: **Counting + Bucket/Heap**  
✅ Tư duy: **Đếm → Phân phối → Lấy K**

---

## 🔗 Links

- [LeetCode Problem](https://leetcode.com/problems/top-k-frequent-elements/)
- [Bucket Sort Explained](https://en.wikipedia.org/wiki/Bucket_sort)
- [Distribution Sort](https://en.wikipedia.org/wiki/Sorting_algorithm#Distribution_sort)
