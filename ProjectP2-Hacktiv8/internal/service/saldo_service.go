package service

import(
	"P2-Hacktiv8/entity"
	"P2-Hacktiv8/repository"
	"P2-Hacktiv8/utils"
	"net/http"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type SaldoService interface{
	TopUp(topUpRequest entity.BalanceRequest) (int, map[string]interface{})
	CallbackWebhook(webhookPayload entity.WebhookPayload) (int)
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

	invoiceID, err := s.userRepository.GetLastIDXendit()
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Top up fail in database: %v", err),
		}
	}

	invoiceString := strconv.Itoa(*invoiceID)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error convert int", err),
		}
	}

	invoiceString = "INV"+invoiceString

	invoice, err := utils.CreateInvoice(*findUser, topUpRequest, invoiceString)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Top up fail in database: %v", err),
		}
	}

	xenditPayment := entity.WebhookXenditPayment{
		InvoiceID:invoiceString,
		UserIDApp:topUpRequest.UserID,
		Status:invoice.Status,
	}

	_, err = s.userRepository.CreateXenditHistory(xenditPayment)
	if err != nil {
		return http.StatusInternalServerError, map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Creating xendit history fail in database: %v", err),
		}
	}

	topUpResponse := entity.TopUpResponse{
		InvoiceID:    invoice.ID,
		Status:       invoice.Status,
		Description:  invoice.Description,
		Url:          invoice.InvoiceURL,
		MerchantName: invoice.MerchantName,
		ExternalID:   invoice.ExternalID,
	}

	// Send email notification
	to := findUser.Email
	subject := fmt.Sprintf("Payment for Top Up %v is Successful", findUser.FullName)

	// Format the content to include detailed payment and booking information
	content := fmt.Sprintf(`
	Dear %s,

	Thank you for making your payment with us! Your booking has been successfully processed. Here are the details:

	- Payment Created Date: %v

	Payment Details:
	- Payment Status: %s
	- Invoice ID: %s
	- Description: %s
	- Payment Link: %s
	- Merchant: %s

	Please save this email for your reference. If you have any questions or require assistance, feel free to reach out to us.

	Best regards,  
	Your Booking Team
	`, 
		findUser.FullName,
		time.Now().Format("January 2, 2006, 3:04 PM"),
		topUpResponse.Status, 
		topUpResponse.InvoiceID, 
		topUpResponse.Description, 
		topUpResponse.Url, 
		invoice.MerchantName)

	// Send the email
	utils.SendEmailNotification(to, subject, content)

	// Return success response with the updated user data
	return http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Successfully top up balance",
		"data":    topUpResponse,
	}
}

func (s *saldoService) CallbackWebhook(webhookPayload entity.WebhookPayload) (int){
	// Process the webhook data based on the invoice status
	if webhookPayload.Status == "PAID" {
		// Logic for updating user balance or marking invoice as paid
		// You can access the webhook data like this:
		// - webhookPayload.Amount
		// - webhookPayload.UserID
		// - webhookPayload.PayerEmail

		// Example: Update balance based on paid amount
		paymentXendit, err := s.userRepository.GetPaymentIdByInvoiceId(webhookPayload.ExternalID)
		if paymentXendit.Status == "PAID"{
			return http.StatusOK
		}
		findUserApp, err := s.userRepository.GetUserById(paymentXendit.UserIDApp)

		// Add the top-up balance to the user's balance
		topUpRequest := entity.BalanceRequest{
			UserID: findUserApp.UserID,
			Balance: findUserApp.Balance + (float32(webhookPayload.Amount)/float32(1000)),
		}

		// Update the balance in the repository
		_, err = s.userRepository.UpdateBalance(topUpRequest)
		if err != nil {
			return http.StatusInternalServerError
		}

		_, err = s.userRepository.UpdateStatusWebhookXenditPayment(*paymentXendit)
		if err != nil {
			return http.StatusInternalServerError
		}

		return http.StatusOK
	} else if webhookPayload.Status == "EXPIRED" {
		// Handle expired invoices if needed
		// Example: Mark the invoice as expired
		// MarkInvoiceAsExpired(webhookPayload.ID)
	}

	return http.StatusNotFound
}