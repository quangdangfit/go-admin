package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/quangdangfit/go-admin/app/mocks"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/repositories/impl"
)

type UserRepositoryTestSuite struct {
	suite.Suite

	mockDB *mocks.MockIDatabase
	repo   repositories.IUserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()
	s.mockDB = mocks.NewMockIDatabase(mockCtrl)

	s.repo = impl.NewUserRepository(s.mockDB)
}

func (s *UserRepositoryTestSuite) TestLogin() {
	//item := &schema.LoginBodyParam{
	//	Username: "test",
	//	Password: "test",
	//}
	//
	//s.mockDB.EXPECT().GetInstance().Return(&gorm.DB{}).Times(1)
	//
	//user, err := s.repo.Login(item)
	//s.NotNil(err)
	//s.Nil(user)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
