package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/quangdangfit/go-admin/app/dbs"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/repositories/impl"
	"github.com/quangdangfit/go-admin/app/schema"
)

type UserRepositoryTestSuite struct {
	suite.Suite

	db   dbs.IDatabase
	repo repositories.IUserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()

	s.db = dbs.NewDatabase()
	s.repo = impl.NewUserRepository(s.db)
}

func (s *UserRepositoryTestSuite) TestLoginSuccess() {
	item := &schema.LoginBodyParam{
		Username: "test-username-1",
		Password: "test-user-pwd-1",
	}

	user, err := s.repo.Login(item)
	s.Nil(err)
	s.NotNil(user)
}

func (s *UserRepositoryTestSuite) TestLoginFailed() {
	item := &schema.LoginBodyParam{
		Username: "test-username-1",
		Password: "wrong-password",
	}

	user, err := s.repo.Login(item)
	s.NotNil(err)
	s.Nil(user)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
