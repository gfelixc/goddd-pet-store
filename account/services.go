package account

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type Service interface {
	CreateAccount(username string, password string) (ID, error)
}

func NewService(accounts Repository) Service {
	return &service{
		accounts: accounts,
	}
}

type service struct {
	accounts Repository
}

func (s service) CreateAccount(user string, pass string) (ID, error) {
	if user == "" || pass == "" {
		return ID(""), ErrInvalidArgument
	}
	newAccount := &Account{
		ID:       ID(uuid.NewV4().String()),
		username: username(user),
		password: password(pass),
	}
	if err := s.accounts.Store(newAccount); err != nil {
		return ID(""), err
	}

	return newAccount.ID, nil
}

var (
	ErrInvalidArgument = errors.New("Invalid argument")
)
