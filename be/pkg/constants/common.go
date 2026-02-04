package constants

import "errors"

const (
	ENUM_ROLE_ADMIN = "admin"
	ENUM_ROLE_USER  = "user"

	ENUM_RUN_PRODUCTION = "production"
	ENUM_RUN_TESTING    = "testing"

	ENUM_PAGINATION_PER_PAGE = 10
	ENUM_PAGINATION_PAGE     = 1

	DB         = "db"
	JWTService = "JWTService"

	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_CREATE_DATA        = "failed create data"
	MESSAGE_FAILED_GET_LIST_DATA      = "failed get list data"
	MESSAGE_FAILED_TOKEN_NOT_VALID    = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND    = "token not found"
	MESSAGE_FAILED_GET_DATA           = "failed get data"
	MESSAGE_FAILED_LOGIN              = "failed login"
	MESSAGE_FAILED_AUTH_WRONG         = "username or password is wrong"
	MESSAGE_FAILED_UPDATE_DATA        = "failed update data"
	MESSAGE_FAILED_DELETE_DATA        = "failed delete data"
	MESSAGE_FAILED_PROSES_REQUEST     = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS      = "denied access"
	MESSAGE_FAILED_AUTH_NOT_FOUND     = "auth not found"
	MESSAGE_FAILED_DATA_NOT_FOUND     = "data not found"

	// Success
	MESSAGE_SUCCESS_CREATE_DATA   = "success create data"
	MESSAGE_SUCCESS_GET_LIST_DATA = "success get list data"
	MESSAGE_SUCCESS_GET_DATA      = "success get data"
	MESSAGE_SUCCESS_UPDATE_DATA   = "success update data"
	MESSAGE_SUCCESS_DELETE_DATA   = "success delete data"
	Charset                       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	ErrForbiddenUpdate = errors.New("failed to update data")
	ErrDataNotFound    = errors.New("data not found")
	ErrProcessData     = errors.New("error process data")
)
