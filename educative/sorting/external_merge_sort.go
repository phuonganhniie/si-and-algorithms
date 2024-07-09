package sorting

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const chunkSize = 10000

func SplitAndSortChunks(inputFile string) []string {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	chunkFiles := []string{}
	chunk := []int{}

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		chunk = append(chunk, num)

		if len(chunk) >= chunkSize {
			sort.Ints(chunk)
			chunkFiles = append(chunkFiles, writeChunk(chunk))
			chunk = []int{}
		}
	}

	if len(chunk) > 0 {
		sort.Ints(chunk)
		chunkFiles = append(chunkFiles, writeChunk(chunk))
	}

	return chunkFiles
}

func writeChunk(chunk []int) string {
	file, err := os.CreateTemp("", "chunk_")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return ""
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, num := range chunk {
		fmt.Fprintln(writer, num)
	}
	writer.Flush()

	return file.Name()
}

func MergeSortedChunks(chunkFiles []string, outputFile string) {
	chunkReaders := make([]*bufio.Scanner, len(chunkFiles))
	for i, file := range chunkFiles {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening chunk file:", err)
			return
		}
		defer f.Close()
		chunkReaders[i] = bufio.NewScanner(f)
		chunkReaders[i].Scan()
	}

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()
	writer := bufio.NewWriter(output)

	for {
		minIndex := -1
		for i, reader := range chunkReaders {
			if reader == nil {
				continue
			}
			if minIndex == -1 || less(reader.Text(), chunkReaders[minIndex].Text()) {
				minIndex = i
			}
		}

		if minIndex == -1 {
			break
		}

		writer.WriteString(chunkReaders[minIndex].Text() + "\n")
		if !chunkReaders[minIndex].Scan() {
			chunkReaders[minIndex] = nil
		}
	}

	writer.Flush()
}

func less(a, b string) bool {
	numA, _ := strconv.Atoi(a)
	numB, _ := strconv.Atoi(b)
	return numA < numB
}
