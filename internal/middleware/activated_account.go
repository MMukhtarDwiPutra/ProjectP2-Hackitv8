package middleware

import (
	"P2-Hacktiv8/repository"
	"P2-Hacktiv8/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"gorm.io/gorm"
)

func CheckUserActivationByEmail(userRepo repository.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			body, _ := ioutil.ReadAll(c.Request().Body)
            c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

            var req entity.LoginRequest
            if err := json.Unmarshal(body, &req); err != nil || req.Email == "" || req.Password == ""{
                return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": http.StatusBadRequest, "message": "invalid email or password in request"})
            }

			// Fetch the user from the repository by email.
			user, err := userRepo.GetUserByEmail(req.Email)
			if err != nil {
				if err == gorm.ErrRecordNotFound{
					return c.JSON(http.StatusNotFound, map[string]interface{}{
						"status":  http.StatusNotFound,
						"message": "Email or password is incorrect!",
					})
				}

				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Error query",
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
