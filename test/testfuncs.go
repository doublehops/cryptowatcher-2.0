package testfuncs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
)

func SetupTestServer(jsonResponse []byte) *httptest.Server {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(jsonResponse))
	}))

	return server
}

func GetServerResponse(file string) []byte {

	//absPath, _ := filepath.Abs("./../../../test/server_responses/test_cmc_list_response.json")
	path := fmt.Sprintf("./../../../test/server_responses/%s", file)
	absPath, _ := filepath.Abs(path)
	testJsonResponse, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("There was an error attempting to pull in json server response: ", file)
		fmt.Println(err.Error())
	}

	return testJsonResponse
}
