package database

import "context"

const table_weather = `
CREATE TABLE IF NOT EXISTS weather
(id SERIAL PRIMARY KEY,
date DATE NOT NULL)
`

var tables = [...]string{table_weather}

func (db *DataBase) CreateTables() error {
	ctx := context.Background()
	for _, table := range tables {
		_, err := db.pool.Exec(ctx, table)
		if err != nil {
			return err
		}
	}
	return nil
}
