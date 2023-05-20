package services

import (
	"gin-boilerplate/models"
	"gin-boilerplate/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	// Perform any necessary validations or transformations on the user data
	// ...

	// Call the UserRepository to create the user
	createdUser, err := us.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (us *UserService) GetUsers() ([]models.User, error) {
	// Call the UserRepository to fetch users
	users, err := us.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) GetUser(userID string) (*models.User, error) {
	// Call the UserRepository to fetch a specific user
	user, err := us.userRepository.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
