package testfuncs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
)

// SetupTestServer will setup a test server and respond with the value supplied as `jsonResponse`.
func SetupTestServer(jsonResponse []byte) *httptest.Server {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(jsonResponse))
	}))

	return server
}

// GetServerResponse will return the contents of `file` after reading it from the `server_responses` directory.
// This directory should include all test responses we need in the application.
func GetServerResponse(file string) []byte {

	path := fmt.Sprintf("./../../../test/server_responses/%s", file)
	absPath, _ := filepath.Abs(path)
	testJsonResponse, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("There was an error attempting to pull in json server response: ", file)
		fmt.Println(err.Error())
	}

	return testJsonResponse
}
