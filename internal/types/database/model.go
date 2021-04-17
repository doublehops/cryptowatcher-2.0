package database

type Coin struct {
	ID        uint32
	Name      string
	Symbol    string
	CreatedAt int32
	UpdatedAt int32
}

type CmcHistory struct {
	ID                int32
	Name              string
	Symbol            string
	Slug              string
	NumMarketPairs    int32
	DateAdded         string
	MaxSupply         float64
	CirculatingSupply float64
	TotalSupply       float64
	CmcRank           int32
	QuotePrice        float64
	Volume24h         float64 `gorm:"column:volume_24h"`
	PercentChange1h   float64 `gorm:"column:percent_change_1h"`
	PercentChange24h  float64 `gorm:"column:percent_change_24h"`
	PercentChange7D   float64 `gorm:"column:percent_change_7d"`
	PercentChange30D  float64 `gorm:"column:percent_change_30d"`
	PercentChange60D  float64 `gorm:"column:percent_change_60d"`
	PercentChange90D  float64 `gorm:"column:percent_change_90d"`
	MarketCap         float64
	CreatedAt         int32
	UpdatedAt         int32
}
