package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gosdk/utils/logger"
	"go.uber.org/dig"

	"github.com/quangdangfit/go-admin/app"
	"github.com/quangdangfit/go-admin/app/models"
	"github.com/quangdangfit/go-admin/app/repositories"
	"github.com/quangdangfit/go-admin/app/router"
	"github.com/quangdangfit/go-admin/pkg/jwt"
)

var (
	engine    *gin.Engine
	container *dig.Container
	token     string

	apiKey = "testAPIkey"

	roles = []*models.Role{
		{
			Model: models.Model{
				ID:        "test-role-id-1",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name:        "test1",
			Description: "test1",
		},
	}
)

func TestMain(m *testing.M) {
	logger.Initialize(false)
	container = app.BuildContainer()
	engine = router.InitGinEngine(container)

	setup()
	code := m.Run()
	teardown()

	os.Exit(code)
}

func setup() {
	logger.Info("============> Setup for testing")
	removeUserData()

	createRoleData()
	createUserData()
	setupToken()
}

func teardown() {
	logger.Info("============> Teardown")
	removeUserData()
}

func createRoleData() {
	container.Invoke(func(
		userRepo repositories.IRoleRepository,
		jwtauth jwt.IJWTAuth,
	) error {
		for _, r := range roles {
			if err := userRepo.Create(r); err != nil {
				return err
			}
		}

		return nil
	})
}

func createUserData() {
	container.Invoke(func(
		userRepo repositories.IUserRepository,
		jwtauth jwt.IJWTAuth,
	) error {
		for _, u := range users {
			if err := userRepo.Create(u); err != nil {
				continue
			}
		}

		return nil
	})
}

func setupToken() {
	container.Invoke(func(
		jwtauth jwt.IJWTAuth,
	) error {
		tokenInfo, err := jwtauth.GenerateToken(user.ID)
		if err != nil {
			return err
		}

		token = tokenInfo.GetAccessToken()

		return nil
	})
}

func removeUserData() {
	//container.Invoke(func(
	//	userRepo repositories.IUserRepository,
	//) error {
	//	for _, u := range users {
	//		err := userRepo.Delete(u.ID)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})
}

//func clearRedis() {
//	container.Invoke(func(
//		redis redis.IRedis,
//	) error {
//		err := redis.Remove(utils.GenAPIKeyRedis("", apiKey))
//		if err != nil {
//			return err
//		}
//		return nil
//	})
//}

func toReader(v interface{}) io.Reader {
	buf := new(bytes.Buffer)
	_ = json.NewEncoder(buf).Encode(v)
	return buf
}

func parseReader(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func newGetRequest(formatRouter string, v interface{}, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf(formatRouter, args...), toReader(v))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return req
}

func newPostRequest(formatRouter string, v interface{}, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("POST", fmt.Sprintf(formatRouter, args...), toReader(v))
	return req
}

func newPostAuthRequest(formatRouter string, v interface{}, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("POST", fmt.Sprintf(formatRouter, args...), toReader(v))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return req
}

func newInternalPostRequest(formatRouter string, v interface{}, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("POST", fmt.Sprintf(formatRouter, args...), toReader(v))
	req.Header.Add("api-key", apiKey)
	return req
}

func newPutRequest(formatRouter string, v interface{}, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("PUT", fmt.Sprintf(formatRouter, args...), toReader(v))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return req
}

func newDeleteRequest(formatRouter string, args ...interface{}) *http.Request {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf(formatRouter, args...), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return req
}
