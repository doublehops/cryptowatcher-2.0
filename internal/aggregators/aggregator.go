package aggregators

import (
	"cryptowatcher.example/internal/dbinterface"
	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
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
