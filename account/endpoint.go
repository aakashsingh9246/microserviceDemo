package account

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		ok, err := s.CreateUser(req.ID, req.Name, req.City, req.Phone, req.Sal)
		return CreateUserResponse{Ok: ok}, err
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		name, err := s.GetUser(ctx, req.ID)

		return GetUserResponse{
			Name: name,
		}, err
	}
}