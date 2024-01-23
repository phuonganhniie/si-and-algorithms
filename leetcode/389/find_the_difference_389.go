package leetcode

import "fmt"

func FindTheDifference(s string, t string) byte {
	diffS := byte(0)
	diffT := byte(0)

	for i := 0; i < len(s); i++ {
		diffS ^= s[i]
		fmt.Printf("Binary of s[i]: %b\n", s[i])
		fmt.Printf("Char of s[i]: %s\n", string(s[i]))
	}
	fmt.Printf("Diff S: %b\n", diffS)
	fmt.Printf("Char Diff S: %s\n", string(diffS))

	for j := 0; j < len(t); j++ {
		diffT ^= t[j]
		fmt.Printf("Binary of t[j]: %b\n", t[j])
		fmt.Printf("Char of t[j]: %s\n", string(t[j]))
	}
	fmt.Printf("Diff T: %b\n", diffT)
	fmt.Printf("Char Diff T: %s\n", string(diffT))

	result := diffS ^ diffT
	fmt.Printf("DiffS XOR DiffT: %b\n", result)

	return result

	// =================================
	// compareMap := make(map[rune]int)
	// result := []byte{}

	// for _, char := range s {
	// 	compareMap[char]--
	// }

	// for _, char := range t {
	// 	compareMap[char]++
	// 	if compareMap[char] == 1 {
	// 		result = append(result, byte(char))
	// 		// return byte(char)
	// 	}
	// }

	// return result
}
