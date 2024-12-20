package service

import(
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	"P2-Hacktiv8/utils"
	"net/http"
	"fmt"
	"gorm.io/gorm"
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

func (s *saldoService) TopUp(topUpRequest entity.BalanceRequest) (int, map[string]interface{}) {
	// Get the user from the repository
	findUser, err := s.userRepository.GetUserById(topUpRequest.UserID)
	if err != nil {
		// Handle user not found error
		if err == gorm.ErrRecordNotFound {
			return http.StatusNotFound, map[string]interface{}{
				"status":  http.StatusNotFound,
				"message": "User not found",
			}
		}
		// Handle internal server error for other errors
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Top up fail in database: %v", err),
		}
	}

	// Add the top-up balance to the user's balance
	topUpRequest.Balance += findUser.Balance

	invoice, err := utils.CreateInvoice(*findUser, topUpRequest)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Top up fail in database: %v", err),
		}
	}

	// Update the balance in the repository
	_, err = s.userRepository.UpdateBalance(topUpRequest)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Top up fail in database: %v", err),
		}
	}

	// Return success response with the updated user data
	return http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Successfully top up balance",
		"data":    invoice,
	}
}
