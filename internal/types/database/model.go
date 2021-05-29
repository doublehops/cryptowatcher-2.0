package database

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Currencies []Currency

type Currency struct {
	ID     uint32 `gorm:"primarykey"`
	Name   string
	Symbol string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}

type CmcHistories []CmcHistory

type CmcHistory struct {
	ID                uint32
	CurrencyID        uint32
	Currency          Currency `gorm:"foreignKey:CurrencyID"`
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
	gorm.Model
}
