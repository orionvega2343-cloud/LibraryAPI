package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { //Closure
		tokenStr := r.Header.Get("Authorization") //Get key of token
		if tokenStr == "" {                       //Validate token
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		parseToken := strings.TrimPrefix(tokenStr, "Bearer ")                                                       //Unprefix token
		validToken, err := jwt.ParseWithClaims(parseToken, &Claims{}, func(token *jwt.Token) (interface{}, error) { //Get token value from Claims
			return []byte(secret), nil //Returning parsed value
		})
		if err != nil || !validToken.Valid { //If token not nil and not valid
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", validToken.Claims.(*Claims)) //Post value on the context
		next.ServeHTTP(w, r.WithContext(ctx))                                      //Delegation value to the handlers
	})

}
