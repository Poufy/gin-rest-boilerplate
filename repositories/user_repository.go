package repositories

import (
	"database/sql"
	"gin-boilerplate/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	// Hash the password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	lastInsertId := 0
	err = ur.db.QueryRow("INSERT INTO users (name, lastname, email, password) VALUES ($1, $2, $3, $4) RETURNING id", user.FirstName, user.LastName, user.Email, user.Password).Scan(&lastInsertId)

	if err != nil {
		return nil, err
	}

	// Get the ID of the newly inserted user
	userID := lastInsertId

	user.ID = int(userID)

	return user, nil
}

func (ur *UserRepository) GetUsers() ([]models.User, error) {
	rows, err := ur.db.Query("SELECT id, name, lastname, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) GetUser(userID string) (*models.User, error) {
	var user models.User
	err := ur.db.QueryRow("SELECT id, name, lastname, email FROM users WHERE id = $1", userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := ur.db.QueryRow("SELECT id, name, lastname, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
