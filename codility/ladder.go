package codility

/**
### Step 1: Find the Maximum Value in `A`
	1. Objective: Determine the largest number of rungs on any ladder.
	2. Why: To know up to which Fibonacci number you need to calculate, ensuring efficiency by avoiding redundant calculations.

### Step 2: Precompute Fibonacci Sequence
	1. Objective: Calculate and store the Fibonacci sequence up to `Fib(max(A))`.
	2. Why: The Fibonacci sequence directly gives us the number of ways to climb a ladder with `n` rungs. Storing these values prevents recalculating them for each ladder.

### Step 3: Initialize Result Array
	1. Objective: Create an array to store the result for each ladder.
	2. Why: You need a place to store your answers for the problem's output.

### Step 4: Calculate Results for Each Ladder
	1. Objective: For each ladder, calculate the number of ways to climb it using the precomputed Fibonacci sequence, and then apply the modulo operation as specified by the corresponding value in `B`.
	2. Why: This is the core calculation required by the problem, adjusted to the specific modulo requirement.
		Formula: `Result[i] = Fib(A[i]) % 2^B[i]`
	3. How:
		- Retrieve the Fibonacci number corresponding to the number of rungs on the ladder: `Fib(A[i])`.
		- Calculate `2^B[i]` (this can be done efficiently using bit shifting: `1 << B[i]`).
		- Apply the modulo operation: `Result[i] = Fib(A[i]) % (1 << B[i])`.

### Step 5: Return the Result Array
	1. Objective: Provide the final output of the program.
	2. Why: To adhere to the function’s specified return type and complete the task.
	3. How: Simply return the array that you’ve been populating with results in step 4.

By following these steps, you ensure that you're solving the problem efficiently, preventing redundant calculations, and handling potentially large numbers in a safe manner through the modulo operation. This approach also ensures that you adhere to the specific requirements of the problem, providing accurate and optimized results.
*/

// func SolutionLesson13_2(A []int, B []int) []int {
// 	maxVal := max(A)
// 	l := len(A)

// 	result := make([]int, l)
// 	fib := make([]int, maxVal+2)
// 	fib[1] = 1
// 	fib[2] = 2

// 	for i := 3; i <= maxVal; i++ {
// 		fib[i] = (fib[i-1] + fib[i-2]) % (1 << 30) // to avoid overflow
// 	}

// 	for i := 0; i < l; i++ {
// 		modulo := 1 << B[i]
// 		result[i] = fib[A[i]] % modulo
// 	}
// 	return result
// }

// func max(arr []int) int {
// 	maxVal := arr[0]
// 	for _, num := range arr {
// 		if num > maxVal {
// 			maxVal = num
// 		}
// 	}
// 	return maxVal
// }
