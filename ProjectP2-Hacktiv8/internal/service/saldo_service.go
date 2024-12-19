package service

import(
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	"net/http"
	"fmt"
)

type SaldoService interface{
	TopUp(topUpRequest entity.BalanceRequest) (int, map[string]interface{})
}

type saldoService struct{
	userRepository repository.UserRepository
}

func NewSaldoService(saldoRepository repository.UserRepository) *saldoService{
	return &saldoService{saldoRepository}
}

func (s *saldoService) TopUp(topUpRequest entity.BalanceRequest) (int, map[string]interface{}){
	findUser, err := s.userRepository.GetUserById(topUpRequest.UserID)
	if err != nil{
		return http.StatusInternalServerError, map[string]interface{}{
			"status" : http.StatusInternalServerError,
			"message": fmt.Sprintf("Top up fail in database: %v",err),
		}
	}

	topUpRequest.Balance += findUser.Balance

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