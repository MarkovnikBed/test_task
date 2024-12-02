package repository

import (
	_ "github.com/lib/pq"
)

func (rep *Repository) PrepareTable() error {
	_, err := rep.DB.Exec("CREATE TABLE IF NOT EXISTS Refresh_tokens (id INTEGER UNIQUE, hash VARCHAR)")
	if err != nil {
		return err
	}
	return nil
}
