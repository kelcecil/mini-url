package main

// MapURLStorage ... This is just a simple structure for URL storage backed by Go's
// always-dependable map.
type MapURLStorage struct {
	Storage map[string]string
}

// InitMapStorage ... Convience method for creatking a new map-based URL storage.
func InitMapStorage() *MapURLStorage {
	return &MapURLStorage{
		Storage: make(map[string]string),
	}
}

// AddNewURL ... Add a new url into the map storage using the hash identifier
// as the key and the url to be shortened as the value.
func (s *MapURLStorage) AddNewURL(url string) string {
	hashIdentifier := IDToHash(len(s.Storage))

	// TODO - Valudate URL and return if URL is invalid.
	s.Storage[hashIdentifier] = url
	return hashIdentifier
}

// GetURLByShortHash ... Retrieve a stored URL by providing a short hash.
func (s *MapURLStorage) GetURLByShortHash(shortHash string) string {
	return s.Storage[shortHash]
}
