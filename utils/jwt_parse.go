package utils

import (
	"errors"
	"fmt"
	"time"
	"os"

	"github.com/golang-jwt/jwt/v4"
)


func ParseRegisTokenString(tokenString string) (map[string]string, error) {
	secretKey := os.Getenv("REGIST_SECRET_KEY")

	// Parse the token using the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		// Return error if token parsing fails
		return nil, fmt.Errorf("token error: %w", err)
	}

	// Validate the token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration time of the token
		exp, ok := claims["exp"].(float64)
		if !ok {
			return nil, errors.New("invalid expiration time")
		}

		// Ensure the token has not expired
		if time.Unix(int64(exp), 0).Before(time.Now()) {
			return nil, errors.New("token has expired")
		}

		// Retrieve the email claim
		email, ok := claims["email"].(string)
		if !ok {
			return nil, errors.New("email not found in token")
		}

		// Return the parsed token details
		return map[string]string{
			"email": email,
		}, nil
	}

	// Return an error if token claims are invalid
	return nil, errors.New("invalid token claims")
}
