package codility

import "sort"

func SolutionLesson2_2(A []int) int {
	countMap := map[int]int{}

	for _, num := range A {
		countMap[num]++
	}

	for num, count := range countMap {
		if count%2 != 0 {
			return num
		}
	}

	return -1
}

// func Solution(A []int) int:
//     // Create a map to track the presence of elements
//     presenceMap := make(map[int]bool)

//     // Fill the map with elements of A
//     for _, value := range A:
//         if value > 0:
//             presenceMap[value] = true

//     // Find the smallest positive integer not in the map
//     for i := 1; i <= len(A) + 1; i++:
//         if !presenceMap[i]:
//             return i

//     // This line should never be reached with correct logic
//     return -1 // Indicates an error

// function Solution(S):
//     hashmap = new Hashmap()
//     for length from 1 to N: # N is the length of string S
//         for i from 0 to N-length:
//             substring = S[i:i+length]
//             if substring in hashmap:
//                 hashmap[substring] = hashmap[substring] + 1
//             else:
//                 hashmap[substring] = 1

//     shortest_unique_length = N + 1 # Initialize to a length larger than the maximum possible length of S
//     for substring, count in hashmap:
//         if count == 1 and length(substring) < shortest_unique_length:
//             shortest_unique_length = length(substring)

//     if shortest_unique_length == N + 1:
//         return -1 # This means no unique substring was found
//     else:
//         return shortest_unique_length

// function Solution(S):
// hashmap = new Hashmap()
// for length from 1 to N: # N is the length of string S
// 	for i from 0 to N-length:
// 		substring = S[i:i+length]
// 		if substring in hashmap:
// 			hashmap[substring] = hashmap[substring] + 1
// 		else:
// 			hashmap[substring] = 1

// shortest_unique_length = N + 1 # Initialize to a length larger than the maximum possible length of S
// for substring, count in hashmap:
// 	if count == 1 and length(substring) < shortest_unique_length:
// 		shortest_unique_length = length(substring)

// if shortest_unique_length == N + 1:
// 	return -1 # This means no unique substring was found
// else:
// 	return shortest_unique_length

// function Solution(S):
// N = length of S
// for l = 1 to N:
// 	hashMap = new Hash Map
// 	for i = 0 to N-l:
// 		substring = S[i:i+l]
// 		if substring in hashMap:
// 			hashMap[substring] += 1
// 		else:
// 			hashMap[substring] = 1

// 	for key in hashMap:
// 		if hashMap[key] == 1:
// 			return l

// return -1

func SolutionTest1(S string) int {
	n := len(S)
	uniqueMap := make(map[string]int)

	for l := 1; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			uniqueStr := S[i : i+l]

			if _, ok := uniqueMap[uniqueStr]; ok {
				uniqueMap[uniqueStr] += 1
			} else {
				uniqueMap[uniqueStr] = 1
			}
		}

		for k := range uniqueMap {
			if uniqueMap[k] == 1 {
				return l
			}
		}
	}

	return -1
}

/*
In order to finish a game, a player has to complete N missions. The missions are numbered from 0 to N-1. The K-th mission has an integer D[K] assigned, representing its difficulty level.

During a day, you can perform any number of missions given the two following rules:

- missions should be performed in the specified order, in other words, a mission can be undertaken only if all of the missions preceding it have already been completed;
- the difference between the difficulty levels of any two missions performed on the same day should not be greater than an integer X.

Write a function:
```
func Solution(D []int, X int) int
```
that, given an array D of N integers and an integer X, returns the minimum number of days required to complete all of the missions in the game.

Examples:

1. Given D = [5, 8, 2, 7] and X = 3, your function should return 3. The first two missions can be performed on the first day, the third mission on the second day, and the fourth mission on the third day.
The text in the image is as follows:

2. Given D = [2, 5, 9, 2, 1, 4] and X = 4, your function should return 3. The first two missions can be performed on the first day, the third mission on the second day and all of the remaining missions on the third day. Note that it is possible to perform the first mission on the first day and the next two missions on the second day. In both of these cases, the minimum number of days required to complete all of the missions is 3.

3. Given D = [1, 12, 10, 4, 5, 2] and X = 2, your function should return 4. The first mission can be performed on the first day, the next two missions on the second day, the fourth and fifth missions on the third day, and the last remaining mission on the fourth day. It is not possible to complete all of the missions in fewer days.

Write an efficient algorithm for the following assumptions:

- N is an integer within the range [1..200,000];
- X is an integer within the range [0..1,000,000,000];
- each element of array D is an integer within the range [1..1,000,000,000].
*/

func SolutionTest2(D []int, X int) int {
	n := len(D)
	finalDay := 1
	minLevel := D[0]
	maxLevel := D[0]

	for i := 1; i <= n-1; i++ {
		currentLevel := D[i]

		// Check if the current level fits in the current day
		if currentLevel-minLevel <= X && maxLevel-currentLevel <= X {
			minLevel = min(minLevel, currentLevel)
			maxLevel = max(maxLevel, currentLevel)
		} else {
			// increase the final day
			finalDay++
			minLevel = currentLevel
			maxLevel = currentLevel
		}
	}

	return finalDay
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
There are N potholes on a straight road parallel to the y-axis. The positions of the potholes are described by two arrays of integers, X and Y. The K-th pothole (for K within the range [0..N-1]) is located at coordinates (X[K], Y[K]).

In order to fix the potholes, a road roller of width W will be used. The road roller can only drive along the street (parallel to the y-axis). It can start driving from any x coordinate at the beginning of the road (a point whose y coordinate is equal to 0). During one drive, for a chosen starting point (x, 0) of the road roller's left end, the road roller drives upwards and patches all potholes on its way that fall within the width of the road roller, W, and are to the right of its starting position, x. In other words, it patches all the potholes whose first coordinate is between x and x+W inclusive.

What is the minimum number of road roller drives required to patch all the potholes?

Write a function:

```
func Solution(X [int], Y [int], W int) int
```

that, given two arrays of integers X and Y, describing the positions of the N potholes, and an integer W, specifying the width of the road roller, returns the minimum number of drives needed to patch all the potholes.

Examples:
1. Given X = [2, 4, 2, 6, 7], Y = [0, 5, 3, 2, 1], and W = 2, the answer is 3. At least three drives will be needed to patch all the potholes. For example, the drives could start respectively at points (1, 0), (4, 0) and (6, 0). During the first drive, the road roller would patch the

2. Given X = [4, 8, 2, 2, 1, 4], Y = [1, 2, 3, 1, 2, 3] and W = 3, the answer is 2. If the first drive of the road roller started at point (1, 0) and the second drive at the point (6, 0), all the potholes would be patched.

3. Given X = [0, 3, 6, 5], Y = [0, 3, 2, 4] and W = 1, the answer is 3. The first drive of the road roller could start at point (0, 0), the second drive at point (3, 0) and the third at point (5, 0).

Write an efficient algorithm for the following assumptions:
N is an integer within the range [1..100,000];
each element of arrays X and Y is an integer within the range [0..1,000,000,000];
W is an integer within the range [1..1,000,000,000].
*/
type Potholes struct {
	x, y int
}

func SolutionTest4(X, Y []int, W int) int {
	n := len(X)
	potholes := make([]Potholes, n)

	// Tạo cặp toạ độ ổ gà X, Y từ mảng X, Y được cho
	for i := 0; i < n; i++ {
		potholes[i] = Potholes{x: X[i], y: Y[i]}
	}

	// Sort toạ độ của ổ gà theo toạ độ X
	sort.Slice(potholes, func(i, j int) bool {
		return potholes[i].x < potholes[j].x
	})

	drives := 0
	maxRepairHoles := -1
	for _, pothole := range potholes {
		// Nếu có ổ gà tại toạ độ đang đứng thì sửa nó và tăng số lần lái thêm 1
		if pothole.x > maxRepairHoles {
			drives++
			maxRepairHoles = pothole.x + W - 1
		}
	}

	return drives
}

func SolutionTest6(numbers []int) int {
	startDigit := make([]int, 10)
	endDigit := make([]int, 10)

	for _, num := range numbers {
		start := getFirstDigit(num)
		end := getEndDigit(num)
		startDigit[start] += 1
		endDigit[end] += 1
	}

	pairs := 0
	for i := 0; i <= 9; i++ {
		pairs = pairs + startDigit[i]*endDigit[i]
	}

	return pairs
}

func getFirstDigit(number int) int {
	for number >= 10 {
		number = number / 10
	}
	return number
}

func getEndDigit(number int) int {
	return number % 10
}
