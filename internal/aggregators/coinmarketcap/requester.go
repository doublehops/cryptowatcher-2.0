package coinmarketcap

import (
	"fmt"
	"io"
	"net/http"
)

// MakeRequest will make an HTTP request to CoinMarketCap.
func (r *Runner) MakeRequest(method, path string, params map[string]string, payload interface{}) (string, []byte, error) {

	l := r.l.Lg.With().Str(packageName, "MakeRequest").Logger()

	l.Info().Msgf("coinmarketcap.MakeRequest: %s %s", method, path)

	req, err := http.NewRequest(method, r.aggregatorConfig.HostConfig.ApiHost+path, nil)
	if err != nil {
		l.Error().Msg("There was an error instantiating request client for coinmarketcap")
		l.Error().Msg(err.Error())
		return "", nil, err
	}

	req.Header.Add("X-CMC_PRO_API_KEY", r.aggregatorConfig.HostConfig.ApiKey)
	req.Header.Add("Accept", "application/json")

	if params != nil {
		q := req.URL.Query()
		for key, value := range params {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := r.client.Do(req)
	if err != nil {
		errMsg := fmt.Errorf("there was an error making request to cmc. %w", err)
		l.Error().Msg(errMsg.Error())

		return "", nil, errMsg
	}
	defer resp.Body.Close()

	statusCode := resp.Status
	respBody, _ := io.ReadAll(resp.Body)

	return statusCode, respBody, nil
}
