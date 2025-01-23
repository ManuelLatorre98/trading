package userRepository

func insertUserPSQLQuery() string {
	return "INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING user_id, username, email, " +
		"register_date"
}
func getUserByEmailPSQLQuery() string {
	return "SELECT user_id, username, email, register_date FROM users WHERE email = $1"
}

func getUserByUserNamePSQLQuery() string {
	return "SELECT user_id, username, email, register_date FROM users WHERE username = $1"
}

func getUserByIdPSQLQuery() string {
	return "SELECT user_id, username, email, register_date FROM users WHERE user_id = $1"
}

func getEmailPasswordPSQLQuery() string {
	return "SELECT password FROM users WHERE email = $1"
}
