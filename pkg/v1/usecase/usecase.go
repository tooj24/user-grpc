package usecase

import (
	"errors"
	"gorm.io/gorm"
	"grpc-clean/internal/models"
	interfaces "grpc-clean/pkg/v1"
)

type UserUseCaseImpl struct {
	repository interfaces.UserRepository
}

func (useCase UserUseCaseImpl) Create(user models.User) (models.User, error) {
	if _, err := useCase.repository.GetByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("the email is already associated with another user")
	}
	return useCase.repository.Create(user)
}

func (useCase UserUseCaseImpl) Get(id string) (models.User, error) {
	var user models.User
	var err error

	if user, err = useCase.repository.Get(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("no such user with id supplied")
		}

		return models.User{}, err
	}

	return user, nil
}

func (useCase UserUseCaseImpl) Update(updateUser models.User) error {
	var user models.User
	var err error

	if user, err = useCase.Get(string(updateUser.ID)); err != nil {
		return err
	}

	if user.Email != updateUser.Email {
		return errors.New("email cannot be changed")
	}

	return useCase.repository.Update(updateUser)
}

func (useCase UserUseCaseImpl) Delete(id string) error {
	var err error

	if _, err = useCase.Get(id); err != nil {
		return err
	}

	return useCase.repository.Delete(id)
}

func NewUserUseCaseImpl(repository interfaces.UserRepository) interfaces.UserUseCase {
	return &UserUseCaseImpl{repository: repository}
}
