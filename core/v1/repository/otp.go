package repository

import v1 "github.com/klovercloud-ci-cd/security/core/v1"

// Otp Repository operations otp.
type Otp interface {
	Store(otp v1.Otp) error
	FindByOtp(otp string) v1.Otp
}
