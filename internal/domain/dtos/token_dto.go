package dtos

type TokenDTO struct {
	Token string `json:"token"`
}

func NewTokenDTO(token string) *TokenDTO {
	return &TokenDTO{
		Token: token,
	}
}
