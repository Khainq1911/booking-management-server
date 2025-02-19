package email

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(receivedMail, message string) error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	auth := smtp.PlainAuth("", os.Getenv("SMTP_EMAIL"), os.Getenv("EMAIL_PASSWORD"), "smtp.gmail.com")

	to := []string{receivedMail}
	msg := []byte("Subject: Bảng lương nhân viên\n\n" + message)

	err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("SMTP_EMAIL"), to, msg)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}
