package account

type ID string
type username string
type password string

type Repository interface {
	Store(account *Account) error
}


type Account struct {
	ID       ID
	username username
	password password
}
