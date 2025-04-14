package repository

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"mystore/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserById(id int64) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUserById(id int64) error
}

type userRepository struct {
	db  *sql.DB
	Log *zap.Logger
}

func NewUserRepository(db *sql.DB, logger *zap.Logger) UserRepository {

	return &userRepository{db: db,
		Log: logger}
}

func (r *userRepository) CreateUser(user *models.User) error {
	err := r.db.QueryRow(`
  INSERT INTO users (name, email, password, role) 
  VALUES ($1, $2, $3, $4) 
  RETURNING id, created_at
`, user.Name, user.Email, user.Password, user.Role).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		r.Log.Error("Failed to insert user", zap.Error(err))
		return errors.New("failed to insert user")
	}
	return nil
}

func (r userRepository) GetUserById(id int64) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email, password, role, created_at FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		r.Log.Error("Failed to get user by ID", zap.Error(err))
		return nil, errors.New("failed to get user")
	}
	return user, nil
}

func (r userRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email, password, role, created_at FROM users WHERE name = $1"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		r.Log.Error("Failed to get user by username", zap.Error(err))
		return nil, errors.New("error: userRepository.GetUserByUsername")
	}
	return user, nil
}

func (r userRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, name, email, password, role, created_at FROM users WHERE email = $1"
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		r.Log.Error("Failed to get user by email", zap.Error(err))
		return nil, errors.New("error: userRepository.GetUserByEmail")
	}
	return user, nil
}

func (r userRepository) UpdateUser(user *models.User) error {
	query := `UPDATE users SET name = $2, email = $3, password = $4, role = $5 WHERE id = $1`

	res, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		r.Log.Error("Failed to update user", zap.Error(err))
		return errors.New("error: userRepository.UpdateUser")
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("error: userRepository.UpdateUser no rows affected")
	}
	return nil
}

func (r userRepository) DeleteUserById(id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		r.Log.Error("Failed to delete user", zap.Error(err))
		return errors.New("error: userRepository.DeleteUserById")
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("error: userRepository.DeleteUserById no rows affected")
	}

	return nil
}
