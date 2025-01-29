package dolarRepository

func insertDolarPSQLQuery() string {
	return `INSERT INTO dolars (price_blue, price_official, price_mep) VALUES ($1, $2, $3) RETURNING dolar_date, price_blue, price_official, price_mep`
}
