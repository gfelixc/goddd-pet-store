package mock

import "github.com/gfelixc/goddd-pet-store/account"

type AccountRepository struct{}

func (AccountRepository) Store(account *account.Account) error {
	return nil
}
