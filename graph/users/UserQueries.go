package users

func insertUserPSQLQuery() string {
	return "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING user_id, username, email, registerDate"
}
func getUserByEmailPSQLQuery() string {
	return "SELECT user_id, username, email, registerDate FROM users WHERE email = $1"
}

func getUserByIdPSQLQuery() string {
	return "SELECT user_id, username, email, registerDate FROM users WHERE user_id = $1"
}
