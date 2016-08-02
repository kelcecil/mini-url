package main

var BASE62_ALPHABET string = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789"

// IdToHash ... Take a numeric id and convert to a user friendly string
// for use.
func IdToHash(id int) string {
	digits := FindDigitsForInt(id)
	shortenedUrlHash := ""

	for i := range digits {
		// Get the alphabet indice from our converted digits
		indice := digits[i]

		// Get a one letter range to easily get a string and add to the hash
		shortenedUrlHash = shortenedUrlHash + BASE62_ALPHABET[indice:indice+1]
	}
	return shortenedUrlHash
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
