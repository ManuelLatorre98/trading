package dolarRepository

import (
	"database/sql"
	"manulatorre98/trading/graph/model"
)

type DolarPsqlRepository struct {
	db *sql.DB
}

func NewDolarPsqlRepository(db *sql.DB) *DolarPsqlRepository {
	return &DolarPsqlRepository{db: db}
}

func (r *DolarPsqlRepository) CreateDolar(dolar *model.DolarInput) (*model.Dolar, error) {
	var newDolar model.Dolar
	err := r.db.QueryRow(insertDolarPSQLQuery(), dolar.PriceBlue, dolar.PriceOfficial, dolar.PriceMep).Scan(
		&newDolar.DolarDate,
		&newDolar.PriceBlue,
		&newDolar.PriceOfficial,
		&newDolar.PriceMep,
	)
	return &newDolar, err
}

func (r *DolarPsqlRepository) getDolar(date string) (*model.Dolar, error) {
	panic("not implemented")
}

func (r *DolarPsqlRepository) getDolarInDateRange(startDate string, endDate string) ([]*model.Dolar, error) {
	panic("not implemented")
}

func (r *DolarPsqlRepository) getHistoricalDolar() ([]*model.Dolar, error) {
	panic("not implemented")
}
