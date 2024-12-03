package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	configs "github.com/MatthewAraujo/min-ecommerce/config"
	"github.com/MatthewAraujo/min-ecommerce/repository"
	"github.com/MatthewAraujo/min-ecommerce/utils"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserKey contextKey = "userID"

var logger = utils.NewParentLogger("AUTH")

func WithJWTAuth(handleFunc http.HandlerFunc, store repository.Queries, requiredRole string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("VALIDATE JWT")
		tokenString := getTokenFromRequest(r)

		token, err := validateJWT(tokenString)
		if err != nil {
			logger.Info("error validating token: %v", err.Error())
			permissionDenied(w)
			return
		}

		if !token.Valid {
			logger.Info("token is invalid")
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userIDValue, ok := claims["userID"]
		if !ok {
			logger.Info("userID not found in claims")
			permissionDenied(w)
			return
		}

		userID, ok := userIDValue.(float64)
		if !ok {
			logger.Info("userID is not a number")
			permissionDenied(w)
			return
		}

		roleValue, ok := claims["role"]
		if !ok {
			logger.Info("role not found in claims")
			permissionDenied(w)
			return
		}

		role, ok := roleValue.(string)
		if !ok {
			logger.Info("role is not a string")
			permissionDenied(w)
			return
		}

		if role != requiredRole {
			logger.Info("user does not have the required role: %v", requiredRole)
			permissionDenied(w)
			return
		}

		u, err := store.FindCustomerByID(context.Background(), int32(userID)) // Converta para int32
		if err != nil {
			logger.Info("error fetching user: %v", err.Error())
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handleFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	if tokenAuth == "" {
		return ""
	}

	tokenParts := strings.Split(tokenAuth, " ")
	if len(tokenParts) == 2 {
		return tokenParts[1]
	}

	return ""
}

func CreateJWT(secret []byte, userID int32, role string) (string, error) {
	expiration := time.Second * time.Duration(configs.Envs.JWT.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":  userID,
		"role":    role,
		"expires": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(configs.Envs.JWT.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}
