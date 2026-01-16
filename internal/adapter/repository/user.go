package repository

import (
	"context"
	"go-boilerplate/internal/core/domain/entity"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	// Define methods for user repository
	CreateUser(ctx context.Context, req entity.UserEntity) error
	ListUsers(ctx context.Context, limit int, offset int) ([]entity.UserEntity, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
	GetUserByID(ctx context.Context, id int64) (*entity.UserEntity, error)
	UpdateUser(ctx context.Context, req entity.UserEntity) error
	UpdateUserPassword(ctx context.Context, req entity.UserEntity) error
	UpdateUserVerificationStatus(ctx context.Context, userID int64) (*entity.UserEntity, error)
	DeleteUser(ctx context.Context, id int64) error
}

type UserRepository struct {
	// Add necessary fields, e.g., database connection
	db *gorm.DB
}

// CreateUser implements [UserRepositoryInterface].
func (u *UserRepository) CreateUser(ctx context.Context, req entity.UserEntity) error {
	panic("unimplemented")
}

// DeleteUser implements [UserRepositoryInterface].
func (u *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetUserByEmail implements [UserRepositoryInterface].
func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	panic("unimplemented")
}

// GetUserByID implements [UserRepositoryInterface].
func (u *UserRepository) GetUserByID(ctx context.Context, id int64) (*entity.UserEntity, error) {
	panic("unimplemented")
}

// ListUsers implements [UserRepositoryInterface].
func (u *UserRepository) ListUsers(ctx context.Context, limit int, offset int) ([]entity.UserEntity, error) {
	panic("unimplemented")
}

// UpdateUser implements [UserRepositoryInterface].
func (u *UserRepository) UpdateUser(ctx context.Context, req entity.UserEntity) error {
	panic("unimplemented")
}

// UpdateUserPassword implements [UserRepositoryInterface].
func (u *UserRepository) UpdateUserPassword(ctx context.Context, req entity.UserEntity) error {
	panic("unimplemented")
}

// UpdateUserVerificationStatus implements [UserRepositoryInterface].
func (u *UserRepository) UpdateUserVerificationStatus(ctx context.Context, userID int64) (*entity.UserEntity, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		db: db,
	}
}
