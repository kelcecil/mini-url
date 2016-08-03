package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ShortUrlForwardingHandler .. Struct representing the handler that accepts
// new URLs and redirects short urls. The struct accepts a struct conformed to
// the ShortUrlStorage interface to allow interchangable storage based on cloud
// provider.
type ShortUrlForwardingHandler struct {
	Storage ShortUrlStorage
}

// NewUrlSubmission .. This struct contains attributes from a PUT to add a new
// URL.
type NewUrlSubmission struct {
	URL string `json:"url"`
}

// ShortenedUrl .. This struct contains attributes from a response following
// a PUT to add a new URL.
type ShortenedUrl struct {
	Key string `json:"key"`
}

// ServeHTTP .. This handler either adds a new url if the request is a PUT
// action or sends a redirect if a URL exists for a given key.
// This method satisfies the Handler interface and is intended to be
// passed to http.handle
func (handler ShortUrlForwardingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method == http.MethodPut {
		handler.handleAddingNewShortUrl(w, r)
	} else {
		handler.handleGettingShortUrl(w, r)
	}
}

// handleAddingNewShortUrl .. This method adds a new URL to our storage and
// returns a key for use in a short-url. This method is intended to be called
// as the result of a PUT operation.
func (handler ShortUrlForwardingHandler) handleAddingNewShortUrl(w http.ResponseWriter, r *http.Request) {
	// Read the body of the request and parse the new URL to be added.
	data, err := ioutil.ReadAll(r.Body)
	if checkAndHandleError(err, w, r) != nil {
		return
	}
	var newUrl NewUrlSubmission
	if err := json.Unmarshal(data, &newUrl); err != nil {
		log.Printf("Failed to unmarshal json: %v", err.Error())
	}

	key := handler.Storage.AddNewUrl(newUrl.URL)

	// Prepare a JSON response to let the other end know the new key for the URL.
	response := ShortenedUrl{Key: key}
	jsonResponse, err := json.Marshal(response)
	if checkAndHandleError(err, w, r) != nil {
		return
	}

	// Write response.
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	log.Printf("Added new short URL: %v as key: %v", newUrl.URL, key)
}

// handleGettingShortUrl .. Thos method parses the key from the URL, finds the URL
// for the redirect, and sends a redirect to the URL back to the user.
func (handler ShortUrlForwardingHandler) handleGettingShortUrl(w http.ResponseWriter, r *http.Request) {
	// Fetch the hash from the URL
	urlPath := strings.Split(r.URL.Path[1:len(r.URL.Path)], "/")
	key := urlPath[0]
	log.Printf("Fetching URL for key: %v", key)

	// Look up the URL and redirect to the stored URL
	newUrl := handler.Storage.GetUrlByShortHash(key)
	http.Redirect(w, r, newUrl, http.StatusTemporaryRedirect)
}

// checkAndHandleError ... This function serves to DRY up a common sequence
// of error checking and writing a response in the event of a failure.
func checkAndHandleError(err error, w http.ResponseWriter, r *http.Request) error {
	if err != nil {
		log.Printf("Issue processing HTTP request: %v", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	return err
}
