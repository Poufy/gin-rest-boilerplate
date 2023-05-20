package repositories

import (
	"database/sql"

	"gin-boilerplate/models"
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
	result, err := ur.db.Exec("INSERT INTO users (name) VALUES ($1)", user.FirstName)

	if err != nil {
		return nil, err
	}
	// Get the ID of the newly inserted user
	userID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(userID)

	return user, nil
}

func (ur *UserRepository) GetUsers() ([]models.User, error) {
	rows, err := ur.db.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) GetUser(userID string) (*models.User, error) {
	var user models.User
	err := ur.db.QueryRow("SELECT id, name FROM users WHERE id = $1", userID).Scan(&user.ID, &user.FirstName)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
