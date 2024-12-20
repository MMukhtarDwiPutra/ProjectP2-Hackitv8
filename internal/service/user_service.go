package service

import(
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"fmt"
	"P2-Hacktiv8/utils"
	"os"
)

type UserService interface{
	RegisterUser(request entity.RegisterRequest) (int, map[string]interface{})
	LoginUser(request entity.LoginRequest) (int, map[string]interface{})
	UserInfo(id int) (int, map[string]interface{})
	ConfirmHandler(tokenString string) (int, map[string]interface{})
}

type userService struct{
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService{
	return &userService{userRepository}
}

func (s *userService) RegisterUser(request entity.RegisterRequest) (int, map[string]interface{}){
	secretKey := os.Getenv("REGIST_SECRET_KEY")
	appUrl := os.Getenv("APP_URL")

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

	// Membuat token JWT untuk siswa yang berhasil login
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": request.Email,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})
	// Menandatangani token dengan secret key
	tokenString, err := token.SignedString([]byte(secretKey))

	user := entity.User{
		Email: request.Email,
		FullName: request.FullName,
		Password: string(hashedPassword),
		Balance: 0,
		JwtToken: tokenString,
		IsActivated: "NOT YET",
	}

	userResult, err := s.userRepository.CreateUser(user);
	if err != nil{
		fmt.Println(err)
		return http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		}
	}

	userResponse := entity.UserResponse{
		UserID: userResult.UserID,
		Email: request.Email,
		FullName: request.FullName,
		Balance: 0,
	}

	confirmationLink := fmt.Sprintf("%v/confirm?token=%s", appUrl, tokenString)

	to := request.Email
	subject := "Register P2-Hacktiv8 Successfully"
	content := fmt.Sprintf("Your register in our website is success! Please use this link %v to activate your account!", confirmationLink)
	utils.SendEmailNotification(to, subject, content)

	return http.StatusCreated, map[string]interface{}{
		"status": http.StatusCreated,
		"message": "User response created successfully",
		"data": userResponse,
	}
}

func (s *userService) LoginUser(request entity.LoginRequest) (int, map[string]interface{}){
	secretKey := os.Getenv("LOGIN_SECRET_KEY")

	// Mengecek apakah email siswa ada di database
	user, err := s.userRepository.GetUserByEmail(request.Email)
	if user == nil {
		return http.StatusNotFound, map[string]interface{}{"message": "Invalid Email or Password"}
	}
	if err != nil {
		if err != gorm.ErrRecordNotFound{
			return http.StatusNotFound, map[string]interface{}{"message": "Invalid Email or Password"}
		}

		return http.StatusInternalServerError, map[string]interface{}{
			"message": "Error querying the database.",
		}
	}

	// Memverifikasi password siswa
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return http.StatusNotFound, map[string]interface{}{
			"status": http.StatusNotFound,
			"message": "Invalid Email or Password",
		}
	}

	// Membuat token JWT untuk siswa yang berhasil login
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})

	// Menandatangani token dengan secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status": http.StatusInternalServerError,
			"message": "error signed jwt token",
		}
	}

	return http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "Login successfuly",
		"token": tokenString,
	}
}

func (s *userService) UserInfo(id int) (int, map[string]interface{}){
	user, err := s.userRepository.GetUserById(id)
	if user == nil {
		return http.StatusNotFound, map[string]interface{}{"message": "User not found"}
	}
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status": http.StatusInternalServerError,
			"message": "error query get user by email",
		}
	}

	userResponse := entity.UserResponse{
		UserID: user.UserID,
		Email: user.Email,
		FullName: user.FullName,
		Balance: user.Balance,
	}

	return http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "Successfully getting user info",
		"data": userResponse,
	}
}

func (s *userService) ConfirmHandler(tokenString string) (int, map[string]interface{}){
	if tokenString == "" {
		return http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"message": "Token is missing",
		}
	}

	user, err := utils.ParseRegisTokenString(tokenString)
	if err != nil{
		return http.StatusUnauthorized, map[string]interface{}{
			"status": http.StatusUnauthorized,
			"message": fmt.Sprintf("Invalid token: %v", err),
		}
	}

	findUser, err := s.userRepository.GetUserByEmail(user["email"])
	if findUser != nil{
		if(findUser.IsActivated == "Activated"){
			return http.StatusOK, map[string]interface{}{
				"status": http.StatusOK,
				"message": "Your account has been activated before! You can use your account!",
			}
		}

		s.userRepository.UpdateIsActivatedById(findUser.UserID, "Activated")
		return http.StatusOK, map[string]interface{}{
			"status": http.StatusOK,
			"message": "Activated account successfuly",
		}
	}

	// Mengembalikan respons Unauthorized jika klaim dalam token tidak valid.
	return http.StatusUnauthorized, map[string]interface{}{
		"status": http.StatusUnauthorized,
		"message": fmt.Sprintf("Invalid token claims"),
	}
}