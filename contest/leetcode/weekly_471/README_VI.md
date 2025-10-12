# LeetCode Weekly Contest 471 - Lời Giải

Thư mục này chứa lời giải cho tất cả các bài trong LeetCode Weekly Contest 471.

## Tổng Quan Các Bài

| # | Bài Toán | Độ Khó | Trạng Thái |
|---|----------|---------|------------|
| Q1 | Tổng Chia Hết Cho K | Dễ | ✅ Hoàn Thành |
| Q2 | Chuỗi Con Cân Bằng Dài Nhất | Trung Bình | ✅ Hoàn Thành |
| Q3 | Chuỗi Con Cân Bằng Dài Nhất (chỉ a,b,c) | Trung Bình | ✅ Hoàn Thành |
| Q4 | Tổng Các Cặp Tổ Tiên (Số Chính Phương) | Khó | ✅ Hoàn Thành |

---

## Q1: Tổng Chia Hết Cho K

### Đề Bài
Cho mảng `nums` và số nguyên `k`, trả về tổng của tất cả các phần tử mà tần suất xuất hiện của chúng chia hết cho giá trị của chính phần tử đó.

### Ví Dụ
```
Input: nums = [1,2,2,2,3,3,3,3], k = 2
Giải thích:
- 1 xuất hiện 1 lần (1 % 1 == 0) → thêm 1 * 1 = 1
- 2 xuất hiện 3 lần (3 % 2 != 0) → không thêm
- 3 xuất hiện 4 lần (4 % 3 != 0) → không thêm
```

### Phương Pháp Giải

**Thuật Toán**: Đếm Tần Suất Bằng Hash Map
- Sử dụng hash map để đếm tần suất xuất hiện của mỗi số
- Với mỗi số duy nhất, kiểm tra xem tần suất có chia hết cho chính số đó không
- Nếu có, cộng `số × tần_suất` vào tổng

**Độ Phức Tạp Thời Gian**: O(n) - Duyệt mảng một lần + duyệt các giá trị duy nhất
**Độ Phức Tạp Không Gian**: O(n) - Lưu trữ hash map

### Điểm Chú Ý
- Điều kiện `tần_suất % giá_trị == 0` kiểm tra xem tần suất có chia hết cho chính giá trị đó, không phải cho k
- Đây là bài toán đếm đơn giản với một yêu cầu đặc biệt

---

## Q2: Chuỗi Con Cân Bằng Dài Nhất

### Đề Bài
Tìm chuỗi con dài nhất mà tất cả các ký tự phân biệt xuất hiện với cùng một tần suất.

### Ví Dụ
```
Input: s = "aabbc"
Output: 4
Giải thích: "aabb" có 'a' xuất hiện 2 lần, 'b' xuất hiện 2 lần → cân bằng!
```

### Phương Pháp Giải

**Thuật Toán**: Brute Force Với Kiểm Tra Tần Suất
1. Thử tất cả các chuỗi con có thể bằng hai vòng lặp lồng nhau
2. Với mỗi chuỗi con, duy trì một map tần suất
3. Kiểm tra xem tất cả ký tự có cùng tần suất không
4. Theo dõi độ dài lớn nhất tìm được

**Độ Phức Tạp Thời Gian**: O(n²) để sinh chuỗi con, O(kích_thước_bảng_chữ_cái) để kiểm tra → O(n²)
**Độ Phức Tạp Không Gian**: O(kích_thước_bảng_chữ_cái) cho map tần suất

### Điểm Chú Ý
- Chuỗi cân bằng nghĩa là TẤT CẢ các ký tự phân biệt phải xuất hiện cùng số lần
- Phương pháp brute force hoạt động tốt vì:
  - Ta cần kiểm tra tất cả chuỗi con
  - Kích thước bảng chữ cái thường nhỏ (26 chữ cái tiếng Anh thường)
- Với bài này, độ phức tạp O(n²) chấp nhận được với ràng buộc điển hình

---

## Q3: Chuỗi Con Cân Bằng Dài Nhất (Chỉ 'a', 'b', 'c')

### Đề Bài
Tìm chuỗi con dài nhất mà tất cả các ký tự phân biệt (chỉ từ 'a', 'b', 'c') xuất hiện với cùng tần suất.

### Ví Dụ
```
Input: s = "abbac"
Output: 4
Giải thích: "abba" có 'a' xuất hiện 2 lần, 'b' xuất hiện 2 lần → độ dài 4

Input: s = "accc"
Output: 3
Giải thích: "ccc" chỉ có 'c' xuất hiện 3 lần → cân bằng, độ dài 3
```

### Phương Pháp Giải

**Thuật Toán**: Hash Map Dựa Trên Trạng Thái Với Liệt Kê Bitmask

Đây là một tối ưu hóa nâng cao so với Q2, tận dụng ràng buộc chỉ có 3 ký tự.

**Đổi Mới Chính**: Thay vì kiểm tra tất cả chuỗi con, ta sử dụng phương pháp "mẫu khác biệt trạng thái":

1. **Liệt kê tất cả tập ký tự**: Thử tất cả 7 tập con có thể {a}, {b}, {c}, {ab}, {ac}, {bc}, {abc}
2. **Biểu diễn trạng thái**: Với mỗi tập con, theo dõi độ chênh lệch giữa số lượng ký tự
   - Trạng thái = (số_lượng[0] - số_lượng[1], số_lượng[1] - số_lượng[2])
   - Khi trạng thái quay về một trạng thái trước đó, các ký tự có số lượng cân bằng
3. **Xử lý ký tự không hợp lệ**: Khi gặp ký tự không trong tập con hiện tại, đặt lại việc theo dõi

**Hiểu Về Mẫu Trạng Thái**:
- Nếu ta thấy trạng thái (0, 0) tại vị trí i và j, thì từ i+1 đến j, tất cả ký tự có số lượng bằng nhau
- Ví dụ: Với "aabbc" theo dõi {a,b,c}:
  - Sau "aa": số_lượng=[2,0,0], trạng_thái=(2,0)
  - Sau "aabb": số_lượng=[2,2,0], trạng_thái=(0,2)
  - Sau "aabbc": số_lượng=[2,2,1], trạng_thái=(0,1)

**Độ Phức Tạp Thời Gian**: O(7n) = O(n) - Quét tuyến tính cho mỗi tập 7 ký tự
**Độ Phức Tạp Không Gian**: O(n) - Hash map trạng thái trong trường hợp xấu nhất

### Tại Sao Cách Này Hiệu Quả
Phương pháp dựa trên trạng thái nhanh hơn nhiều so với O(n²) vì:
1. Ta không sinh tất cả chuỗi con
2. Sự lặp lại trạng thái tự động tìm ra các đoạn cân bằng
3. Chỉ có 3 ký tự nghĩa là không gian trạng thái hạn chế

### Các Tối Ưu Hóa Được Áp Dụng
1. **Loại bỏ hàm xác minh**: Tin tưởng vào phát hiện dựa trên trạng thái
2. **Đặt lại khi gặp ký tự không hợp lệ**: Ngăn việc kéo dài qua các đoạn không hợp lệ
3. **Theo dõi vị trí bắt đầu hợp lệ**: Đảm bảo tính đúng đắn sau khi đặt lại
4. **Quét đơn cho mỗi tập con**: Không cần vòng lặp lồng nhau

---

## Q4: Tổng Các Cặp Tổ Tiên (Số Chính Phương)

### Đề Bài
Cho một cây có gốc tại node 0, với mỗi node không phải gốc i, đếm có bao nhiêu tổ tiên có tính chất `nums[i] × nums[tổ_tiên]` là số chính phương. Trả về tổng của tất cả các số đếm.

### Ví Dụ
```
Input: n = 3, edges = [[0,1],[1,2]], nums = [2,8,2]
Output: 3
Giải thích:
- Node 1: tổ tiên 0 → 8 × 2 = 16 = 4² ✓ (đếm = 1)
- Node 2: tổ tiên [1,0]
  - 2 × 8 = 16 = 4² ✓
  - 2 × 2 = 4 = 2² ✓
  - (đếm = 2)
- Tổng: 1 + 2 = 3
```

### Phương Pháp Giải

**Thuật Toán**: DFS Với Lý Thuyết Số Square-Free

**Hiểu Biết Toán Học Cốt Lõi**: 
Hai số nhân với nhau tạo thành số chính phương **khi và chỉ khi** chúng có cùng "phần square-free".

**Phần Square-Free Là Gì?**
- Phần square-free của một số là tích của tất cả các thừa số nguyên tố có số mũ lẻ
- Ví dụ:
  - 8 = 2³ → square-free = 2 (số mũ 3 là lẻ)
  - 4 = 2² → square-free = 1 (số mũ 2 là chẵn, bị loại bỏ)
  - 18 = 2¹ × 3² → square-free = 2 (chỉ 2 có số mũ lẻ)
  - 12 = 2² × 3¹ → square-free = 3

**Tại Sao Cách Này Hoạt Động?**
```
Nếu a = p1^e1 × p2^e2 × ... và b = p1^f1 × p2^f2 × ...
Thì a × b là số chính phương ⟺ tất cả (ei + fi) đều chẵn

Điều này xảy ra ⟺ ei và fi có cùng tính chẵn lẻ với mọi số nguyên tố
⟺ square-free(a) == square-free(b)
```

**Các Bước Thuật Toán**:
1. **Tính trước các giá trị square-free** cho tất cả nums[i]
2. **Xây dựng cây** dưới dạng danh sách kề
3. **Duyệt DFS** với theo dõi tổ tiên:
   - Duy trì một map các giá trị square-free dọc theo đường đi từ gốc
   - Với mỗi node, đếm các tổ tiên có giá trị square-free khớp
   - Sử dụng backtracking để quản lý map tổ tiên đúng cách

**Độ Phức Tạp Thời Gian**: 
- Tính square-free: O(n × √max(nums))
- DFS: O(n)
- **Tổng**: O(n × √max(nums))

**Độ Phức Tạp Không Gian**: O(n) cho cấu trúc cây và theo dõi tổ tiên

### Chi Tiết Cài Đặt

**Tối Ưu Hóa Tính Square-Free**:
```go
func getSquareFree(num int) int {
    result := 1
    // Xử lý 2 riêng
    count := 0
    for num % 2 == 0 {
        count++
        num /= 2
    }
    if count % 2 == 1 {
        result *= 2
    }
    
    // Chỉ kiểm tra các thừa số lẻ
    for i := 3; i*i <= num; i += 2 {
        count := 0
        for num % i == 0 {
            count++
            num /= i
        }
        if count % 2 == 1 {
            result *= i
        }
    }
    
    // Thừa số còn lại là số nguyên tố
    if num > 1 {
        result *= num
    }
    return result
}
```

**DFS Với Backtracking**:
```go
func dfs(node, parent int, ancestors map[int]int) {
    // Đếm các tổ tiên khớp
    if node != 0 {
        totalSum += ancestors[squareFree[node]]
    }
    
    // Thêm bản thân vào đường đi
    ancestors[squareFree[node]]++
    
    // Thăm các node con
    for _, child := range graph[node] {
        if child != parent {
            dfs(child, node, ancestors)
        }
    }
    
    // Backtrack (xóa bản thân khỏi đường đi)
    ancestors[squareFree[node]]--
}
```

### Điểm Chú Ý
1. **Lý thuyết số là quan trọng**: Nhận ra tính chất square-free biến đổi một bài toán có vẻ O(n²) thành O(n × √max(nums))
2. **DFS với backtracking**: Thiết yếu để theo dõi tổ tiên hiệu quả mà không cần sao chép map
3. **Hash map vs mảng**: Sử dụng map tốt hơn mảng kích thước 10⁵ vì chỉ một tập con các giá trị square-free xuất hiện
4. **Tối ưu hóa quan trọng**: Xử lý thừa số 2 riêng cắt giảm một nửa số lần lặp

### Tại Sao Đây Là Giải Pháp Tối Ưu
- Không thể tránh tính square-free: O(√num) là tối ưu cho phân tích thừa số nguyên tố
- Không thể tránh thăm mỗi node: DFS O(n) là cần thiết
- Thao tác hash map là O(1) trường hợp trung bình
- Không có công việc dư thừa được thực hiện

---

## Chạy Kiểm Thử

Để chạy tất cả các test:
```bash
cd contest/leetcode/weekly_471
go test -v
```

Để chạy bài cụ thể:
```bash
go test -v -run TestSumDivisibleByK           # Q1
go test -v -run TestLongestBalanced           # Q2  
go test -v -run TestLongestBalancedSubstring  # Q3
go test -v -run TestSumOfAncestorPairs        # Q4
```

---

## Thống Kê Contest
- **Độ Khó Tăng Dần**: Dễ → Trung Bình → Trung Bình → Khó
- **Kỹ Năng Cần Thiết**:
  - Hash map và đếm tần suất (Q1, Q2)
  - Tối ưu hóa dựa trên trạng thái (Q3)
  - Lý thuyết số và thuật toán cây (Q4)
- **Điểm Học Tập**:
  - Tối ưu hóa dựa trên ràng buộc (Q3 với ràng buộc 3 ký tự)
  - Hiểu biết toán học có thể cải thiện độ phức tạp đáng kể (Q4)
  - Backtracking cho các bài toán phụ thuộc đường đi (Q4)

---

## So Sánh Phương Pháp

### Q3: Brute Force vs State-Based
| Phương Pháp | Độ Phức Tạp | Ưu Điểm | Nhược Điểm |
|-------------|-------------|---------|-----------|
| Brute Force (Q2) | O(n²) | Đơn giản, dễ hiểu | Chậm với n lớn |
| State-Based (Q3) | O(n) | Nhanh, tối ưu | Phức tạp hơn, cần hiểu trạng thái |

### Q4: Naive vs Square-Free
| Phương Pháp | Độ Phức Tạp | Ý Tưởng |
|-------------|-------------|---------|
| Naive | O(n² × √max) | Kiểm tra mọi cặp (node, tổ tiên) |
| Square-Free | O(n × √max) | Sử dụng lý thuyết số để nhóm |

---

## Bài Học Rút Ra

1. **Ràng buộc là manh mối**: Q3 với chỉ 3 ký tự cho phép thuật toán hoàn toàn khác
2. **Toán học giúp tối ưu**: Q4 cho thấy lý thuyết số có thể giảm độ phức tạp từ O(n²) xuống O(n)
3. **Backtracking mạnh mẽ**: Kỹ thuật quan trọng cho bài toán đường đi trong cây
4. **State pattern**: Theo dõi "độ khác biệt" thay vì "giá trị tuyệt đối" mở ra nhiều tối ưu hóa
5. **Không ngại brute force**: Đôi khi O(n²) với hằng số nhỏ chạy nhanh hơn O(n log n) với hằng số lớn
