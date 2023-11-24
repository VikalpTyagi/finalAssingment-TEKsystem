package pkg

import (
	"finalAssing/internal/config"
	"fmt"
	"net/smtp"

	"github.com/rs/zerolog/log"
)

func EmailSender(userEmail string, otp int) error {
	// Sender's email address and password
	cfg := config.GetConfig()
	// from := "job.Project99@gmail.com"
	// password := "nnyi uvzv epbn jbzv"
	from := cfg.EmailConfig.SenderMail
	password := cfg.EmailConfig.Password

	// Recipient's email address
	to := userEmail

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	// smtpPort := 587
	smtpPort := cfg.EmailConfig.Port

	// Message content
	message := []byte(fmt.Sprintf("Subject: Slavery Password Reset\n\nThis is a test email body.\nWe have recieved your password reset request successfully!\nPlease use the %d OTP, It's valid for only 5 minutes.\n\n\nRegards,\nSlavers Team", otp))

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err := smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if err != nil {
		log.Error().Err(err).Str("User Email", userEmail).Msg("Error in sending email")
		return err
	}

	log.Info().Msg("Email sent successfully!")
	return nil
}
