package main

import "testing"

func TestAddNewUrlToMap(t *testing.T) {
	storage := InitMapStorage()
	shortUrl1 := storage.AddNewUrl("http://kelcecil.com")
	shortUrl2 := storage.AddNewUrl("http://www.google.com")

	if shortUrl1 == shortUrl2 {
		t.Errorf("Short Urls for two different URLs should not be the same.")
	}

	if shortUrl1 != "a" {
		t.Errorf("Short URL for first insert wasn't what was expected; was: %v", shortUrl1)
	}

	if shortUrl2 != "b" {
		t.Errorf("Short URL for second insert wasn't what was expected; was: %v", shortUrl2)
	}
}

func TestGetUrlsFromMap(t *testing.T) {
	storage := InitMapStorage()
	testCases := map[string]string{
		"http://kelcecil.com": storage.AddNewUrl("http://kelcecil.com"),
		"http://boxcast.com":  storage.AddNewUrl("http://boxcast.com"),
	}

	for originalUrl, shortUrl := range testCases {
		retrievedUrl := storage.GetUrlByShortHash(shortUrl)
		if originalUrl != retrievedUrl {
			t.Errorf("Look up by short hash failed for short url; got: %v, wanted %v",
				retrievedUrl,
				originalUrl)
		}
	}
}
