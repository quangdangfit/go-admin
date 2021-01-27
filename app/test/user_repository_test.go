package test

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/quangdangfit/go-admin/app/dbs"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/repositories/impl"
	"github.com/quangdangfit/go-admin/app/schema"
)

var (
	users = []*models.User{
		{
			Model: models.Model{
				ID:        "test-user-id-1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Username:     "test-username-1",
			Email:        "testuseremail1@tokoin.io",
			Password:     "test-user-pwd-1",
			RefreshToken: "test-user-refresh-token-1",
			RoleID:       roles[0].ID,
		},
		{
			Model: models.Model{
				ID:        "test-user-id-2",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Username:     "test-username-2",
			Email:        "testuseremail2@tokoin.io",
			Password:     "test-user-pwd-2",
			RefreshToken: "test-user-refresh-token-2",
			RoleID:       roles[0].ID,
		},
		{
			Model: models.Model{
				ID:        "test-user-id-3",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Username:     "test-username-3",
			Email:        "testuseremail3@tokoin.io",
			Password:     "test-user-pwd-3",
			RefreshToken: "test-user-refresh-token-3",
			RoleID:       roles[0].ID,
		},
	}

	user = users[0]
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

func (s *UserRepositoryTestSuite) TestGetByIDSuccess() {
	u, err := s.repo.GetByID(user.ID)
	s.Nil(err)
	s.NotNil(u)
	s.Equal(user.ID, u.ID)
}

func (s *UserRepositoryTestSuite) TestGetByIDNotFound() {
	u, err := s.repo.GetByID("not-found-id")
	s.NotNil(err)
	s.Nil(u)
}

func (s *UserRepositoryTestSuite) TestGetByTokenSuccess() {
	u, err := s.repo.GetUserByToken(user.RefreshToken)
	s.Nil(err)
	s.NotNil(u)
	s.Equal(user.ID, u.ID)
}

func (s *UserRepositoryTestSuite) TestGetByTokenNotFound() {
	u, err := s.repo.GetUserByToken("not-found-token")
	s.NotNil(err)
	s.Nil(u)
}

func (s *UserRepositoryTestSuite) TestListFull() {
	usrs, err := s.repo.List(&schema.UserQueryParam{
		Offset: 0,
		Limit:  100000,
	})
	s.Nil(err)
	s.NotNil(usrs)
	s.Equal(len(users), len(*usrs))
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

func (s *UserRepositoryTestSuite) TestRemoveTokenSuccess() {
	u, err := s.repo.RemoveToken(users[1].ID)
	s.Nil(err)
	s.NotNil(u)

	u, err = s.repo.GetByID(users[1].ID)
	s.Nil(err)
	s.NotNil(u)
	s.Equal("", u.RefreshToken)
}

func (s *UserRepositoryTestSuite) TestRemoveTokenNotFound() {
	_, err := s.repo.RemoveToken("not-found-id")
	s.Nil(err)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
