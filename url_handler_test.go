package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

func checkForError(t *testing.T, message string, err error) {
	if err != nil {
		t.Errorf("%v. Reason: %v", message, err.Error())
	}
}
