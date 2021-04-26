package cmcmodule

import (
	"io/ioutil"
	"net/http"
)

const host = "https://pro-api.coinmarketcap.com"

func (mm *CmcModule) MakeRequest(method, path string, params map[string]string, payload interface{}) (string, []byte, error) {

	l := mm.l.Lg.With().Str("cmcmodule", "MakeRequest").Logger()

	client := &http.Client{}

	l.Info().Msgf("cmcmodule.MakeRequest: %s %s\n", method, path)

	req, err := http.NewRequest(method, host+path, nil)
	if err != nil {
		l.Error().Msg("There was an error instantiating request client for cmcmodule")
		l.Error().Msg(err.Error())
		return "", nil, err
	}

	req.Header.Add("X-CMC_PRO_API_KEY", mm.ApiKey)
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
		l.Error().Msg("There was an error making request to cmc")
		l.Error().Msg(err.Error())
		return "", nil, err
	}

	statusCode := resp.Status
	respBody, _ := ioutil.ReadAll(resp.Body)

	return statusCode, respBody, nil
}
