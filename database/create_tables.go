package database

import "context"

// SQL request for creating table City
const tableCity = `
CREATE TABLE IF NOT EXISTS City
(Id SERIAL PRIMARY KEY,
Name CHARACTER VARYING(40) UNIQUE NOT NULL,
Latitude DECIMAL,
Longitude DECIMAL
);
`

// SQL request for creating table Weather
const tableWeather = `
CREATE TABLE IF NOT EXISTS Weather
(Id SERIAL PRIMARY KEY,
CityId INTEGER,
Date DATE NOT NULL,
MinTemp DECIMAL,
MaxTemp DECIMAL,
FOREIGN KEY (CityId) REFERENCES City (Id),
UNIQUE (CityId, Date)
);
`

var tables = [...]string{tableWeather, tableCity}

var cities = [...]City{
	{"Москва", 55.7, 37.53},
	{"Санкт-Петербург", 59.94, 30.31},
	{"Казань", 55.8, 49.11},
	{"Новосибирск", 55.0, 82.92},
	{"Воронеж", 51.66, 39.2},
}

// SQL request for initializing table City
const initTableCity = `
INSERT INTO City (Name, Latitude, Longitude)
SELECT * FROM UNNEST(
	$1::text[],
	$2::decimal[],
	$3::decimal[]
)
ON CONFLICT (Name)
DO NOTHING
`

// Create tables in database
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

// Initialise tables in database
func (db *DataBase) InitTables() error {
	ctx := context.Background()
	names := make([]string, 0, len(cities))
	latitudes := make([]float64, 0, len(cities))
	longitudes := make([]float64, 0, len(cities))
	for _, city := range cities {
		names = append(names, city.Name)
		latitudes = append(latitudes, city.Latitude)
		longitudes = append(longitudes, city.Longitude)
	}
	_, err := db.pool.Exec(ctx, initTableCity,
		names, latitudes, longitudes)
	return err
}
