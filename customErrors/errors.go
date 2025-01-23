package customErrors

import "manulatorre98/trading/graph/directives"

const (
	ErrUserEmailNotFound      = "user email not found"
	ErrUserEmailAlreadyExists = "email already exists"
	ErrUserNameNotFound       = "user name not found"
	ErrUserNameAlreadyExists  = "username already exists"
	ErrUserIdNotFound         = "user id not found"
	ErrInternalServer         = "internal server error"

	ErrEmailOrPassWrong = "Email or password is wrong"
)

func DefaultTranslation() {
	directives.ValidateAddTranslation("required", " is required")
	directives.ValidateAddTranslation("email", " must be a valid email")
	directives.ValidateAddTranslation("min", " must have at least %s characters")
	directives.ValidateAddTranslation("max", " must have at most %s characters")
	directives.ValidateAddTranslation("unique", " already exists")
}
