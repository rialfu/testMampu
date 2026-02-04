package controller

import (
	"errors"
	"net/http"
	"rialfu/wallet/modules/wallet/dto"
	"rialfu/wallet/modules/wallet/service"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"
	"rialfu/wallet/pkg/utils"

	_ "rialfu/wallet/pkg/example"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	WalletController interface {
		CheckBalancee(ctx *gin.Context)
		WithdrawProcess(ctx *gin.Context)
		StoreProcess(ctx *gin.Context)
	}

	walletController struct {
		walletService service.WalletService
		db            *gorm.DB
	}
)

func NewUserController(injector *do.Injector, ws service.WalletService) WalletController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)

	return &walletController{
		walletService: ws,

		db: db,
	}
}

// Inquiry Balance godoc
// @Summary Untuk Mengecek sisa saldo
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} example.ResponseCheckBalance
// @Router /api/wallet/balance [get]
func (c *walletController) CheckBalancee(ctx *gin.Context) {
	val, exist := ctx.Get("user_id")
	var userId string
	var res utils.Response

	if exist == false {
		res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return

	}
	userId = val.(string)
	data, err := c.walletService.CheckBalance(ctx.Request.Context(), userId)
	if err != nil {
		var status int
		if errors.Is(err, constants.ErrDataNotFound) {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
			status = http.StatusNotFound
		} else {
			res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, err.Error(), nil)
			status = http.StatusInternalServerError
		}

		ctx.AbortWithStatusJSON(status, res)
		return
	}
	res = utils.BuildResponseSuccess(constants.MESSAGE_SUCCESS_GET_DATA, data)
	ctx.JSON(http.StatusOK, res)
}

// Withdraw Wallet godoc
// @Summary Untuk menarik saldo
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.WithdrawRequest true "Withdraw payload"
// @Success 200 {object} example.ResponseWithdraw
// @Router /api/wallet/withdraw [post]
func (c *walletController) WithdrawProcess(ctx *gin.Context) {
	val, exist := ctx.Get("user_id")
	var userId string
	var res utils.Response

	if exist == false {
		res = utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return

	}
	userId = val.(string)
	var req dto.WithdrawRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Validate request

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var validationMessages = map[string]map[string]string{}
		errs := helpers.TranslateValidationError(err, validationMessages)
		res := utils.BuildResponseFailedValidation(errs, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, err := c.walletService.WithdrawProcess(ctx.Request.Context(), userId, req)
	if err != nil {
		var status int
		if errors.Is(err, constants.ErrDataNotFound) {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_WITHDRAW, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
			status = http.StatusNotFound
		} else if errors.Is(err, dto.ErrTragetBankNotFound) {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_WITHDRAW, dto.MESSAGE_FAILED_TARGET_BANK_NOT_FOUND, nil)
			status = http.StatusBadRequest
		} else if errors.Is(err, dto.ErrInsufficientBalance) {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_WITHDRAW, dto.MESSAGE_FAILED_BALANCE_NOT_ENOUGH, nil)
			status = http.StatusBadRequest
		} else {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_WITHDRAW, err.Error(), nil)
			status = http.StatusInternalServerError
		}

		ctx.AbortWithStatusJSON(status, res)
		return
	}
	res = utils.BuildResponseSuccess(dto.MESSSAGE_SUCCESS_WITHDRAW, data)
	ctx.JSON(http.StatusOK, res)
}

// Store Wallet godoc
// @Summary Untuk menyetor saldo
// @description API for terima request dari sistem lain untuk store wallet
// @Accept json
// @Produce json
// @Param request body dto.StoreRequest true "Store payload"
// @Success 200 {object} example.ResponseStoreWallet
// @Router /api/wallet/deposit [post]
func (c *walletController) StoreProcess(ctx *gin.Context) {
	var res utils.Response
	var req dto.StoreRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Validate request

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var validationMessages = map[string]map[string]string{}
		errs := helpers.TranslateValidationError(err, validationMessages)
		res := utils.BuildResponseFailedValidation(errs, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data, err := c.walletService.StoreBalance(ctx.Request.Context(), req)
	if err != nil {
		var status int
		if errors.Is(err, constants.ErrDataNotFound) {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_WITHDRAW, constants.MESSAGE_FAILED_DATA_NOT_FOUND, nil)
			status = http.StatusNotFound
		} else if errors.Is(err, dto.ErrTragetBankNotFound) {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_WITHDRAW, dto.MESSAGE_FAILED_TARGET_BANK_NOT_FOUND, nil)
			status = http.StatusBadRequest
		} else if errors.Is(err, dto.ErrInsufficientBalance) {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_WITHDRAW, dto.MESSAGE_FAILED_BALANCE_NOT_ENOUGH, nil)
			status = http.StatusBadRequest
		} else {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_WITHDRAW, err.Error(), nil)
			status = http.StatusInternalServerError
		}

		ctx.AbortWithStatusJSON(status, res)
		return
	}
	res = utils.BuildResponseSuccess(dto.MESSSAGE_SUCCESS_WITHDRAW, data)
	ctx.JSON(http.StatusOK, res)
}
