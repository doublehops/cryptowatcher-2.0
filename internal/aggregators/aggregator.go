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
	//cfg  config.CMCAggregator
	L  *logga.Logga
	DB dbinterface.QueryAble
	//cmcm *cmcmodule.CmcModule
}
