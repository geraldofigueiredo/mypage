package handler

import (
	"context"
	"net/http"
	"strings"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extração do token da request e armaenamento no contexto
		token := extractHeaderToken(r)
		ctx := r.Context()
		ctx = context.WithValue(ctx, "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractHeaderToken(r *http.Request) string {
	headerAuthString := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerAuthString, " ")
	if headerAuthString == "" || len(bearerToken) == 1 {
		return ""
	}

	return bearerToken[1]
}
