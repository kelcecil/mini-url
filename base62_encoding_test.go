package main

import "testing"

// TestIdToHash ... Test that identifier numbers are properly converted into
// short ids properly.
func TestIDToHash(t *testing.T) {
	testCases := map[int]string{
		100:   "Mb",
		62:    "ab",
		63:    "bb",
		10000: "sLc",
	}

	for testID, desiredTestResult := range testCases {
		result := IDToHash(testID)
		if result != desiredTestResult {
			t.Errorf("id to shortened string failed for value: %v; got: %v", testID, result)
		}
	}
}

// TestDigitsForInt ... Ensure that identifiers are broken down into
// base62 parts correctly.
func TestDigitsForInt(t *testing.T) {
	testCases := map[int][]int{
		100: []int{38, 1},
		62:  []int{0, 1},
	}

	for testID, desiredTestResult := range testCases {
		computedResult := FindDigitsForInt(testID)
		if !slicesAreEqual(computedResult, desiredTestResult) {
			t.Errorf("Converting id to base 62 failed for value: %v; got: %v", testID, computedResult)
		}
	}
}

// slicesAreEqual ... Helper for comparing two slices of integers for equality.
func slicesAreEqual(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
