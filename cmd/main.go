package main

import (
	"log"
	"net/http"

	encode "github.com/gfelixc/goddd-pet-store/account/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/gfelixc/goddd-pet-store/account"
	"github.com/gfelixc/goddd-pet-store/account/mock"
)

func main() {
	repo := mock.AccountRepository{}
	svc := account.NewService(repo)
	createAccountHandler := httptransport.NewServer(
		account.MakeCreateAccountEndpoint(svc),
		encode.DecodeCreateAccountRequest,
		encode.EncodeResponse,
	)

	http.Handle("/account/create-account", createAccountHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
