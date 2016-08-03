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

func TestAddNewUrlViaHandler(t *testing.T) {
	handler := ShortUrlForwardingHandler{
		Storage: InitMapStorage(),
	}

	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	newUrl := NewUrlSubmission{
		URL: "http://google.com",
	}

	jsonData, err := json.Marshal(newUrl)
	checkForError(t, "Failed when marshalling json", err)

	req, err := http.NewRequest(http.MethodPut, testServer.URL, bytes.NewReader(jsonData))
	checkForError(t, "Failed to create new put request", err)

	client := http.Client{}
	response, err := client.Do(req)
	checkForError(t, "Failed to execute PUT action on server", err)

	buf, err := ioutil.ReadAll(response.Body)
	checkForError(t, "Failed to read response body from PUT", err)

	var key ShortenedUrl
	err = json.Unmarshal(buf, &key)
	checkForError(t, "Failed to unmarshal JSON", err)

	if handler.Storage.GetUrlByShortHash(key.Key) != "http://google.com" {
		t.Errorf("Hash provided in response to URL did not match URL in store. Key: %v, URL: %v", key.Key, newUrl.URL)
	}
}

func TestRedirectToStoredUrl(t *testing.T) {
	redirectError := errors.New("Redirect occurred")
	storage := MapUrlStorage{
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

func tryGetURL(t *testing.T, storage MapUrlStorage, client *http.Client, key, redirectUrl, desiredUrl string) {
	targetUrl := fmt.Sprintf("%v/%v", redirectUrl, key)
	req, err := http.NewRequest(http.MethodGet, targetUrl, nil)
	checkForError(t, "Request for get failed", err)

	t.Logf("Fetching target url: %v", targetUrl)
	response, _ := client.Do(req)

	redirectedUrl, err := response.Location()
	checkForError(t, "Unable to retrieve response location", err)

	if !strings.Contains(desiredUrl, redirectedUrl.Host) {
		t.Errorf("The url was expected to be %v but was %v instead.", desiredUrl, redirectedUrl.Host)
	}
}

func initializeTestServer(storage ShortUrlStorage) *httptest.Server {
	handler := ShortUrlForwardingHandler{
		Storage: storage,
	}
	return httptest.NewServer(handler)
}

func checkForError(t *testing.T, message string, err error) {
	if err != nil {
		t.Errorf("%v. Reason: %v", message, err.Error())
	}
}
