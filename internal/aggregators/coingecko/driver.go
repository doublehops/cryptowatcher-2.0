package coingecko

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/doublehops/cryptowatcher-2.0/internal/dbinterface"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/types/database"
)

const (
	packageName = "coingecko"

	aggregatorID   uint32 = 2
	aggregatorName string = "coingecko"
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
	config, err := parseConfig("internal/aggregators/coingecko/config.json")
	if err != nil {
		// todo - add log message
		return nil, err
	}

	if config.HostConfig.APIHost == "" {
		return nil, fmt.Errorf("coingecko configuration not set")
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
	f, err := os.ReadFile(configFile)
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
		r.l.Error("Unable to get currency listing from Coingecko module")
		return histories, err
	}

	for _, c := range currencies {
		history := &database.History{
			AggregatorID:      r.GetAggregatorID(),
			Name:              c.Name,
			Symbol:            c.Symbol,
			CirculatingSupply: c.CirculatingSupply,
			TotalSupply:       c.TotalSupply,
			Rank:              c.MarketCapRank,
			QuotePrice:        c.CurrentPrice,
			High24hr:          c.High24H,
			Low24hr:           c.Low24H,
			Volume24h:         float64(c.TotalVolume),
			PercentChange24h:  c.PriceChange24H,
			MarketCap:         float64(c.MarketCap),
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
