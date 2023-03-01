package aggregators

import (
	"github.com/doublehops/cryptowatcher-2.0/internal/dbinterface"
	"github.com/doublehops/cryptowatcher-2.0/internal/pkg/logga"
	"github.com/doublehops/cryptowatcher-2.0/internal/types/database"
)

type Aggregator interface {
	New() *Runner
	FetchLatestHistory() (database.History, error)
	GetAggregatorName() string
	GetAggregatorID() uint32
}

type Runner struct {
	L  *logga.Logga
	DB dbinterface.QueryAble
}
