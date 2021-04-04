package database

type Coin struct {
	ID        uint
	Name      string
	Symbol    string
	CreatedAt int
	UpdatedAt int
}

type CmcHistory struct {
	ID                uint
	Name              string
	Symbol            string
	NumMarketPairs    int
	DateAdded         string
	MaxSupply         int
	CirculatingSupply int
	TotalSupply       int
	CmcRank           int
	QuotePrice        float32
	Volume24h         float32
	PercentChange1h   float32
	PercentChange24h  float32
	PercentChange7D   float32
	PercentChange30D  float32
	PercentChange60D  float32
	PercentChange90D  float32
	MarketCap         float32
	CreatedAt         int
	UpdatedAt         int
}
