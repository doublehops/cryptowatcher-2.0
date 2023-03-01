package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/doublehops/cryptowatcher-2.0/internal/dbinterface"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/types/database"
)

const (
	packageName = "coinmarketcap"

	aggregatorID   uint32 = 1
	aggregatorName string = "coinmarketcap"
)

type Runner struct {
	l                *logga.Logga
	db               dbinterface.QueryAble
	aggregatorConfig *aggregatorConfig
	client           *http.Client
}

type aggregatorConfig struct {
	Name       string     `json:"name"`
	Label      string     `json:"label"`
	HostConfig HostConfig `json:"hostConfig"`
}

type HostConfig struct {
	APIKey  string `json:"apiKey"`
	APIHost string `json:"apiHost"`
}

// New will instantiate Runner.
func New(l *logga.Logga, db dbinterface.QueryAble, client *http.Client) (*Runner, error) {
	config, err := parseConfig("./internal/aggregators/coinmarketcap/config.json")
	if err != nil {
		// todo - add log message
		return nil, err
	}

	if config.HostConfig.APIHost == "" || config.HostConfig.APIKey == "" {
		return nil, fmt.Errorf("coinmarketcap configuration not set")
	}

	return &Runner{
		aggregatorConfig: config,
		l:                l,
		db:               db,
		client:           client,
	}, nil
}

func parseConfig(configFile string) (*aggregatorConfig, error) {
	var config aggregatorConfig
	absPath, _ := filepath.Abs(configFile)
	f, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file `%s`. %w", configFile, err)
	}

	if err := json.Unmarshal(f, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// FetchLatestHistory will fetch the latest history populate a database.History struct.
func (r *Runner) FetchLatestHistory() ([]*database.History, error) {
	var histories []*database.History

	l := r.l.Lg.With().Str(packageName, "FetchLatestHistory").Logger()
	l.Info().Msg("Running currency fetcher")

	currencies, err := r.FetchCurrencyListing(20)
	if err != nil {
		r.l.Error("Unable to get currency listing from CMC module")
		return histories, err
	}

	for _, c := range currencies {

		history := &database.History{
			AggregatorID:      r.GetAggregatorID(),
			Name:              c.Name,
			Symbol:            c.Symbol,
			Slug:              c.Slug,
			NumMarketPairs:    c.NumMarketPairs,
			DateAdded:         c.DateAdded,
			MaxSupply:         c.MaxSupply,
			CirculatingSupply: c.CirculatingSupply,
			TotalSupply:       c.TotalSupply,
			Rank:              c.CmcRank,
			QuotePrice:        c.Quote.USDObj.Price,
			Volume24h:         c.Quote.USDObj.Volume24Hours,
			PercentChange1h:   c.Quote.USDObj.PercentChange1Hour,
			PercentChange24h:  c.Quote.USDObj.PercentChange24Hours,
			PercentChange7D:   c.Quote.USDObj.PercentChange7Days,
			PercentChange30D:  c.Quote.USDObj.PercentChange30Days,
			PercentChange60D:  c.Quote.USDObj.PercentChange60Days,
			PercentChange90D:  c.Quote.USDObj.PercentChange90Days,
			MarketCap:         c.Quote.USDObj.MarketCap,
		}

		histories = append(histories, history)
	}

	return histories, nil
}

// GetAggregatorName will return the aggregator name for log messages
func (r *Runner) GetAggregatorName() string {
	return aggregatorName
}

// GetAggregatorID will return the aggregator ID to distinguish which aggregator each currency record came from.
func (r *Runner) GetAggregatorID() uint32 {
	return aggregatorID
}
