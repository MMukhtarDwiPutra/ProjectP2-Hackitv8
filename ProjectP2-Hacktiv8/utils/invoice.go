package utils

import (
	"bytes"
	"encoding/json"
	"P2-Hacktiv8/entity"
	"net/http"
	"os"
	"encoding/base64"
	"fmt"
)



func CreateInvoice(user entity.User, topUpRequest entity.BalanceRequest) (*entity.ChargeResponse, error) {
	apiKey := os.Getenv("3RD_PARTY_XENDIT_API")+":"
	apiUrl := "https://api.xendit.co/ewallets/charges"

	bodyRequest := map[string]interface{}{
		"currency": "IDR",
		"amount":   topUpRequest.Balance,
		"payment_method": map[string]interface{}{
			"type":        "EWALLET",
			"reusability": "ONE_TIME_USE",
			"ewallet": map[string]interface{}{
				"channel_code": "DANA",
				"channel_properties": map[string]interface{}{
					"success_return_url": "https://your-redirect-website.com/success",
				},
			},
		},
		"customer_id": "cust-acbd66b4-c7c1-4a92-ab5d-6c5fe93b7917",
		"metadata": map[string]interface{}{
			"sku": "ABCDEFGH",
		},
	}

	reqBody, err := json.Marshal(bodyRequest)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	apiKey = "Basic "+base64.StdEncoding.EncodeToString([]byte(apiKey+":"))
	request.SetBasicAuth(apiKey, "")
	request.Header.Set("callback_url", "https://your-callback-url.com/endpoint")
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	fmt.Println(response.Body)
	var paymentResponse entity.ChargeResponse
	if err := json.NewDecoder(response.Body).Decode(&paymentResponse); err != nil {
		return nil, err
	}

	return &paymentResponse, nil
}

// {"id":"65a738de263bd04f8f4b7194","external_id":"payment-link-example","user_id":"57fdbb445eec38910d3a4c47","status":"PENDING","merchant_name":"Your Company","merchant_profile_picture_url":"https://xnd-companies.s3.amazonaws.com/prod/1476344224287_930.png","amount":100000,"description":"Invoice Demo #123","expiry_date":"2024-01-18T02:18:06.783Z","invoice_url":"https://checkout-staging.xendit.co/latest/65a738de263bd04f8f4b7194","available_banks":[{"bank_code":"MANDIRI","collection_type":"POOL","transfer_amount":100000,"bank_branch":"Virtual Account","account_holder_name":"YOUR COMPANY","identity_amount":0},{"bank_code":"BRI","collection_type":"POOL","transfer_amount":100000,"bank_branch":"Virtual Account","account_holder_name":"YOUR COMPANY","identity_amount":0},{"bank_code":"BNI","collection_type":"POOL","transfer_amount":100000,"bank_branch":"Virtual Account","account_holder_name":"YOUR COMPANY","identity_amount":0},{"bank_code":"PERMATA","collection_type":"POOL","transfer_amount":100000,"bank_branch":"Virtual Account","account_holder_name":"YOUR COMPANY","identity_amount":0}],"available_retail_outlets":[],"available_ewallets":[],"available_qr_codes":[{"qr_code_type":"QRIS"}],"available_direct_debits":[],"available_paylaters":[{"paylater_type":"KREDIVO"},{"paylater_type":"UANGME"},{"paylater_type":"AKULAKU"},{"paylater_type":"ATOME"}],"should_exclude_credit_card":false,"should_send_email":false,"success_redirect_url":"https://www.google.com","failure_redirect_url":"https://www.google.com","created":"2024-01-17T02:18:06.932Z","updated":"2024-01-17T02:18:06.932Z","currency":"IDR","items":[{"name":"Air Conditioner","quantity":1,"price":100000,"category":"Electronic","url":"https://yourcompany.com/example_item"}],"fees":[{"type":"ADMIN","value":5000}],"customer":{"given_names":"John","surname":"Doe","email":"johndoe@example.com","mobile_number":"+6287774441111","addresses":[{"city":"Jakarta Selatan","country":"Indonesia","postal_code":"12345","state":"Daerah Khusus Ibukota Jakarta","street_line1":"Jalan Makan","street_line2":"Kecamatan Kebayoran Baru"}]},"customer_notification_preference":{"invoice_created":["whatsapp","email","viber"],"invoice_reminder":["whatsapp","email","viber"],"invoice_paid":["whatsapp","email","viber"]}}%
