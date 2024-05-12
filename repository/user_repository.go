package repository

import (
	"todo-deck-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{ db }
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	err := ur.db.Where("email=?", email).First(user).Error;
	if err != nil {
		return err
	}

	return nil
}