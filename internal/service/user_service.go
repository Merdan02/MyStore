package service

import (
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"mystore/internal/models"
	"mystore/internal/repository"
	"strings"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]*models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUserById(id int) error
	//Login(name, password, email string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
	Log  *zap.Logger
}

func NewUserService(repo repository.UserRepository, logger *zap.Logger) UserService {
	return &userService{
		repo: repo,
		Log:  logger,
	}
}
func (s *userService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Password == "" || user.Email == "" || user.Role == "" {
		s.Log.Error("user data is empty")
		return errors.New("data is empty")
	}

	existingUser, err := s.repo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		s.Log.Error("user already exists", zap.String("email", user.Email))
		return errors.New("user with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.Error("failed to hash password", zap.Error(err))
		return errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)

	if err := s.repo.CreateUser(user); err != nil {
		s.Log.Error("user creation failed", zap.Error(err))
		return err
	}

	return nil
}

func (s *userService) GetAllUsers() ([]*models.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		s.Log.Error("users retrieval failed", zap.Error(err))
		return nil, err
	}
	return users, nil
}

func (s *userService) GetUserById(id int) (*models.User, error) {

	if id == 0 {
		s.Log.Error("user id is empty")
		return nil, errors.New("user id is empty")
	}

	user, err := s.repo.GetUserById(id)
	if err != nil {
		s.Log.Error("user not found with this id", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByUsername(username string) (*models.User, error) {
	if username == "" {
		s.Log.Error("username is empty")
		return nil, errors.New("username is empty")
	}
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		s.Log.Error("user not found with this username", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	if email == "" {
		s.Log.Error("email is empty")
		return nil, errors.New("email is empty")
	}
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		s.Log.Error("user not found with this email", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(user *models.User) error {

	if user.Name == "" || user.Password == "" || user.Email == "" || user.Role == "" {
		s.Log.Error("user data is empty")
		return errors.New("data is empty")
	}

	if user.ID == 0 {
		s.Log.Error("user id is empty")
		return errors.New("user id is empty")
	}

	if !strings.HasPrefix(user.Password, "$2a$") {
		// только если это НЕ хеш
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			s.Log.Error("failed to hash password", zap.Error(err))
			return errors.New("failed to hash password")
		}
		user.Password = string(hashedPassword)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Log.Error("failed to hash password", zap.Error(err))
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	res := s.repo.UpdateUser(user)
	if res != nil {
		s.Log.Error("user update failed", zap.Error(err))
		return err
	}
	return nil
}

func (s *userService) DeleteUserById(id int) error {
	if id == 0 {
		s.Log.Error("user id is empty")
		return errors.New("user id is empty")
	}
	err := s.repo.DeleteUserById(id)
	if err != nil {
		s.Log.Error("user delete failed", zap.Error(err))
		return err
	}
	return nil
}

//func (s *userService) Login(name, password, email string) (*models.User, error) {
//	user, err := s.repo.GetUserByEmail(email)
//	if err != nil {
//		s.Log.Error("user not found with this email", zap.Error(err))
//		return nil, err
//	}
//
//	if !CheckPassword(password, user.Password) {
//		s.Log.Error("invalid user password", zap.Error(err))
//		return nil, errors.New("invalid user password")
//	}
//	return user, nil
//}
