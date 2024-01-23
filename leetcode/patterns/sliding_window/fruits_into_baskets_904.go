package sliding_window

func TotalFruit(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	}

	maxFruits, left := 0, 0
	baskets := make(map[int]int)

	for right, fruit := range fruits {
		// Add current fruit to basket
		baskets[fruit]++

		// If we have more than two types of fruits, move the left pointer
		for len(baskets) > 2 {
			leftFruit := fruits[left]
			baskets[leftFruit]-- // decrease the number of fruits of previous type in basket

			if baskets[leftFruit] == 0 {
				delete(baskets, leftFruit) // remove previous type fruit in basket
			}
			left++
		}

		// Update the maximum number of fruits in the baskets
		currentWindowFruits := right - left + 1
		if currentWindowFruits > maxFruits {
			maxFruits = currentWindowFruits
		}
	}
	return maxFruits
}

/**
Array: [ 1, 2, 3, 4, 5 ]
Indexes: 0  1  2  3  4
		 |-----------|
		left       right

		The number of elements between left and right, inclusive:
		right - left + 1 = 4 - 0 + 1 = 5
*/
