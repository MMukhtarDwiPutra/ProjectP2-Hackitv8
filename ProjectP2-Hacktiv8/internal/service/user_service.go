package service

import(
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	middleware "P2-Hacktiv8/internal/middleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"fmt"
)

type UserService interface{
	RegisterUser(request entity.RegisterRequest) (int, map[string]interface{})
	LoginUser(request entity.LoginRequest) (int, map[string]interface{})
}

type userService struct{
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService{
	return &userService{userRepository}
}

func (s *userService) RegisterUser(request entity.RegisterRequest) (int, map[string]interface{}){
	findUser, err := s.userRepository.GetUserByEmail(request.Email)
	if findUser != nil {
		return http.StatusConflict, map[string]interface{}{"message": "Tidak berhasil! Email sudah terdaftar!"}
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return http.StatusInternalServerError, map[string]interface{}{
			"message": "Error querying the database.",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		}
	}

	user := entity.User{
		Email: request.Email,
		FullName: request.FullName,
		Username: request.Username,
		Password: string(hashedPassword),
		Age: request.Age,
	}

	userResult, err := s.userRepository.CreateUser(user);
	if err != nil{
		fmt.Println(err)
		return http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		}
	}

	userResponse := entity.RegisterResponse{
		ID: userResult.ID,
		Email: request.Email,
		FullName: request.FullName,
		Username: request.Username,
		Age: request.Age,
	}

	return http.StatusCreated, map[string]interface{}{
		"data": userResponse,
	}
}

func (s *userService) LoginUser(request entity.LoginRequest) (int, map[string]interface{}){
	// Mengecek apakah email siswa ada di database
	user, err := s.userRepository.GetUserByEmail(request.Email)
	if user == nil {
		return http.StatusNotFound, map[string]interface{}{"message": "Invalid Email or Password"}
	}
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{"message": "error query."}
	}

	// Memverifikasi password siswa
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return http.StatusNotFound, map[string]interface{}{"message": "Invalid Email or Password"}
	}

	// Membuat token JWT untuk siswa yang berhasil login
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})

	// Menandatangani token dengan secret key
	tokenString, err := token.SignedString([]byte(middleware.SecretKey))
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"}
	}

	return http.StatusOK, map[string]interface{}{
		"token": tokenString,
	}
}