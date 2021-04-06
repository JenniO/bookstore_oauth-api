package access_token

import (
	"github.com/JenniO/bookstore_oauth-api/src/domain/access_token"
	"github.com/JenniO/bookstore_oauth-api/src/repository/db"
	"github.com/JenniO/bookstore_oauth-api/src/repository/rest"
	"github.com/JenniO/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type (
	Repository interface {
		GetById(string) (*access_token.AccessToken, *errors.RestErr)
		Create(access_token.AccessToken) *errors.RestErr
		UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
	}
	Service interface {
		GetById(string) (*access_token.AccessToken, *errors.RestErr)
		Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
		UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
	}
	service struct {
		dbRepo        db.DbRepository
		restUsersRepo rest.RestUsersRepository
	}
)

func NewService(usersRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		dbRepo:        dbRepo,
		restUsersRepo: usersRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if accessTokenId == "" {
		err := errors.NewBadRequestError("invalid access token id")
		return nil, err
	}
	return s.dbRepo.GetById(accessTokenId)
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
