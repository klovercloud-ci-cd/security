package logic

import (
	"github.com/klovercloud-ci-cd/security/config"
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/service"
	"log"
	"net/smtp"
)

type emailService struct {
}

func (e emailService) Listen(otp v1.Otp) {
	var message string
	if otp.BaseUrl == "" {
		message = `Hi ` + otp.Email + `,` + `
		  Please find your OTP attached below. It will be expired within 5 minutes.
		  OTP:` + otp.Otp
	} else {
		message = `Click on the following link to reset your password:\n` +
			otp.BaseUrl + "?email=" + otp.Email + "&otp=" + otp.Otp
	}
	// Create authentication
	auth := smtp.PlainAuth("", config.MailServerHostEmail, config.MailServerHostEmailSecret, config.SmtpHost)
	// Send actual message
	err := smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, config.MailServerHostEmail, []string{otp.Email}, []byte(message))
	if err != nil {
		log.Println(err.Error())
	}
}

// NewEmailService returns service.Media type service
func NewEmailService() service.Media {
	return &emailService{}
}
