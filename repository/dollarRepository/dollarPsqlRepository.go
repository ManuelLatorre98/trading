package dollarRepository

import (
	"database/sql"
	"manulatorre98/trading/graph/model"
)

type DollarPsqlRepository struct {
	db *sql.DB
}

func NewDollarPsqlRepository(db *sql.DB) *DollarPsqlRepository {
	return &DollarPsqlRepository{db: db}
}

func (r *DollarPsqlRepository) CreateDollar(dollar *model.DollarInput) (*model.Dollar, error) {
	var newDolar model.Dollar
	err := r.db.QueryRow(insertDollarPSQLQuery(), dollar.DolarType, dollar.BuyPrice, dollar.SellPrice).Scan(
		&newDolar.DollarID,
		&newDolar.DollarDate,
		&newDolar.DollarType,
		&newDolar.BuyPrice,
		&newDolar.SellPrice,
	)
	return &newDolar, err
}
func (r *DollarPsqlRepository) getDollarInDate(dollarType string, date string) (*model.Dollar, error) {
	var dollar model.Dollar
	err := r.db.QueryRow(getDollarInDatePSQLQuery(), dollarType, date).Scan(
		&dollar.DollarID,
		&dollar.DollarDate,
		&dollar.DollarType,
		&dollar.BuyPrice,
		&dollar.SellPrice,
	)
	return &dollar, err
}
func (r *DollarPsqlRepository) getAllDollarsInDate(date string) (*model.Dollar, error) {
	//todo fix this to return an array of dollars
	var dollar model.Dollar
	err := r.db.QueryRow(getAllDollarsInDatePSQLQuery(), date).Scan(
		&dollar.DollarID,
		&dollar.DollarDate,
		&dollar.DollarType,
		&dollar.BuyPrice,
		&dollar.SellPrice,
	)
	return &dollar, err
}

func (r *DollarPsqlRepository) getDollarInDateRange(dollarType string, startDate string,
	endDate string) ([]*model.Dollar, error) {

	rows, err := r.db.Query(getDollarInDateRangePSQLQuery(), dollarType, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var dollars []*model.Dollar

	for rows.Next() {
		var dollar model.Dollar
		if err := rows.Scan(
			&dollar.DollarID,
			&dollar.DollarDate,
			&dollar.DollarType,
			&dollar.BuyPrice,
			&dollar.SellPrice,
		); err != nil {
			return nil, err
		}
		dollars = append(dollars, &dollar)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dollars, nil
}
func (r *DollarPsqlRepository) getAllDollarsInDateRange(startDate string, endDate string) ([]*model.Dollar, error) {
	panic("not implemented")
}

func (r *DollarPsqlRepository) getHistoricalAllDollar(dollarType string) ([]*model.Dollar, error) {
	panic("not implemented")
}
func (r *DollarPsqlRepository) getHistoricalAllDollars() ([]*model.Dollar, error) {
	panic("not implemented")
}
