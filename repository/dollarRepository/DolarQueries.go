package dollarRepository

func insertDollarPSQLQuery() string {
	return `INSERT INTO dollars (dollar_type, buy_price, sell_price) VALUES ($1, $2, $3) RETURNING dollar_id, 
dollar_date, dollar_type, buy_price, sell_price`
}

func getDollarInDatePSQLQuery() string {
	return `
		SELECT dollar_id, dollar_date, dollar_type, buy_price, sell_price
		FROM dollars
		WHERE dollar_type = $1 AND dollar_date::DATE = $2::DATE`
}

func getAllDollarsInDatePSQLQuery() string {
	return `
		SELECT dollar_id, dollar_date, dollar_type, buy_price, sell_price
		FROM dollars
		WHERE dollar_date::DATE = $1::DATE`
}

func getDollarInDateRangePSQLQuery() string {
	return `
	SELECT dollar_id, dollar_date, dollar_type, buy_price, sell_price
	FROM dollars
	WHERE dollar_type = $1 AND (dollar_date::DATE BETWEEN $2::DATE AND $2::DATE)`

}
