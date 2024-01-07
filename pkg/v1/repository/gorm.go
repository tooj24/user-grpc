package repository

import (
	"gorm.io/gorm"
	"grpc-clean/internal/models"
	interfaces "grpc-clean/pkg/v1"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (repository UserRepositoryImpl) Create(user models.User) (models.User, error) {
	err := repository.db.Create(&user).Error
	return user, err
}

func (repository UserRepositoryImpl) Get(id string) (models.User, error) {
	var user models.User
	err := repository.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (repository UserRepositoryImpl) Update(user models.User) error {
	if err := repository.db.Where("id = ?", user.ID).First(&user).Error; err != nil {
		return err
	}
	err := repository.db.Model(&user).Update("name", user.Name).Error
	return err
}

func (repository UserRepositoryImpl) Delete(id string) error {
	err := repository.db.Where("id = ?", id).Delete(&models.User{}).Error
	return err
}

func (repository UserRepositoryImpl) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := repository.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func NewUserRepositoryImpl(db *gorm.DB) interfaces.UserRepository {
	return &UserRepositoryImpl{db: db}
}
