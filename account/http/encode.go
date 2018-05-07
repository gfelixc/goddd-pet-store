package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gfelixc/goddd-pet-store/account"
)

func DecodeCreateAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request account.CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
