package middleware

import (
	"P2-Hacktiv8/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"encoding/json"
	"io/ioutil"
	"bytes"
)

func CheckUserActivationByEmail(userRepo repository.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			body, _ := ioutil.ReadAll(c.Request().Body)
            c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

            type Request struct {
                Email string `json:"email"`
            }
            var req Request
            if err := json.Unmarshal(body, &req); err != nil || req.Email == "" {
                return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid email in request"})
            }

			// Fetch the user from the repository by email.
			user, err := userRepo.GetUserByEmail(req.Email)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Failed to fetch user",
				})
			}

			// Check if the user exists and is activated.
			if user == nil || user.IsActivated != "Activated"{
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"status":  http.StatusForbidden,
					"message": "User is not activated",
				})
			}

			// Proceed to the next handler.
			return next(c)
		}
	}
}
