package main

import "testing"

// TestDigitsForInt ... Ensure that identifiers are broken down into
// base62 parts correctly.
func TestDigitsForInt(t *testing.T) {
	testCases := map[int][]int{
		100: []int{38, 1},
		62:  []int{0, 1},
	}

	for testId, desiredTestResult := range testCases {
		computedResult := FindDigitsForInt(testId)
		if !slicesAreEqual(computedResult, desiredTestResult) {
			t.Errorf("Converting id to base 62 failed for value: %v; got: %v", testId, computedResult)
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
