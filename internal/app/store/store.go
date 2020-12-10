package store

import (
	_ "github.com/go-sql-driver/mysql" //...
	"github.com/jmoiron/sqlx"
)

//Store ...
type Store struct {
	config *Config
	db     *sqlx.DB
}

//New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
		db:     &sqlx.DB{},
	}
}

//Open ...
func (s *Store) Open() error {
	db, err := sqlx.Connect("mysql", s.config.DB_url)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

//Close ...
func (s *Store) Close() {
	s.db.Close()
}
