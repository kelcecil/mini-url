package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestAddNewURLViaHandler .. This test is intended to verify that we can make
// a PUT call to the root of the server and add a new short URL for use.
func TestAddNewURLViaHandler(t *testing.T) {
	handler := ShortURLForwardingHandler{
		Storage: InitMapStorage(),
	}

	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	newURL := NewURLSubmission{
		URL: "http://google.com",
	}

	jsonData, err := json.Marshal(newURL)
	checkForError(t, "Failed when marshalling json", err)

	req, err := http.NewRequest(http.MethodPut, testServer.URL, bytes.NewReader(jsonData))
	checkForError(t, "Failed to create new put request", err)

	client := http.Client{}
	response, err := client.Do(req)
	checkForError(t, "Failed to execute PUT action on server", err)

	buf, err := ioutil.ReadAll(response.Body)
	checkForError(t, "Failed to read response body from PUT", err)

	var key ShortenedURL
	err = json.Unmarshal(buf, &key)
	checkForError(t, "Failed to unmarshal JSON", err)

	if handler.Storage.GetURLByShortHash(key.Key) != "http://google.com" {
		t.Errorf("Hash provided in response to URL did not match URL in store. Key: %v, URL: %v", key.Key, newURL.URL)
	}
}

// TestRedirectToStoredURL .. These tests are intended to ensure that we receive
// redirects for URLs that have been added to storage previously.
func TestRedirectToStoredURL(t *testing.T) {
	redirectError := errors.New("Redirect occurred")
	storage := MapURLStorage{
		Storage: map[string]string{
			"a": "http://google.com",
			"b": "http://kelcecil.com",
		},
	}
	testServer := initializeTestServer(&storage)
	defer testServer.Close()

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return redirectError
		},
	}
	for key, value := range storage.Storage {
		tryGetURL(t, storage, client, key, testServer.URL, value)
	}
}

// tryGetURL .. Convenience function to test if a HTTP redirect goes where we
// expect it to go. If it does not, then we fail the test.
func tryGetURL(t *testing.T, storage MapURLStorage, client *http.Client, key, redirectURL, desiredURL string) {
	targetURL := fmt.Sprintf("%v/%v", redirectURL, key)
	req, err := http.NewRequest(http.MethodGet, targetURL, nil)
	checkForError(t, "Request for get failed", err)

	t.Logf("Fetching target URL: %v", targetURL)
	response, _ := client.Do(req)

	redirectedURL, err := response.Location()
	checkForError(t, "Unable to retrieve response location", err)

	if !strings.Contains(desiredURL, redirectedURL.Host) {
		t.Errorf("The URL was expected to be %v but was %v instead.", desiredURL, redirectedURL.Host)
	}
}

// initializeTestServer .. Convenience function to create the test server for
// the event in which an overriden redirect policy is not required.
func initializeTestServer(storage ShortURLStorage) *httptest.Server {
	handler := ShortURLForwardingHandler{
		Storage: storage,
	}
	return httptest.NewServer(handler)
}

// checkForError .. Helpful function to reduce the error checking code in the
// handler check code. Pass the testing pointer, a message, and an error to
// check for failure.
func checkForError(t *testing.T, message string, err error) {
	if err != nil {
		t.Errorf("%v. Reason: %v", message, err.Error())
	}
}
