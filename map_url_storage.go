package main

// MapUrlStorage ... This is just a simple structure for URL storage backed by Go's
// always-dependable map.
type MapUrlStorage struct {
	Storage map[string]string
}

func InitMapStorage() *MapUrlStorage {
	return &MapUrlStorage{
		Storage: make(map[string]string),
	}
}

// AddNewUrl ... Add a new url into the map storage using the hash identifier
// as the key and the url to be shortened as the value.
func (s *MapUrlStorage) AddNewUrl(url string) string {
	hashIdentifier := IdToHash(len(s.Storage))

	// TODO - Valudate URL and return if URL is invalid.
	s.Storage[hashIdentifier] = url
	return hashIdentifier
}

func (s *MapUrlStorage) GetUrlByShortHash(shortHash string) string {
	return s.Storage[shortHash]
}
