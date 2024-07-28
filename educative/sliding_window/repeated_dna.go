package slidingwindow

import "math"

/* --------------------------------- */
/* ----- SOLUTION BY EDUCATIVE ----- */
/* --------------------------------- */

type Set struct {
	hashMap map[interface{}]bool
	order   []interface{}
}

// NewSet will initialize and return a new object of Set.
func NewSet() *Set {
	s := new(Set)
	s.hashMap = make(map[interface{}]bool)
	return s
}

// Add will add the value in the Set.
func (s *Set) Add(value interface{}) {
	if _, exists := s.hashMap[value]; !exists {
		s.hashMap[value] = true
		s.order = append(s.order, value)
	}
}

// Delete will delete the value from the set.
func (s *Set) Delete(value interface{}) {
	if _, exists := s.hashMap[value]; exists {
		delete(s.hashMap, value)
		for i, v := range s.order {
			if v == value {
				s.order = append(s.order[:i], s.order[i+1:]...)
				break
			}
		}
	}
}

// Exists will check if the value exists in the set or not.
func (s *Set) Exists(value interface{}) bool {
	_, ok := s.hashMap[value]
	return ok
}

func (s *Set) SetToArrayStr() []string {
	var rs []string
	for _, v := range s.order {
		rs = append(rs, v.(string))
	}
	return rs
}

func findRepeatedSequences(dna string, k int) Set {
	stringLength := len(dna)

	if stringLength < k {
		return *NewSet()
	}

	mapping := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}

	baseValue := 4

	numbers := make([]int, stringLength)
	for i, ch := range dna {
		numbers[i] = mapping[ch]
	}

	hashValue := 0

	hashSet := *NewSet()
	output := *NewSet()

	for i := 0; i < stringLength-k+1; i++ {

		if i == 0 {
			for j := 0; j < k; j++ {
				hashValue += numbers[j] * int(math.Pow(float64(baseValue), float64(k-j-1)))
			}
		} else {
			previousHashValue := hashValue
			hashValue = ((previousHashValue - numbers[i-1]*int(math.Pow(float64(baseValue), float64(k-1)))) * baseValue) + numbers[i+k-1]
		}

		if hashSet.Exists(hashValue) {
			output.Add(dna[i : i+k])
		}

		hashSet.Add(hashValue)
	}

	return output
}

/* --------------------------------- */
/* ----- SOLUTION BY EDUCATIVE ----- */
/* --------------------------------- */
