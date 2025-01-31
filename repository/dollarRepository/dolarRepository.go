package dollarRepository

import "manulatorre98/trading/graph/model"

type DollarRepository interface {
	CreateDollar(dollar *model.DollarInput) (*model.Dollar, error)
	getDollarInDate(dollarType string, date string) (*model.Dollar, error)
	getAllDollarsInDate(date string) (*model.Dollar, error)
	getDollarInDateRange(dollarType string, startDate string, endDate string) ([]*model.Dollar, error)
	getAllDollarsInDateRange(startDate string, endDate string) ([]*model.Dollar, error)
	getHistoricalAllDollar(dollarType string) ([]*model.Dollar, error)
	getHistoricalAllDollars() ([]*model.Dollar, error)
}
