package testfuncs

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/config"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
)

// GetTestConfig will return the config of the test config file.
// todo - Github is not able to open the path for some reason so hard-coding the response for now.
func GetTestConfig(l *logga.Logga) (*config.Config, error) {
	//cfg, err := config.New(l, "./../../../test/config.json.test")
	//if err != nil {
	//	l.Lg.Error().Msgf("error starting main. %w", err.Error())
	//	return nil, err
	//}

	cfg := &config.Config{
		Aggregator: config.Aggregator{
			Name: "coingecko",
		},
		DB: config.DB{
			Host: "127.0.0.1",
			Name: "cw_test",
			User: "dev",
			Pass: "pass12",
		},
	}

	return cfg, nil
}

// --------  TEST SERVER  --------

// SetupTestServer will setup a test server and respond with the value supplied as `jsonResponse`.
// nolint:errcheck
func SetupTestServer(jsonResponse []byte) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(jsonResponse))
	}))

	return server
}

// GetTestJsonResponse will return the contents of `file` after reading it from the `serverresponses` directory.
// This directory should include all test responses we need in the application.
func GetTestJsonResponse(file string) ([]byte, error) {
	var res []byte
	path := "../../../test/serverresponses/" + file
	absPath, _ := filepath.Abs(path)
	testJsonResponse, err := os.ReadFile(absPath)
	if err != nil {
		return res, err
	}

	return testJsonResponse, nil
}

// --------  TEST CLIENT  --------

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// GetNewTestClient can be called to return a http.Client to use in tests. Just pass in the response you want back.
func GetNewTestClient(testJsonResponse []byte) *http.Client {
	return NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(string(testJsonResponse))),
			Header:     make(http.Header),
		}
	})
}
