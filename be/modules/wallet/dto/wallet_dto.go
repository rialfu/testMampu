package dto

import (
	"errors"

	"github.com/shopspring/decimal"
)

const (
	MESSAGE_FAILED_TARGET_BANK_NOT_FOUND = "target bank not found"
	MESSAGE_FAILED_BALANCE_NOT_ENOUGH    = "balance not enough"
	MESSAGE_FAILED_WITHDRAW              = "failed withdraw"
	MESSAGE_FAILED_STORE                 = "failed store"
	MESSSAGE_SUCCESS_WITHDRAW            = "success withdraw"
	MESSSAGE_SUCCESS_STORE               = "success withdraw"
)

var (
	ErrTragetBankNotFound  = errors.New("target bank not found")
	ErrInsufficientBalance = errors.New("balance not enough")
)

type (
	StoreRequest struct {
		Amount       float64 `json:"amount" validate:"required,min=1"`
		PaymentType  string  `json:"payment_type" validate:"required"`
		PaidAt       string  `json:"paid_at" validate:"required,datetime=2006-01-02T15:04:05Z"`
		SignatureKey string  `json:"signature_key" validate:"required"`
		WalletID     string  `json:"wallet_id" validate:"required"`
	}
	BalanceResponse struct {
		ID         string  `json:"wallet_id"`
		Balance    float64 `json:"balance"`
		LastUpdate string  `json:"last_update"`
	}
	WithdrawRequest struct {
		Balance       float64 `json:"balance" validate:"required,min=1"`
		TargetBank    uint    `json:"target_bank" validate:"required"`
		TargetAccount string  `json:"target_account" validate:"required,max=50"`
	}
	BankCodeResponse struct {
		ID       uint   `json:"id" `
		BankCode string `json:"bank_code"`
		BankName string `json:"bank_name"`
	}
	StoreResponse struct {
		ReferenceNo string          `json:"reference_no"`
		Fee         decimal.Decimal `json:"fee"`
		Amount      decimal.Decimal `json:"amount"`
	}
	WithdrawResponse struct {
		ReferenceNo   string          `json:"reference_no"`
		Fee           decimal.Decimal `json:"fee"`
		Amount        decimal.Decimal `json:"amount"`
		TargetBank    string          `json:"target_bank" `
		TargetAccount string          `json:"target_account" `
	}
)
