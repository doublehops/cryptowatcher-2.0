package funcs

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func KeyExists(key string, m map[string]string) bool {

	_, ok := m[key]
	return ok
}

func TestServer(expectedResponse []byte) (string, []byte, error) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Test server received response")
	}))

	defer ts.Close()

	return "", expectedResponse, nil
}
