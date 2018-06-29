package pgeo

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/volatiletech/sqlboiler/randomize"
)

type NullPoint struct {
	Point
	Valid bool
}

func (p NullPoint) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return valuePoint(p.Point)
}

func (p *NullPoint) Scan(src interface{}) error {
	if src == nil {
		p.Point, p.Valid = NewPoint(0, 0), false
		return nil
	}

	p.Valid = true
	return scanPoint(&p.Point, src)
}

func (p *NullPoint) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(p.Point)
}

func (p *NullPoint) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		p.Valid = false
		return nil
	}

	var err = json.Unmarshal(data, p.Point)
	p.Valid = err == nil
	return err
}

func (o *NullPoint) Randomize(seed *randomize.Seed, fieldType string, shouldBeNull bool) {
	*o = NewNullPoint(NewRandPoint(), !shouldBeNull)
}
