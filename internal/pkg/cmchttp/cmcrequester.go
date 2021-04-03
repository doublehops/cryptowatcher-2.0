package cmchttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const host = "https://pro-api.coinmarketcap.com"

type Requester struct {
	ApiKey           string
}

func New(rApiKey string) *Requester {
	return &Requester{
		ApiKey:           rApiKey,
	}
}

func (r *Requester) MakeRequest(method, path string, params map[string]string, payload interface{}) (string, []byte, error) {

	client := &http.Client{}

	fmt.Printf("r.MakeRequest: %s %s\n", method, path)

	req, err := http.NewRequest(method, host + path, nil)
	if err != nil {
		fmt.Println("There was an error instantiated request client for r")
		fmt.Println(err.Error())
		return "", nil, err
	}

	req.Header.Add("X-CMC_PRO_API_KEY", r.ApiKey)
	req.Header.Add("Accept", "application/json")

	if params != nil {
		q := req.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("There was an error making request to r")
		fmt.Println(err.Error())
		return "", nil, err
	}

	statusCode := resp.Status
	fmt.Printf("----- statusCode: %s\n", statusCode)
	respBody, _ := ioutil.ReadAll(resp.Body)

	return statusCode, respBody, nil
}