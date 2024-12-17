package routes

import(
	"P2-Hacktiv8/internal/controller"
	// internal "P2-Hacktiv8/internal/middleware"
	"P2-Hacktiv8/internal/service"
	"P2-Hacktiv8/repository"
	// _ "P2-Hacktiv8/docs" // Import the generated Swagger docs

	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// NewRouter membuat dan mengonfigurasi router Echo untuk aplikasi ini, termasuk semua route dan controller yang diperlukan.
func NewRouter(db *gorm.DB) *echo.Echo {
	// Inisialisasi repository
	// cfg := LoadConfig()

	userRepository := repository.NewUserRepository(db)

	// Inisialisasi service
	userService := service.NewUserService(userRepository)

	// Inisialisasi controller
	userController := controller.NewUserController(userService)

	// Membuat instance baru dari Echo
	e := echo.New()

	// Middleware global untuk logging, recovery, dan autentikasi
	e.Use(middleware.Logger())   // Log setiap request yang diterima
	e.Use(middleware.Recover()) // Recover dari panic atau error di server
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Route user
	e.POST("/users/register", userController.RegisterUser)
	e.POST("/users/login", userController.LoginUser)


	// Mengembalikan instance Echo yang sudah dikonfigurasi
	return e
}
