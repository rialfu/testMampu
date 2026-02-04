package service

import (
	"context"

	"rialfu/wallet/modules/user/dto"
	"rialfu/wallet/modules/user/repository"
	"rialfu/wallet/pkg/constants"
	"rialfu/wallet/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserById(ctx context.Context, userId string) (dto.UserResponse, error)
	Update(ctx context.Context, req dto.UserUpdateRequest, userId string) (dto.UserUpdateResponse, error)
	GetAllUser(ctx *gin.Context) (helpers.PaginateData[dto.UserResponse], error)
}

type userService struct {
	userRepository repository.UserRepository
	db             *gorm.DB
}

func NewUserService(
	userRepo repository.UserRepository,
	db *gorm.DB,
) UserService {
	return &userService{
		userRepository: userRepo,
		db:             db,
	}
}

func (s *userService) GetUserById(ctx context.Context, userId string) (dto.UserResponse, error) {
	user, isExist, err := s.userRepository.GetUserById(ctx, s.db, userId)
	if err != nil {
		return dto.UserResponse{}, err
	}
	if isExist == false {
		return dto.UserResponse{}, constants.ErrDataNotFound
	}

	return dto.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		TelpNumber: user.TelpNumber,
		// IsVerified: user.IsVerified,
	}, nil
}

func (s *userService) GetAllUser(ctx *gin.Context) (helpers.PaginateData[dto.UserResponse], error) {
	datas, page, limit, total, err := s.userRepository.ReadAll(ctx, s.db, ctx.Request.URL.Query())
	results := make([]dto.UserResponse, 0, len(datas))
	if err != nil {
		return helpers.PaginateData[dto.UserResponse]{}, err
	}
	for _, u := range datas {
		// Mapping manual per item
		item := dto.UserResponse{
			ID:    u.ID,
			Email: u.Email,
			Name:  u.Name,
		}
		results = append(results, item)
	}

	return helpers.PaginateData[dto.UserResponse]{Data: results, Limit: limit, Page: page, Total: total}, nil
}

func (s *userService) Update(ctx context.Context, req dto.UserUpdateRequest, userId string) (dto.UserUpdateResponse, error) {
	user, isExist, err := s.userRepository.GetUserById(ctx, s.db, userId)
	if err != nil {
		return dto.UserUpdateResponse{}, err
	}
	if isExist == false {
		return dto.UserUpdateResponse{}, constants.ErrDataNotFound
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.TelpNumber != "" {
		user.TelpNumber = req.TelpNumber
	}

	updatedUser, err := s.userRepository.Update(ctx, s.db, user)
	if err != nil {
		return dto.UserUpdateResponse{}, err
	}

	return dto.UserUpdateResponse{
		ID:         updatedUser.ID,
		Name:       updatedUser.Name,
		TelpNumber: updatedUser.TelpNumber,
		// Role:       updatedUser.Role,
		Email: updatedUser.Email,
	}, nil
}
