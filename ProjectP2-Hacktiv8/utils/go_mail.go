package utils

import(
	"gopkg.in/gomail.v2"
)

func SendEmailNotification(from, to, subject, content string){
	m := gomail.NewMessage()
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", "recipient@example.com")
	m.SetHeader("Subject", "Test Email Notification")
	m.SetBody("text/plain", "This is a test email sent using Gomail.")

	// Attach a file (optional)
	// m.Attach("/path/to/file")

	// Set up the SMTP server
	d := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-password")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Could not send email: %v", err)
	}

	log.Println("Email sent successfully!")
}