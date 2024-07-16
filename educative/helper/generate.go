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

type Set struct {
	hashMap map[interface{}]bool
}

func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[interface{}]bool)
	return s
}

func (s *Set) Add(value interface{}) {
	s.hashMap[value] = true
}

func (s *Set) Delete(value interface{}) {
	delete(s.hashMap, value)
}

func (s *Set) Exists(value interface{}) bool {
	_, exist := s.hashMap[value]
	return exist
}
