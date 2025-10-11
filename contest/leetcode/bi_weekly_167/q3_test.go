package biweekly167

import "testing"

func TestExamTracker(t *testing.T) {
	examTracker := Constructor()

	// Test case from example
	examTracker.Record(1, 98)
	result := examTracker.TotalScore(1, 1)
	expected := int64(98)
	if result != expected {
		t.Errorf("TotalScore(1, 1) = %d, want %d", result, expected)
	}

	examTracker.Record(5, 99)
	result = examTracker.TotalScore(1, 3)
	expected = int64(98)
	if result != expected {
		t.Errorf("TotalScore(1, 3) = %d, want %d", result, expected)
	}

	result = examTracker.TotalScore(1, 5)
	expected = int64(197)
	if result != expected {
		t.Errorf("TotalScore(1, 5) = %d, want %d", result, expected)
	}

	result = examTracker.TotalScore(3, 4)
	expected = int64(0)
	if result != expected {
		t.Errorf("TotalScore(3, 4) = %d, want %d", result, expected)
	}

	result = examTracker.TotalScore(2, 5)
	expected = int64(99)
	if result != expected {
		t.Errorf("TotalScore(2, 5) = %d, want %d", result, expected)
	}
}

func TestExamTrackerEdgeCases(t *testing.T) {
	// Test empty tracker
	examTracker := Constructor()
	result := examTracker.TotalScore(1, 10)
	expected := int64(0)
	if result != expected {
		t.Errorf("TotalScore on empty tracker = %d, want %d", result, expected)
	}

	// Test single record
	examTracker.Record(5, 100)
	result = examTracker.TotalScore(5, 5)
	expected = int64(100)
	if result != expected {
		t.Errorf("TotalScore(5, 5) = %d, want %d", result, expected)
	}

	// Test range that doesn't include any records
	result = examTracker.TotalScore(1, 4)
	expected = int64(0)
	if result != expected {
		t.Errorf("TotalScore(1, 4) = %d, want %d", result, expected)
	}

	// Test range after all records
	result = examTracker.TotalScore(6, 10)
	expected = int64(0)
	if result != expected {
		t.Errorf("TotalScore(6, 10) = %d, want %d", result, expected)
	}
}
