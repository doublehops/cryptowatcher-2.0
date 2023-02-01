package cgmodule

import (
	"fmt"

	"io/ioutil"
	"net/http"
)

// MakeRequest will make an HTTP request to CoinMarketCap.
func (mm *RequestModule) MakeRequest(method, path string, params map[string]string, payload interface{}) (string, []byte, error) {

	l := mm.l.Lg.With().Str("coingeckomodule", "MakeRequest").Logger()

	client := &http.Client{}

	l.Info().Msgf("coingeckomodule.MakeRequest: %s %s", method, path)

	req, err := http.NewRequest(method, mm.ApiHost+path, nil)
	if err != nil {
		l.Error().Msg("There was an error instantiating request client for coingeckomodule")
		l.Error().Msg(err.Error())
		return "", nil, err
	}

	//req.Header.Add("X-CMC_PRO_API_KEY", mm.ApiKey)
	//req.Header.Add("Accept", "application/json")

	if params != nil {
		q := req.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Errorf("there was an error making request to cmc. %w", err)
		l.Error().Msg(errMsg.Error())

		return "", nil, errMsg
	}

	statusCode := resp.Status
	respBody, _ := ioutil.ReadAll(resp.Body)

	return statusCode, respBody, nil
}
