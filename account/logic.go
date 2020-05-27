package account

import (
	
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}


func (s service) CreateUser(id,phone,sal int, city,name string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	user := User{
		ID:   id,
		Name: name,
		City: city,
		Phone: phone,
		Sal: sal
	}

	if err := s.repostory.CreateUser(user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return "Success", nil
}

func (s service) GetUser(id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	name, err := s.repostory.GetUser(id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get user", id)

	return name, nil
}