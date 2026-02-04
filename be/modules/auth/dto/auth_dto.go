package dto

import (
	"errors"
)

const (
	MESSAGE_FAILED_LOGIN                = "failed login"
	MESSAGE_FAILED_REFRESH_TOKEN        = "failed refresh token"
	MESSAGE_SUCCESS_REFRESH_TOKEN       = "success refresh token"
	MESSAGE_FAILED_LOGOUT               = "failed logout"
	MESSAGE_SUCCESS_LOGOUT              = "success logout"
	MESSAGE_FAILED_SEND_PASSWORD_RESET  = "failed send password reset"
	MESSAGE_SUCCESS_SEND_PASSWORD_RESET = "success send password reset"
	MESSAGE_FAILED_RESET_PASSWORD       = "failed reset password"
	MESSAGE_FAILED_ACCOUNT_OR_PASSWRONG = "account or password is wrong"
	MESSAGE_SUCCESS_RESET_PASSWORD      = "success reset password"
)

var (
	ErrRefreshTokenNotFound   = errors.New("refresh token not found")
	ErrRefreshTokenExpired    = errors.New("refresh token expired")
	ErrInvalidCredentials     = errors.New("invalid credentials")
	ErrPasswordResetToken     = errors.New("password reset token invalid")
	ErrAccountOrPasswordWrong = errors.New("account wrong")
)

type (
	TokenResponse struct {
		AccessToken string `json:"access_token"`
		Name        string `json:"name"`
		Telp        string `json:"telp"`
		Email       string `json:"email"`
	}
)
