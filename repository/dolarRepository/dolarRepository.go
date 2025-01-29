package dolarRepository

import "manulatorre98/trading/graph/model"

type DolarRepository interface {
	CreateDolar(dolar *model.DolarInput) (*model.Dolar, error)
	getDolar(date string) (*model.Dolar, error)
	getDolarInDateRange(startDate string, endDate string) ([]*model.Dolar, error)
	getHistoricalDolar() ([]*model.Dolar, error)
}
