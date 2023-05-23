package services

import (
	"errors"
	"gin-boilerplate/middleware"
	"gin-boilerplate/models"
	"gin-boilerplate/repositories"

	"golang.org/x/crypto/bcrypt"
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

func (us *UserService) Login(email, password string) (string, error) {
	// Fetch the user by email from the UserRepository
	user, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("Invalid credentials")
	}

	// Generate a secure authentication token
	token, err := middleware.GenerateAuthToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
