package middleware

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"github.com/mojeico/gqlgen-golang/internal/model"
	"github.com/mojeico/gqlgen-golang/internal/repository"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

var CurrentUserKey = "currentUser"

const (
	tokenTime = 24 * 7 * time.Hour // 7 days for token
	signInKey = "registerKey"
)

func AuthMiddleware(userRepo repository.UserRepo) func(handler http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token, err := parseToken(r)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			user, err := userRepo.GetUserByID(claims["jti"].(string))

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

}

//Bearer test_token12313131313131
func stripBearerPrefixFromToken(token string) (string, error) {

	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

func parseToken(r *http.Request) (*jwt.Token, error) {

	authHeaderExtractor := &request.PostExtractionFilter{
		Extractor: request.HeaderExtractor{"Authorization"},
		Filter:    stripBearerPrefixFromToken,
	}

	authExtractor := &request.MultiExtractor{
		authHeaderExtractor,
		request.ArgumentExtractor{"access_token"},
	}

	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(signInKey)
		return t, nil
	})

	return jwtToken, errors.Wrap(err, "parseToken error: ")

}

func GetCurrentUserFromContext(ctx context.Context) (*model.User, error) {

	errNoUserInContext := "no user in context"

	if ctx.Value(CurrentUserKey) == nil {
		return nil, errors.New(errNoUserInContext)
	}

	user, ok := ctx.Value(CurrentUserKey).(*model.User)
	if !ok || user.ID == "" {
		return nil, errors.New(errNoUserInContext)
	}

	return user, nil
}
