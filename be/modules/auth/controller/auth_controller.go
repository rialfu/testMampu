package controller

import (
	"errors"
	"net/http"
	"rialfu/wallet/modules/auth/dto"
	"rialfu/wallet/modules/auth/service"
	userDto "rialfu/wallet/modules/user/dto"
	"rialfu/wallet/pkg/constants"
	_ "rialfu/wallet/pkg/example"
	"rialfu/wallet/pkg/helpers"
	"rialfu/wallet/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type (
	AuthController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
	}

	authController struct {
		authService service.AuthService
		db          *gorm.DB
	}
)

func NewAuthController(injector *do.Injector, as service.AuthService) AuthController {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)

	return &authController{
		authService: as,
		db:          db,
	}
}

// Register godoc
// @Summary Untuk Register
// @Accept json
// @Produce json
// @Param request body userDto.UserCreateRequest true "Create user payload"
// @Success 200 {object} example.ResponseRegister
// @Router /api/auth/register [post]
func (c *authController) Register(ctx *gin.Context) {
	var req userDto.UserCreateRequest
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
	//

	result, err := c.authService.Register(ctx.Request.Context(), req)
	if err != nil {
		var res utils.Response
		if errors.Is(err, userDto.ErrEmailAlreadyExists) {
			res = utils.BuildResponseFailedValidation(map[string]string{"email": userDto.MESSAGE_FAILED_EMAIL_ALREADY_EXIST}, nil)
		} else if errors.Is(err, userDto.ErrPhoneAlreadyExists) {
			res = utils.BuildResponseFailedValidation(map[string]string{"email": userDto.MESSAGE_FAILED_PHONE_ALREADY_EXIST}, nil)
		} else {
			res = utils.BuildResponseFailed(userDto.MESSAGE_FAILED_REGISTER_USER, err.Error(), nil)
		}

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(userDto.MESSAGE_SUCCESS_REGISTER_USER, result)
	ctx.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary Untuk Login
// @Accept json
// @Produce json
// @Param request body userDto.UserLoginRequest true "Login payload"
// @Success 200 {object} example.ResponseLogin
// @Router /api/auth/login [post]
func (c *authController) Login(ctx *gin.Context) {
	var req userDto.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response := utils.BuildResponseFailed(constants.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var validationMessages = map[string]map[string]string{}
		errs := helpers.TranslateValidationError(err, validationMessages)
		res := utils.BuildResponseFailedValidation(errs, nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		var res utils.Response
		if errors.Is(err, userDto.ErrEmailAlreadyExists) {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, dto.MESSAGE_FAILED_ACCOUNT_OR_PASSWRONG, nil)
		} else {
			res = utils.BuildResponseFailed(dto.MESSAGE_FAILED_LOGIN, err.Error(), nil)
		}

		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(userDto.MESSAGE_SUCCESS_LOGIN, result)
	ctx.JSON(http.StatusOK, res)
}
