package service

import(
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	// middleware "P2-Hacktiv8/internal/middleware"
	// "golang.org/x/crypto/bcrypt"
	"net/http"
	// "time"
	// "github.com/golang-jwt/jwt/v4"
	// // "gorm.io/gorm"
	"fmt"
	// "P2-Hacktiv8/utils"
)

type SaldoService interface{
	TopUp(topUpRequest entity.TopUpRequest) (int, map[string]interface{})
}

type saldoService struct{
	userRepository repository.UserRepository
}

func NewSaldoService(saldoRepository repository.UserRepository) *saldoService{
	return &saldoService{saldoRepository}
}

func (s *saldoService) TopUp(topUpRequest entity.TopUpRequest) (int, map[string]interface{}){
	user, err := s.userRepository.UpdateBalance(topUpRequest)
	if err != nil{
		return http.StatusInternalServerError, map[string]interface{}{
			"status" : http.StatusInternalServerError,
			"message": fmt.Sprintf("Top up fail in database: %v",err),
		}
	}

	return http.StatusOK, map[string]interface{}{
		"status" : http.StatusOK,
		"message": "Successfully top up balance",
		"data": user,
	}
}