package cmcmodule

type CurrencySortBase struct {
	Name           string
	Rank           int32
	Symbol         string
	MarketCap      float64
	Volume24h      float64
	CapVolumeRatio float64
}

type CurrencyData struct {
	Currencies []*Currency `json:"data"`
}

type Currency struct {
	ID                uint     `json:"id"`
	Name              string   `json:"name"`
	Symbol            string   `json:"symbol"`
	Slug              string   `json:"slug"`
	NumMarketPairs    int32    `json:"num_market_pairs"`
	DateAdded         string   `json:"date_added"`
	Tags              []string `json:"tags"`
	MaxSupply         float64  `json:"max_supply,omitempty"` // It seems that some currencies are missing this property.
	CirculatingSupply float64  `json:"circulating_supply"`
	TotalSupply       float64  `json:"total_supply"`
	CmcRank           int32    `json:"cmc_rank"`
	LastUpdated       string   `json:"last_updated"`
	Quote             Quote    `json:"quote"`
}

type Quote struct {
	USDObj PriceObj `json:"USD"`
}

type PriceObj struct {
	Price                float64 `json:"price"`
	Volume24Hours        float64 `json:"volume_24h"`
	PercentChange1Hour   float64 `json:"percent_change_1h"`
	PercentChange24Hours float64 `json:"percent_change_24h"`
	PercentChange7Days   float64 `json:"percent_change_7d"`
	PercentChange30Days  float64 `json:"percent_change_30d"`
	PercentChange60Days  float64 `json:"percent_change_60d"`
	PercentChange90Days  float64 `json:"percent_change_90d"`
	MarketCap            float64 `json:"market_cap"`
}

//{
//"id": 6951,
//"name": "Reef",
//"symbol": "REEF",
//"slug": "reef",
//"num_market_pairs": 53,
//"date_added": "2020-09-08T00:00:00.000Z",
//"tags": [
//	"defi",
//	"substrate",
//	"polkadot",
//	"dot-ecosystem"
//],
//"max_supply": 20000000000,
//"circulating_supply": 11268898338,
//"total_supply": 15934019762,
//"platform": {
//	"id": 1027,
//	"name": "Ethereum",
//	"symbol": "ETH",
//	"slug": "ethereum",
//	"token_address": "0xfe3e6a25e6b192a42a44ecddcd13796471735acf"
//},
//"cmc_rank": 99,
//"last_updated": "2021-03-12T23:55:13.000Z",
//"quote": {
//	"USD": {
//		"price": 0.04585198755086,
//		"volume_24h": 560482777.3297944,
//		"percent_change_1h": -3.3673205,
//		"percent_change_24h": 23.71972341,
//		"percent_change_7d": 15.65177067,
//		"percent_change_30d": 16.99732327,
//		"percent_change_60d": 470.25121697,
//		"percent_change_90d": 0,
//		"market_cap": 516701386.30588293,
//		"last_updated": "2021-03-12T23:55:13.000Z"
//		}
//	}
//}
