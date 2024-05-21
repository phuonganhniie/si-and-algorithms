package main

import (
	"fmt"
	"time"

	"github.com/phuonganhniie/educative/sorting"
)

const numCount = 100000
const fileName = "./data/not_sorted_data.txt"
const filePathSorted = "./data/sorted_data.txt"

func main() {
	// err := helper.GenerateRandomNumberToFile(fileName, numCount)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("Generated file with", numCount, "random numbers.")

	start := time.Now()

	// Step 1: Split and sort chunks
	chunkFiles := sorting.SplitAndSortChunks(fileName)

	// Step 2: Merge sorted chunks
	sorting.MergeSortedChunks(chunkFiles, filePathSorted)

	elapsed := time.Since(start)
	fmt.Println("External merge sort completed in:", elapsed)
}
