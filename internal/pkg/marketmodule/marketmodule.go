package marketmodule

import (
	"encoding/json"
	"fmt"
	"strconv"

	"cryptowatcher.example/internal/pkg/cmchttp"
	"cryptowatcher.example/internal/types"
)

type marketmodule struct {
	ApiKey string
	cmcr	cmchttp.Requester
}

func New(mmApiKey string) *marketmodule {

	return &marketmodule{
		ApiKey:           mmApiKey,
	}
}

func (mm *marketmodule) SaveCurrencyListing(numberToRetrieve int) (string, error) {

	fmt.Println("---  Fetching currencies  ---")
	fmt.Printf("Fetching top %d but dispalying top %d\n\n", numberToRetrieve)

	currencies, _ := mm.GetCurrencyListing(numberToRetrieve)

	for _, c := range currencies.Currencies {

		ratio := c.Quote.USDObj.MarketCap / c.Quote.USDObj.Volume24Hours

		_ = types.CurrencySortBase{
			Name:           c.Name,
			Rank:           c.CmcRank,
			Symbol:         c.Symbol,
			MarketCap:      c.Quote.USDObj.MarketCap,
			Volume24h:      c.Quote.USDObj.Volume24Hours,
			CapVolumeRatio: ratio,
		}

		// @TODO: Save to database
	}

	return "", nil
}

func (mm *marketmodule) GetCurrencyListing(limit int) (*types.CurrencyListing, error) {

	params := map[string]string{
		"start":   "1",
		"limit":   strconv.Itoa(limit),
		"convert": "USD",
	}

	cmcr := cmchttp.New(mm.ApiKey)

	_, data, err := cmcr.MakeRequest("GET", "/v1/cryptocurrency/listings/latest", params, nil)
	if err != nil {
		fmt.Println("There was an error instantiating mm request client")
		fmt.Println(err.Error())
		return nil, err
	}

	var listing types.CurrencyListing
	err = json.Unmarshal(data, &listing)
	if err != nil {
		fmt.Println("There was an error unmarshalling json mm response")
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Printf("%d currencies returned\n\n", len(listing.Currencies))

	return &listing, nil
}
