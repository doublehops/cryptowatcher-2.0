package cmcmodule

import (
	"io/ioutil"
	"net/http"
)

const host = "https://pro-api.coinmarketcap.com"

//type Requester struct {
//	ApiKey string
//	l *logga.Logga
//}

//func New(rApiKey string, logger *logga.Logga) *Requester {
//	return &Requester{
//		ApiKey: rApiKey,
//		l: logger,
//	}
//}

func (mm *CmcModule) MakeRequest(method, path string, params map[string]string, payload interface{}) (string, []byte, error) {

	l := mm.l.Lg.With().Str("cmcmodule", "MakeRequest").Logger()

	client := &http.Client{}

	l.Info().Msgf("r.MakeRequest: %s %s\n", method, path)

	req, err := http.NewRequest(method, host+path, nil)
	if err != nil {
		l.Error().Msg("There was an error instantiated request client for r")
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
		l.Error().Msg("There was an error making request to r")
		l.Error().Msg(err.Error())
		return "", nil, err
	}

	statusCode := resp.Status
	respBody, _ := ioutil.ReadAll(resp.Body)

	return statusCode, respBody, nil
}
