package handlers

import (
	"net/http"
	"pasour/internal/domain/commands"
	"pasour/internal/domain/services"
	"pasour/internal/interfaces/utils"
)

type UserHandler struct {
	UserService services.UserService
	Router      *http.ServeMux
}

func NewUserHandler(userService services.UserService, router *http.ServeMux) *UserHandler {
	return &UserHandler{
		UserService: userService,
		Router:      router,
	}
}

func (h *UserHandler) RegisterRoutes() {
	h.Router.HandleFunc("POST /signup", func(w http.ResponseWriter, r *http.Request) {
		signupHandler(w, r, h.UserService)
	})
}

func signupHandler(w http.ResponseWriter, r *http.Request, userSrv services.UserService) {
	cmd := &commands.UserSignUpCmd{}

	if err := utils.ValidateReqBody(r, cmd); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err)
		return
	}

	user, err := userSrv.SignUp(cmd)

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	utils.RespondJSON(w, http.StatusCreated, user)
}
