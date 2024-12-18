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
// @Description Registers a user by providing their email, username, full name, age, and password.
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
