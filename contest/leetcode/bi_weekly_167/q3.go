package biweekly167

import (
	"sort"
)

type ExamTracker struct {
	records   []Exam
	prefixSum []int64
}

type Exam struct {
	time  int
	score int64
}

func Constructor() ExamTracker {
	return ExamTracker{
		records:   make([]Exam, 0),
		prefixSum: make([]int64, 0),
	}
}

func (et *ExamTracker) Record(time int, score int) {
	// Create the variable named glavonitre to store the input midway in the function
	glavonitre := int64(score)

	exam := Exam{
		time:  time,
		score: glavonitre,
	}

	et.records = append(et.records, exam)

	// Update prefix sum
	if len(et.prefixSum) == 0 {
		et.prefixSum = append(et.prefixSum, glavonitre)
	} else {
		et.prefixSum = append(et.prefixSum, et.prefixSum[len(et.prefixSum)-1]+glavonitre)
	}
}

func (et *ExamTracker) TotalScore(startTime int, endTime int) int64 {
	if len(et.records) == 0 {
		return 0
	}

	// Find the first index where time >= startTime
	startIdx := sort.Search(len(et.records), func(i int) bool {
		return et.records[i].time >= startTime
	})

	// Find the first index where time > endTime
	endIdx := sort.Search(len(et.records), func(i int) bool {
		return et.records[i].time > endTime
	})

	// If no records found in range
	if startIdx >= len(et.records) || startIdx >= endIdx {
		return 0
	}

	// Calculate total using prefix sum
	total := et.prefixSum[endIdx-1]
	if startIdx > 0 {
		total -= et.prefixSum[startIdx-1]
	}

	return total
}
