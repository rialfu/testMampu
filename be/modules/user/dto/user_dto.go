package dto

import (
	"errors"
)

const (
	// Failed
	MESSAGE_FAILED_REGISTER_USER       = "failed create user"
	MESSAGE_FAILED_TOKEN_NOT_VALID     = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND     = "token not found"
	MESSAGE_FAILED_GET_USER            = "failed get user"
	MESSAGE_FAILED_DENIED_ACCESS       = "denied access"
	MESSAGE_FAILED_VERIFY_EMAIL        = "failed verify email"
	MESSAGE_FAILED_EMAIL_ALREADY_EXIST = "email already exists"
	MESSAGE_FAILED_PHONE_ALREADY_EXIST = "telp already exists"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"
)

var (
	ErrGetUserById        = errors.New("failed to get user by id")
	ErrGetUserByEmail     = errors.New("failed to get user by email")
	ErrEmailAlreadyExists = errors.New("email already exist")
	ErrPhoneAlreadyExists = errors.New("phone already exist")
	ErrTokenInvalid       = errors.New("token invalid")
	ErrTokenExpired       = errors.New("token expired")
)

type (
	UserCreateRequest struct {
		Name       string `json:"name" form:"name" validate:"required,min=2,max=100"`
		TelpNumber string `json:"telp" form:"telp" validate:"required,min=8,max=14"`
		Email      string `json:"email" form:"email" validate:"required,email"`
		Password   string `json:"password" form:"password" validate:"required,min=8"`
	}

	UserResponse struct {
		ID         uint64 `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		TelpNumber string `json:"telp"`
		// IsVerified bool   `json:"is_verified"`
	}
	UserUpdateRequest struct {
		Name       string `json:"name" form:"name" binding:"omitempty,min=2,max=100"`
		TelpNumber string `json:"telp" form:"telp" binding:"omitempty,min=8,max=20"`
		Email      string `json:"email" form:"email" binding:"omitempty,email"`
	}

	UserUpdateResponse struct {
		ID         uint64 `json:"id"`
		Name       string `json:"name"`
		TelpNumber string `json:"telp"`
		Role       string `json:"role"`
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}

	UserLoginRequest struct {
		EmailOrPhone string `json:"email_phone" form:"email_phone" validate:"required"`
		Password     string `json:"password" form:"password" validate:"required"`
	}
)
