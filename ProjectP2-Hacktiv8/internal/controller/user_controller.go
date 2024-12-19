package controller

import (
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/internal/service"
	"fmt"
	"github.com/go-playground/validator/v10" // Validator untuk memvalidasi request body.
	"github.com/labstack/echo/v4"            // Import Echo framework untuk pengelolaan HTTP API.
	"net/http"
)

// Validator global yang digunakan untuk memvalidasi struktur request body.
var validate = validator.New()

type userController struct {
	userService service.UserService
}

// NewUserController creates a new instance of userController.
func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Registers a user by providing their email, full name, and password. A confirmation email will be sent to the provided email address.
// @Tags User
// @Accept json
// @Produce json
// @Param request body entity.RegisterRequest true "Register Request"
// @Success 201 {object} entity.RegisterResponse "User successfully registered"
// @Failure 400 {object} map[string]string "Invalid request parameters"
// @Failure 409 {object} map[string]string "Email already registered"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /register [post]
func (h *userController) RegisterUser(c echo.Context) error {
	var userRequest entity.RegisterRequest

	// Melakukan bind request body ke struct.
	if err := c.Bind(&userRequest); err != nil {
		// Mengembalikan respons jika request body tidak valid.
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	// Memvalidasi data request menggunakan validator.
	err := validate.Struct(userRequest)
	if err != nil {
		// Mengembalikan respons jika validasi gagal.
		return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("invalid request parameters: %v", err)})
	}

	// Memanggil service untuk mendaftarkan user.
	status, webResponse := h.userService.RegisterUser(userRequest)

	// Mengembalikan respons dalam format JSON.
	return c.JSON(status, webResponse)
}

// LoginUser godoc
// @Summary User login
// @Description Authenticates a user by email and password, and returns a JWT token.
// @Tags User
// @Accept json
// @Produce json
// @Param request body entity.LoginRequest true "Login Request"
// @Success 200 {object} map[string]string "User successfully logged in"
// @Failure 400 {object} map[string]string "Invalid request parameters"
// @Failure 404 {object} map[string]string "Invalid email or password"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /login [post]
func (h *userController) LoginUser(c echo.Context) error {
	var loginRequest entity.LoginRequest

	// Melakukan bind request body ke struct.
	if err := c.Bind(&loginRequest); err != nil {
		// Mengembalikan respons jika request body tidak valid.
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request parameters"})
	}

	// Memvalidasi data request menggunakan validator.
	err := validate.Struct(loginRequest)
	if err != nil {
		// Mengembalikan respons jika validasi gagal.
		return c.JSON(http.StatusBadRequest, map[string]string{"message": fmt.Sprintf("invalid request parameters: %v", err)})
	}

	// Memanggil service untuk melakukan login user.
	status, webResponse := h.userService.LoginUser(loginRequest)

	// Mengembalikan respons dalam format JSON.
	return c.JSON(status, webResponse)
}

// UserInfo godoc
// @Summary Get user information
// @Description Retrieves user information such as email, full name, and balance by user ID.
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} entity.UserResponse "User information retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized - User not authenticated"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user-info [get]
func (h *userController) UserInfo(c echo.Context) error {
	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": fmt.Sprintf("User ID is not valid!"),
		})
	}

	status, webResponse := h.userService.UserInfo(userID)

	return c.JSON(status, webResponse)
}

// ConfirmHandler godoc
// @Summary Confirm user account activation
// @Description Handles account activation via a confirmation token passed in the query string. It activates the user's account.
// @Tags User
// @Accept json
// @Produce html
// @Param token query string true "Confirmation Token"
// @Success 200 {object} map[string]string "Account successfully activated"
// @Failure 400 {object} map[string]string "Token is missing"
// @Failure 401 {object} map[string]string "Invalid or expired token"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /confirm [get]
func (h *userController) ConfirmHandler(c echo.Context) error {
	// Extract the token from query parameters
	tokenString := c.QueryParam("token")

	status, webResponse := h.userService.ConfirmHandler(tokenString)

	// Check the status to decide the template to render
	if status == http.StatusOK {
		return c.HTML(http.StatusOK, fmt.Sprintf(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Account Activation</title>
				<style>
					body {
						display: flex;
						justify-content: center;
						align-items: center;
						height: 100vh;
						margin: 0;
						font-family: Arial, sans-serif;
						background-color: #f4f4f4;
					}
					.container {
						text-align: center;
						padding: 20px;
						background: white;
						border-radius: 10px;
						box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
					}
					h1 {
						color: #333;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Account Activation</h1>
					<h1>%s</h1>
				</div>
			</body>
			</html>`, webResponse["message"]))
	}

	// Render an error message for other statuses
	return c.HTML(status, fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Activation Error</title>
			<style>
				body {
					display: flex;
					justify-content: center;
					align-items: center;
					height: 100vh;
					margin: 0;
					font-family: Arial, sans-serif;
					background-color: #f4f4f4;
				}
				.container {
					text-align: center;
					padding: 20px;
					background: white;
					border-radius: 10px;
					box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
				}
				h1 {
					color: #333;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Activation Error</h1>
				<h1>%s</h1>
			</div>
		</body>
		</html>`, webResponse["message"]))
}
