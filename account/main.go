package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/satori/go.uuid"
)

type Service interface {
	CreateAccount(username string, password string) (ID, error)
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

type ID string
type username string
type password string

type Repository interface {
	Store(account *Account) error
}

type inmem struct{}

func (inmem) Store(account *Account) error {
	return nil
}

type Account struct {
	ID       ID
	username username
	password password
}

type createAccountRequest struct {
	User string `json:"username"`
	Pass string `json:"password"`
}

type createAccountResponse struct {
	ID  string `json:"id"`
	Err string `json:"err,omitempty"`
}

func makeCreateAccountEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createAccountRequest)
		v, err := svc.CreateAccount(req.User, req.Pass)
		if err != nil {
			return createAccountResponse{string(v), err.Error()}, nil
		}
		return createAccountResponse{string(v), ""}, nil
	}
}

func main() {
	repo := inmem{}
	svc := service{
		accounts: repo,
	}
	createAccountHandler := httptransport.NewServer(
		makeCreateAccountEndpoint(svc),
		decodeCreateAccountRequest,
		encodeResponse,
	)

	http.Handle("/account/create-account", createAccountHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func decodeCreateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
