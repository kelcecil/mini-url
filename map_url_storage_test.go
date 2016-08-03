package main

import "testing"

// TestAddNewURLToMap ... Test to ensure that URLs are properly added and
// retrieved from the map storage driver.
func TestAddNewURLToMap(t *testing.T) {
	storage := InitMapStorage()
	shortURL1 := storage.AddNewURL("http://kelcecil.com")
	shortURL2 := storage.AddNewURL("http://www.google.com")

	if shortURL1 == shortURL2 {
		t.Errorf("Short URLs for two different URLs should not be the same.")
	}

	if shortURL1 != "a" {
		t.Errorf("Short URL for first insert wasn't what was expected; was: %v", shortURL1)
	}

	if shortURL2 != "b" {
		t.Errorf("Short URL for second insert wasn't what was expected; was: %v", shortURL2)
	}
}

// TestGetURLsFromMap ... Ensure that short URL parts retrieved when storing URLs
// are also able to retrieve the same URLs later.
func TestGetURLsFromMap(t *testing.T) {
	storage := InitMapStorage()
	testCases := map[string]string{
		"http://kelcecil.com": storage.AddNewURL("http://kelcecil.com"),
		"http://boxcast.com":  storage.AddNewURL("http://boxcast.com"),
	}

	for originalURL, shortURL := range testCases {
		retrievedURL := storage.GetURLByShortHash(shortURL)
		if originalURL != retrievedURL {
			t.Errorf("Look up by short hash failed for short URL; got: %v, wanted %v",
				retrievedURL,
				originalURL)
		}
	}
}
