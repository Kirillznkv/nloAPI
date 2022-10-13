package store

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Store struct {
	config            *Config
	db                *sql.DB
	anomalyRepository *AnomalyRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) MigrateUP() error {
	m, err := migrate.New(
		s.config.MigrateDirPath,
		s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Anomaly() *AnomalyRepository {
	if s.anomalyRepository != nil {
		return s.anomalyRepository
	}
	s.anomalyRepository = &AnomalyRepository{
		store: s,
	}
	return s.anomalyRepository
}
