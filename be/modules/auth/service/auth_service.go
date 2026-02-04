package service

import (
	"context"

	"rialfu/wallet/database/entities"
	"rialfu/wallet/modules/auth/dto"
	userDto "rialfu/wallet/modules/user/dto"
	userRepo "rialfu/wallet/modules/user/repository"
	walletRepo "rialfu/wallet/modules/wallet/repository"
	"rialfu/wallet/pkg/helpers"

	"gorm.io/gorm"
)

type AuthService interface {
	Register(ctx context.Context, req userDto.UserCreateRequest) (userDto.UserResponse, error)
	Login(ctx context.Context, req userDto.UserLoginRequest) (dto.TokenResponse, error)
}

type authService struct {
	userRepository   userRepo.UserRepository
	jwtService       JWTService
	db               *gorm.DB
	transactor       helpers.Transactor
	walletRepository walletRepo.WalletRepository
}

func NewAuthService(
	userRepo userRepo.UserRepository,
	walletRepo walletRepo.WalletRepository,
	jwtService JWTService,
	db *gorm.DB,
) AuthService {

	transactor := helpers.NewGormTransactor(db)
	return &authService{
		userRepository:   userRepo,
		jwtService:       jwtService,
		db:               db,
		transactor:       transactor,
		walletRepository: walletRepo,
	}
}

func (s *authService) Register(ctx context.Context, req userDto.UserCreateRequest) (userDto.UserResponse, error) {
	var data userDto.UserResponse
	err := s.transactor.WithTransaction(ctx, func(txCtx context.Context) error {
		_, isExist, err := s.userRepository.CheckEmail(txCtx, s.db, req.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if isExist {
			return userDto.ErrEmailAlreadyExists
		}
		_, isExist, err = s.userRepository.CheckNumber(txCtx, s.db, req.TelpNumber)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if isExist {
			return userDto.ErrPhoneAlreadyExists
		}

		hashedPassword, err := helpers.HashPassword(req.Password)
		if err != nil {
			return err
		}

		user := entities.User{
			Name:       req.Name,
			Email:      req.Email,
			TelpNumber: req.TelpNumber,
			Password:   hashedPassword,
		}

		createdUser, err := s.userRepository.Register(txCtx, s.db, user)
		if err != nil {
			return err
		}
		information := entities.InformationUser{
			UserID: createdUser.ID,
		}
		information, err = s.userRepository.InitialInformation(txCtx, s.db, information)
		if err != nil {
			return err
		}
		walletInfo := entities.Wallet{
			UserID: createdUser.ID,
		}
		walletInfo, err = s.walletRepository.Create(txCtx, s.db, walletInfo)
		if err != nil {
			return err
		}

		data = userDto.UserResponse{
			ID:         createdUser.ID,
			Name:       createdUser.Name,
			Email:      createdUser.Email,
			TelpNumber: createdUser.TelpNumber,
		}

		return nil
	})
	return data, err

}

func (s *authService) Login(ctx context.Context, req userDto.UserLoginRequest) (dto.TokenResponse, error) {
	var user entities.User
	var err error
	var isExist bool
	isEmail := helpers.IsEmailValid(req.EmailOrPhone)
	if isEmail {
		user, isExist, err = s.userRepository.CheckEmail(ctx, s.db, req.EmailOrPhone)
	} else {
		user, isExist, err = s.userRepository.CheckNumber(ctx, s.db, req.EmailOrPhone)
	}

	if err != nil {
		return dto.TokenResponse{}, err
	}
	if !isExist {
		return dto.TokenResponse{}, dto.ErrAccountOrPasswordWrong
	}

	isValid, err := helpers.CheckPassword(user.Password, req.Password)
	if err != nil {
		return dto.TokenResponse{}, err
	}
	if !isValid {
		return dto.TokenResponse{}, dto.ErrAccountOrPasswordWrong
	}
	accessToken := s.jwtService.GenerateAccessToken(user.ID)

	if err != nil {
		return dto.TokenResponse{}, err
	}

	return dto.TokenResponse{
		AccessToken: accessToken,
		Name:        user.Name,
		Email:       user.Email,
		Telp:        user.TelpNumber,
	}, nil
}
