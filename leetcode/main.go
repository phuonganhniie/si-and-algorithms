package leetcode

func main() {
	// ================ LEETCODE ====================
	// #1 Valid Parentheses
	// rs := isValid("()")
	// fmt.Println(rs)

	// #2 Unable eat lunch students
	// students := []int{1, 1, 0, 0}
	// sandwiches := []int{0, 1, 0, 1}
	// rs2 := countStudents(students, sandwiches)
	// fmt.Println(rs2)

	// #3 Two sum
	// nums := []int{3, 2, 4}
	// target := 6
	// rs3 := twoSum(nums, target)
	// fmt.Println(rs3)

	// #4 Find the difference
	// s := "abc"
	// t := "bceax"
	// rs4 := leetcode.FindTheDifference(s, t)
	// fmt.Println(rs4)

	// #5 Product except self
	// nums := []int{1, 2, 3, 4}
	// rs5 := leetcode_238.ProductExceptSelf(nums)
	// fmt.Println(rs5)

	// #6 Minimum String Length After Removing Substrings
	// s := "ABFCACDB"
	// rs6 := minLength(s)
	// fmt.Println(rs6)

	// #7 Longest Substring Without Repeating Characters
	// s := "ppwwkew"
	// rs7 := lengthOfLongestSubstring(s)
	// fmt.Println(rs7)

	// #8 Middle of the Linked List
	// head := leetcode.ListNode{
	// 	Val: 1,
	// 	Next: &leetcode.ListNode{
	// 		Val: 2,
	// 		Next: &leetcode.ListNode{
	// 			Val: 3,
	// 			Next: &leetcode.ListNode{
	// 				Val: 4,
	// 				Next: &leetcode.ListNode{
	// 					Val: 5,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// leetcode.DeleteMiddle(&head)

	// #9 Longest Increasing Subsequence
	// nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	// nums := []int{0, 1, 0, 3, 2, 3}
	// rs9 := leetcode.LengthOfLISBS(nums)
	// fmt.Println(rs9)

	// #10 Longest Palindromic Substring

	// #11
	// word := "a123bc34d8ef34"
	// rs := leetcode.NumDifferentIntegers(word)
	// fmt.Println(rs)

	// #647
	// s := "babad"
	// result := leetcode_647.CountSubstrings(s)
	// fmt.Println("Number of palindromic substrings (DP):", result)

	// #2149
	// nums := []int{3, -1, -2, -5, -2, 4}
	// _ = leetcode_2149.RearrangeArray(nums)

	// #1481
	// arr := []int{4, 3, 1, 1, 3, 3, 2}
	// k := 3
	// leetcode_1481.FindLeastNumOfUniqueInts(arr, k)

	// #787
	// n := 4
	// flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	// src := 0
	// dst := 3
	// k := 1
	// fmt.Println("Starting BFS...")
	// cheapestPrice := leetcode_787.FindCheapestPrice(n, flights, src, dst, k)
	// fmt.Printf("Cheapest price: %d\n", cheapestPrice)

	// #2864
	// s := "010"
	// fmt.Println(leetcode_2864.MaximumOddBinaryNumber(s))

	// #977
	// nums := []int{-4, -1, 0, 3, 10}
	// leetcode_977.SortedSquares(nums)

	// #1750
	// s := "bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb"
	// s := "cabaabac"
	// s := "aaaaaaaaaaaa"
	// s := "aabccabba"
	// leetcode_1750.MinimumLength(s)

	// #349
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}
	// leetcode_349.Intersection(nums1, nums2)

	// #1669

	// Others
	// arr := []int{1, 2, 3, 4, 5}
	// leetcode.ShiftArrayRight(arr)

	// ================ CODILITY ====================
	// Lesson 3
	// arr := []int{3, 1, 2, 4, 3}
	// codility.SolutionLesson3_3a(arr)

	// arr1 := []int{2, 3, 1, 5}
	// codility.SolutionLesson3_2(arr1)

	// Lesson 7
	// s := "(()(())())"
	// codility.SolutionLesson7_3(s)

	// Lesson 10
	// codility.SolutionLesson10_1(24)

	// Lesson 13
	// A := []int{4, 4, 5, 5, 1}
	// B := []int{3, 2, 4, 3, 1}
	// codility.SolutionLesson13_2(A, B)

	// arr := []int{10, 1, 3, 1, 2, 2, 1, 0, 4}
	// codility.RewriteMaxNonIntersectingPairs(arr)

	// ================ CODING PATTERNS ====================
	/* Sliding Window */
	// fruits := []int{1, 2, 3, 2, 2}
	// sliding_window.TotalFruit(fruits)

	/* Two Pointers */
	// nums := []int{10, 5, 2, 6}
	// k := 100
	// two_pointers.NumSubarrayProductLessThanK(nums, k)

	// nums := []int{-4, -1, 0, 3, 10}
	// two_pointers.SortedSquares(nums)

	/* Breadth First Search */
	// graph := make(map[int][]int)
	// graph[1] = []int{2, 3, 5, 10}
	// graph[2] = []int{1, 4}
	// graph[3] = []int{1, 6, 7, 9}
	// graph[5] = []int{1, 8}
	// graph[10] = []int{1}
	// graph[4] = []int{2}
	// graph[6] = []int{3, 7}
	// graph[7] = []int{3, 6}
	// graph[9] = []int{3, 8}
	// graph[8] = []int{5, 9}
	// bfs.SimpleBFS(1, graph)

	// root := &bfs.TreeNode{Val: 3}
	// root.Left = &bfs.TreeNode{Val: 9}
	// root.Right = &bfs.TreeNode{Val: 20}
	// root.Right.Left = &bfs.TreeNode{Val: 15}
	// root.Right.Right = &bfs.TreeNode{Val: 7}
	// bfs.AverageOfLevels(root)

	// root := &bfs.TreeNode{Val: 3}
	// root.Left = &bfs.TreeNode{Val: 9}
	// root.Right = &bfs.TreeNode{Val: 20}
	// root.Right.Left = nil
	// root.Right.Right = nil
	// root.Right.Right = &bfs.TreeNode{Val: 15}
	// root.Right.Right = &bfs.TreeNode{Val: 7}
	// bfs.MinDepth(root)

	// time := bfs.OrangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
	// fmt.Println(time)

	// Depth First Search
	// root := &dfs.TreeNode{Val: 5}
	// root.Left = &dfs.TreeNode{Val: 4}
	// root.Right = &dfs.TreeNode{Val: 8}
	// root.Left.Left = &dfs.TreeNode{Val: 11}
	// root.Left.Left.Left = &dfs.TreeNode{Val: 7}
	// root.Left.Left.Right = &dfs.TreeNode{Val: 2}
	// root.Right.Left = &dfs.TreeNode{Val: 13}
	// root.Right.Right = &dfs.TreeNode{Val: 4}
	// root.Right.Right.Right = &dfs.TreeNode{Val: 1}
	// dfs.HasPathSum(root, 22)

	// s := "aabbbabaaa"
	// codility.SolutionTest1(s)

	// X := []int{2, 4, 2, 6, 7}
	// Y := []int{0, 5, 3, 2, 1}
	// W := 2
	// codility.SolutionTest4(X, Y, W)
}
