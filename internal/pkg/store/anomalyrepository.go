package store

import "github.com/Kirillznkv/nloAPI/internal/pkg/model"

type AnomalyRepository struct {
	store *Store
}

func (r *AnomalyRepository) Create(a *model.Anomaly) error {
	if err := r.store.db.QueryRow(
		"INSERT INTO anomaly (SessionId, Frequency, Timestamp) VALUES ($1, $2, $3)",
		a.SessionId,
		a.Frequency,
		a.Timestamp,
	); err != nil {
		return err.Err()
	}

	return nil
}
