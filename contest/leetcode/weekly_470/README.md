# No Zero Pairs - LeetCode Weekly Contest 470

## Đề bài

Một số nguyên no-zero là số nguyên dương không chứa chữ số 0 trong biểu diễn thập phân của nó.

Cho một số nguyên `n`, đếm số lượng cặp `(a, b)` sao cho:
- `a` và `b` là các số nguyên no-zero
- `a + b = n`

Trả về số nguyên biểu thị số lượng các cặp như vậy.

### Ví dụ

**Example 1:**
```
Input: n = 2
Output: 1
Explanation: Cặp duy nhất là (1, 1)
```

**Example 2:**
```
Input: n = 3
Output: 2
Explanation: Các cặp là (1, 2) và (2, 1)
```

**Example 3:**
```
Input: n = 11
Output: 8
Explanation: Các cặp là (2, 9), (3, 8), (4, 7), (5, 6), (6, 5), (7, 4), (8, 3), và (9, 2)
Lưu ý: (1, 10) và (10, 1) không thỏa mãn vì 10 chứa chữ số 0
```

**Large Test Case:**
```
Input: n = 34201859
Output: 7587520
```

### Constraints

- `2 <= n <= 10^15`

## Phương pháp giải

### Thuật toán: Digit Dynamic Programming

Thay vì duyệt trực tiếp hoặc sinh tất cả các số không chứa 0, ta sử dụng **Digit DP** để đếm trực tiếp số cặp hợp lệ.

#### Ý tưởng chính:

1. **Xử lý từng chữ số**: Xét n theo từng vị trí chữ số từ phải sang trái
2. **Thử các khả năng**: Với mỗi vị trí, thử tất cả các cặp chữ số (da, db) từ 1-9
3. **Kiểm tra tổng**: da + db + carry phải cho chữ số đúng của n
4. **Memoization**: Lưu cache để tránh tính lại các trạng thái đã xét

#### Trạng thái DP:

```
solve(pos, carry, lenA, lenB)
```

- `pos`: Vị trí chữ số đang xét (0 = chữ số cuối cùng)
- `carry`: Số nhớ từ phép cộng trước
- `lenA`: Số chữ số của số a
- `lenB`: Số chữ số của số b

#### Điều kiện:

- Nếu `pos < lenA`: chữ số của a ∈ {1,2,3,4,5,6,7,8,9}
- Nếu `pos >= lenA`: chữ số của a = 0 (leading zero)
- Tương tự cho b

### Độ phức tạp

- **Thời gian**: O(k² × 81 × k) với k là số chữ số của n
  - k² cho tất cả các cặp (lenA, lenB)
  - 81 = 9×9 cho các cặp chữ số (da, db)
  - k cho số vị trí
  - Với memoization, mỗi trạng thái chỉ tính 1 lần
- **Không gian**: O(k⁴) cho cache
- **Thực tế với n = 10^15**: ~0.28ms

### Các hàm chính

1. **`CountNoZeroPairs(n int64) int64`**: Hàm chính sử dụng Digit DP
2. **`CountNoZeroPairsOld(n int64) int64`**: Implementation cũ (giữ lại để tham khảo)
3. **`hasZero(num int64) bool`**: Kiểm tra xem số có chứa chữ số 0
4. **`generateNoZeroNumbers(maxNum int64) []int64`**: Sinh tất cả số không chứa 0

## Chạy code

### Chạy tests

```bash
cd contest/leetcode/weekly_470
go test -v
```

### Chạy tests với test case cụ thể

```bash
go test -v -run TestCountNoZeroPairs/Large_test_case
```

### Chạy benchmark

```bash
go test -bench=. -benchmem
```

## Kết quả test

Tất cả test cases đều pass:
- Example 1: n = 2 → 1 cặp ✓
- Example 2: n = 3 → 2 cặp ✓
- Example 3: n = 11 → 8 cặp ✓
- n = 100 → 90 cặp ✓
- n = 1000 → 738 cặp ✓
- **n = 34201859 → 7587520 cặp ✓** (chạy gần như tức thì!)

## Performance Benchmark

```
BenchmarkCountNoZeroPairs/n=11-8           325480    3,373 ns/op     304 B/op    37 allocs/op
BenchmarkCountNoZeroPairs/n=100-8           90462   13,316 ns/op   2,424 B/op   133 allocs/op
BenchmarkCountNoZeroPairs/n=1000-8          33020   32,692 ns/op   8,960 B/op   289 allocs/op
BenchmarkCountNoZeroPairs/n=10000-8         21476   55,920 ns/op  10,864 B/op   525 allocs/op
BenchmarkCountNoZeroPairsLarge/n=34201859-8  4245  278,466 ns/op  72,792 B/op 2,338 allocs/op
```

**So sánh với phương pháp cũ:**
- Phương pháp cũ (sinh tất cả số): ~600ms cho n = 34201859
- **Phương pháp mới (Digit DP): ~0.28ms** cho n = 34201859
- **Nhanh hơn hơn 2000 lần!** 🚀

### Ưu điểm của Digit DP:

1. **Hiệu quả với số lớn**: Không phụ thuộc vào giá trị của n, chỉ phụ thuộc vào số chữ số
2. **Memory efficient**: Chỉ ~73KB cho n = 34201859
3. **Memoization**: Tránh tính lại các trạng thái đã xét
4. **Scalable**: Có thể xử lý n lên đến 10^15

## So sánh các phương pháp

| Phương pháp | Thời gian (n=34201859) | Memory | Độ phức tạp |
|-------------|------------------------|---------|-------------|
| Brute Force | TLE (>1s) | O(n) | O(n) |
| Generate & Check | ~600ms | ~4MB | O(9^k) |
| **Digit DP** | **~0.28ms** | **~73KB** | **O(k³)** |

## Code Python tương đương

Solution tham khao tu: `https://leetcode.com/problems/count-no-zero-pairs-that-sum-to-n/solutions/7249582/easy-to-understand-solution/`

```python
from functools import lru_cache

class Solution:
    def countNoZeroPairs(self, n: int) -> int:
        lenN = len(str(n))
        digits = list(map(int, str(n)))[::-1]

        @lru_cache(None)
        def solve(pos, carry, lenA, lenB):
            if pos == lenN:
                return int(carry == 0)

            rangeA = range(1, 10) if pos < lenA else (0,)
            rangeB = range(1, 10) if pos < lenB else (0,)

            numWays = 0
            for da in rangeA:
                for db in rangeB:
                    summ = da + db + carry
                    if summ % 10 == digits[pos]:
                        numWays += solve(pos + 1, summ // 10, lenA, lenB)
            return numWays

        totalPairs = 0
        for lenA in range(1, lenN + 1):
            for lenB in range(1, lenN + 1):
                totalPairs += solve(0, 0, lenA, lenB)

        return totalPairs
```

## Kết luận

Bài toán này là một ví dụ điển hình về **Digit Dynamic Programming**, một kỹ thuật rất hữu ích cho các bài toán:
- Đếm số có các tính chất đặc biệt trong một khoảng
- Xử lý các ràng buộc về chữ số
- Làm việc với số rất lớn (lên đến 10^18)

Kỹ thuật này thường xuất hiện trong:
- LeetCode Hard problems
- Competitive Programming (Codeforces, AtCoder)
- Google Code Jam, Facebook Hacker Cup

