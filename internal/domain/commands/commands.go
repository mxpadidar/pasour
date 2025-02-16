package commands

type UserSignUpCmd struct {
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=8"`
}

type TokenEncodeCmd struct {
	Sub string
}
