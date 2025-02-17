package handlers

import (
	"net/http"
	"pasour/internal/domain/commands"
	"pasour/internal/domain/services"
	"pasour/internal/interfaces/utils"
	"strconv"
)

type AuthHandler struct {
	userSrv  services.UserService
	tokenSrv services.TokenService
	router   *http.ServeMux
}

func NewAuthHandler(userSrv services.UserService, tokenSrv services.TokenService, router *http.ServeMux) *AuthHandler {
	return &AuthHandler{
		userSrv:  userSrv,
		tokenSrv: tokenSrv,
		router:   router,
	}
}

func (handler *AuthHandler) RegisterRoutes() {
	handler.router.HandleFunc("POST /auth", handler.authenticateUser)
}

func (handler *AuthHandler) authenticateUser(w http.ResponseWriter, r *http.Request) {
	authCmd := &commands.AuthCmd{}
	err := utils.ValidateReqBody(r, authCmd)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	user, err := handler.userSrv.Authenticate(authCmd)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	encodeCmd := &commands.TokenEncodeCmd{
		Sub: strconv.Itoa(user.ID),
	}

	token, err := handler.tokenSrv.Encode(encodeCmd)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondJSON(w, http.StatusOK, token)
}
