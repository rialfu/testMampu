package repository

import (
	"context"
	"errors"

	"rialfu/wallet/database/entities"
	"rialfu/wallet/pkg/helpers"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Register(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error)
		GetUserById(ctx context.Context, tx *gorm.DB, userId string) (entities.User, bool, error)
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entities.User, bool, error)
		CheckNumber(ctx context.Context, tx *gorm.DB, email string) (entities.User, bool, error)
		Update(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error)
		InitialInformation(ctx context.Context, tx *gorm.DB, user entities.InformationUser) (entities.InformationUser, error)
		ReadAll(ctx context.Context, tx *gorm.DB, queryParam map[string][]string) ([]entities.User, int, int, int64, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) getDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("DB_TX").(*gorm.DB); ok {
		return tx
	}
	return r.db
}

func (r *userRepository) Register(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, tx *gorm.DB, userId string) (entities.User, bool, error) {
	db := r.getDB(ctx)
	var user entities.User
	if err := db.WithContext(ctx).Where("id = ?", userId).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, false, nil
		}
		return entities.User{}, false, err
	}

	return user, true, nil
}
func (r *userRepository) ReadAll(ctx context.Context, tx *gorm.DB, queryParam map[string][]string) ([]entities.User, int, int, int64, error) {
	var users []entities.User
	db := r.db.Model(&entities.User{})
	db, pagination := helpers.ApplyPagination(db, &entities.User{}, queryParam)
	var total int64
	if err := db.Find(&users).Error; err != nil {
		return []entities.User{}, 0, 0, 0, err
	}
	if err := db.Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return []entities.User{}, 0, 0, 0, err
	}
	return users, pagination.Page, pagination.Limit, total, nil
}

func (r *userRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entities.User, bool, error) {
	db := r.getDB(ctx)

	var user entities.User
	if err := db.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, false, nil
		}
		return entities.User{}, false, err
	}

	return user, true, nil
}

func (r *userRepository) CheckNumber(ctx context.Context, tx *gorm.DB, email string) (entities.User, bool, error) {
	db := r.getDB(ctx)

	var user entities.User
	if err := db.WithContext(ctx).Where("telp_number = ?", email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, false, nil
		}
		return entities.User{}, false, err
	}

	return user, true, nil
}

func (r *userRepository) Update(ctx context.Context, tx *gorm.DB, user entities.User) (entities.User, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Updates(&user).Error; err != nil {
		return entities.User{}, err
	}

	return user, nil
}
func (r *userRepository) InitialInformation(ctx context.Context, tx *gorm.DB, user entities.InformationUser) (entities.InformationUser, error) {
	db := r.getDB(ctx)

	if err := db.WithContext(ctx).Create(&user).Error; err != nil {
		return entities.InformationUser{}, err
	}

	return user, nil
}
