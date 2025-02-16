package middlewares

import (
	"context"
	"net/http"
	"pasour/internal/domain/services"
	"pasour/internal/domain/types"
	"pasour/internal/interfaces/utils"
)

func AuthMiddleware(next http.Handler, tokenSrv services.TokenService, userSrv services.UserService) http.Handler {

	// authHandler is a middleware that checks if the request has a valid token
	// and if the user exists in the database. If the token is valid and the user exists
	// the userDTO is added to the context and the next handler is called.
	authHandler := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		token, err := tokenSrv.GetTokenFromHeader(authHeader)
		if err != nil {
			utils.WriteErrResponse(err, w)
			return
		}

		sub, err := tokenSrv.Decode(token)
		if err != nil {
			utils.WriteErrResponse(err, w)
			return
		}

		user, err := userSrv.FindByUsername(sub)
		if err != nil {
			utils.WriteErrResponse(err, w)
			return
		}

		// add userDTO to context
		ctx := r.Context()
		ctx = context.WithValue(ctx, types.UserCtxKey, user)
		r = r.WithContext(ctx)

		// call next handler
		next.ServeHTTP(w, r)

	}
	return http.HandlerFunc(authHandler)
}
