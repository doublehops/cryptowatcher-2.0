package database

import (
	"database/sql"
	"time"
)

type Currencies []*Currency

type Currency struct {
	ID        uint32
	Name      string
	Symbol    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// type Histories []*History
type HistoryPriceTimeSeriesData []HistoryPriceTimeSeriesDataItem

type History struct {
	ID                uint32
	AggregatorID      uint32
	CurrencyID        uint32
	Currency          Currency
	Name              string
	Symbol            string
	Slug              string
	NumMarketPairs    int32
	DateAdded         string
	MaxSupply         float64
	CirculatingSupply float64
	TotalSupply       float64
	Rank              int32
	QuotePrice        float64
	High24hr          float64
	Low24hr           float64
	Volume24h         float64
	PercentChange1h   float64
	PercentChange24h  float64
	PercentChange7D   float64
	PercentChange30D  float64
	PercentChange60D  float64
	PercentChange90D  float64
	MarketCap         float64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         sql.NullTime
}

type HistoryPriceTimeSeriesDataItem struct {
	QuotePrice float64
	CreatedAt  time.Time
}
