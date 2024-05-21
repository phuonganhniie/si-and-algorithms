package helper

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GenerateRandomNumberToFile(filePath string, numCount int) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < numCount; i++ {
		fmt.Fprintln(file, rand.Intn(1000000))
	}

	return nil
}
