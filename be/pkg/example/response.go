package example

import (
	authDTO "rialfu/wallet/modules/auth/dto"
	userDTO "rialfu/wallet/modules/user/dto"
	walletDTO "rialfu/wallet/modules/wallet/dto"
)

type ResponseRegister struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Data    userDTO.UserResponse `json:"data,omitempty"`
}

type ResponseLogin struct {
	Status  bool                  `json:"status"`
	Message string                `json:"message"`
	Data    authDTO.TokenResponse `json:"data,omitempty"`
}
type ResponseCheckBalance struct {
	Status  bool                      `json:"status"`
	Message string                    `json:"message"`
	Data    walletDTO.BalanceResponse `json:"data,omitempty"`
}
type ResponseWithdraw struct {
	Status  bool                       `json:"status"`
	Message string                     `json:"message"`
	Data    walletDTO.WithdrawResponse `json:"data,omitempty"`
}
type ResponseStoreWallet struct {
	Status  bool                    `json:"status"`
	Message string                  `json:"message"`
	Data    walletDTO.StoreResponse `json:"data,omitempty"`
}
