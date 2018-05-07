package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type CreateAccountRequest struct {
	User string `json:"username"`
	Pass string `json:"password"`
}

type CreateAccountResponse struct {
	ID  string `json:"id"`
	Err string `json:"err,omitempty"`
}

func MakeCreateAccountEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		v, err := svc.CreateAccount(req.User, req.Pass)
		if err != nil {
			return CreateAccountResponse{string(v), err.Error()}, nil
		}
		return CreateAccountResponse{string(v), ""}, nil
	}
}
