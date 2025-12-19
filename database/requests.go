package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type City struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type Day struct {
	Time    string
	TempMin float64
	TempMax float64
}

func (db *DataBase) GetCities() ([]City, error) {
	ctx := context.Background()
	sql := `SELECT name, latitude, longitude FROM City;`
	rows, _ := db.pool.Query(ctx, sql)
	cities, err := pgx.CollectRows(rows, pgx.RowToStructByName[City])
	return cities, err
}

func (db *DataBase) GetNamesOfCities() ([]string, error) {
	ctx := context.Background()
	sql := `SELECT name FROM City;`
	rows, _ := db.pool.Query(ctx, sql)
	names, err := pgx.CollectRows(rows, pgx.RowTo[string])
	return names, err
}

func (db *DataBase) GetCityByName(name string) (City, error) {
	ctx := context.Background()
	sql := `SELECT name, latitude, longitude FROM City
			WHERE name = $1;`
	rows, _ := db.pool.Query(ctx, sql, name)
	city, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[City])
	return city, err
}

const addCity = `
INSERT INTO City (name, latitude, longitude)
VALUES ($1, $2, $3)
ON CONFLICT (name)
DO UPDATE SET
	latitude = EXCLUDED.latitude,
	longitude = EXCLUDED.longitude
`

func (db *DataBase) AddCity(city City) error {
	ctx := context.Background()
	_, err := db.pool.Exec(ctx, addCity, city.Name, city.Latitude, city.Longitude)
	return err
}
