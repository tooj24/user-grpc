package v1

import "grpc-clean/internal/models"

type UserRepository interface {
	Create(user models.User) (models.User, error)
	Get(id string) (models.User, error)
	Update(user models.User) error
	Delete(id string) error
	GetByEmail(email string) (models.User, error)
}

type UserUseCase interface {
	Create(user models.User) (models.User, error)
	Get(id string) (models.User, error)
	Update(user models.User) error
	Delete(id string) error
}
