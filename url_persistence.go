package main

// ShortURLStorage ... Interface to allow storage of small URLs and
// facilitate mocks for testing.
type ShortURLStorage interface {
	AddNewURL(string) string
	GetURLByShortHash(string) string
}
