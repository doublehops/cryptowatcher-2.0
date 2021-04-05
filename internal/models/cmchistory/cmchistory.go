package cmchistory

import (
	"gorm.io/gorm"

	"cryptowatcher.example/internal/pkg/logga"
	"cryptowatcher.example/internal/types/database"
)

type Model struct {
	db *gorm.DB
	l *logga.Logga
}

func New(db *gorm.DB, logger *logga.Logga) *Model {

	return &Model{
		db: db,
		l: logger,
	}
}

func (m *Model) CreateRecord(r *database.CmcHistory) *gorm.DB {

	l := m.l.Lg.With().Str("cmchistory", "CreateRecord").Logger()

	l.Info().Msgf("Adding cmc record: %s", r.Symbol)

	result := m.db.Create(&r)
	if result.Error != nil {
		l.Error().Msgf("There was an error saving record to database. %v", result.Error)
	}

	return result
}
