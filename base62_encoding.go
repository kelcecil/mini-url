package main

// IdToHash ... Take a numeric id and convert to a user friendly string
// for use.
func IdToHash(id int64) string {
	return ""
}

// findDigitsForInt ... Obtain the individual digits that will be used
// to find the replacement letters in our base62 alphabet.
func FindDigitsForInt(dividend int) []int {
	digits := make([]int, 0)
	var remainder int

	for dividend > 0 {
		remainder = dividend % 62
		dividend = dividend / 62
		digits = append(digits, remainder)
	}

	return digits
}
