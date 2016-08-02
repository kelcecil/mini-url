package main

// ShortUrlStorage ... Interface to allow storage of small URLs and
// facilitate mocks for testing.
type ShortUrlStorage interface {
	AddNewUrl(string) string
	GetUrlByShortHash(string) string
}
