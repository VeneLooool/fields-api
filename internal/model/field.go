package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Coordinates []Coordinate

type Coordinate struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Field struct {
	ID          uint64      `db:"id"`
	Name        string      `db:"name"`
	Culture     string      `db:"culture"`
	CreatedBy   string      `db:"created_by"`
	Coordinates Coordinates `db:"coordinates"`
}

func (cs Coordinates) Value() (driver.Value, error) {
	value, err := json.Marshal(cs)
	if err != nil {
		return nil, err
	}

	return string(value), nil
}

func (cs *Coordinates) Scan(src any) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &cs)
}
