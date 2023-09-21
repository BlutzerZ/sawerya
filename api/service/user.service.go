package service

import (
	"blutzerz/sawerya/api/dto"
	"blutzerz/sawerya/api/models"
	"blutzerz/sawerya/api/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService() *UserService {
	repo := repository.NewUserRepository()
	return &UserService{
		repository: repo,
	}
}

func (s *UserService) CreateUser(req *dto.RegisterUserRequest) error {
	u := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repository.Create(u)

	return err
}

func (s *UserService) GetUserByID(ID uint) (models.User, error) {
	user, err := s.repository.FindByID(ID)

	return user, err
}

func (s *UserService) GetUserByUsername(username string) (models.User, error) {
	user, err := s.repository.FindByUsername(username)

	return user, err
}

func (s *UserService) GetAllUser() ([]models.User, error) {
	users, err := s.repository.FindAll()

	return users, err
}

func (s *UserService) UpdateUsername(ID uint, req *dto.UpdateUsernameRequest) error {
	var err error

	// validate password
	isMatch, err := s.isPasswordMatch(ID, req.Password)
	if err != nil {
		return err
	}
	if !isMatch {
		msg := "invalid password"
		err = errors.New(msg)

		return err
	}
	// update username
	user := models.User{
		ID:       ID,
		Username: req.Username,
	}

	err = s.repository.Update(&user)

	return err
}

func (s *UserService) UpdatePassword(ID uint, req *dto.UpdatePasswordRequest) error {

	// validate password
	isMatch, err := s.isPasswordMatch(ID, req.OldPassword)
	if err != nil {
		return err
	}
	if !isMatch {
		msg := "invalid password"
		err = errors.New(msg)

		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// update password
	user := models.User{
		ID:       ID,
		Password: string(hash),
	}
	err = s.repository.Update(&user)

	return err
}

func (s *UserService) DeleteUser(ID uint, req *dto.DeleteUserRequest) error {
	isMatch, err := s.isPasswordMatch(ID, req.Password)
	if err != nil {
		return err
	}
	if !isMatch {
		msg := "invalid password"
		err = errors.New(msg)

		return err
	}

	err = s.repository.Delete(ID)

	return err
}

func (s *UserService) isPasswordMatch(ID uint, inputPassword string) (bool, error) {
	var user models.User
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
