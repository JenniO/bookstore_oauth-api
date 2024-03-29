package rest

import (
	"encoding/json"
	"github.com/JenniO/bookstore_oauth-api/src/domain/users"
	"github.com/JenniO/bookstore_oauth-api/src/utils/errors"
	"github.com/federicoleon/golang-restclient/rest"
	"time"
)

type (
	RestUsersRepository interface {
		LoginUser(string, string) (*users.User, *errors.RestErr)
	}
	usersRepository struct {
	}
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8082",
		Timeout: 100 * time.Millisecond,
	}
)

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response, when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}
