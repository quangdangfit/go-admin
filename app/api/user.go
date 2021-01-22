package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/quangdangfit/gosdk/utils/logger"

	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/app/services"
	"github.com/quangdangfit/go-admin/pkg/errors"
	gohttp "github.com/quangdangfit/go-admin/pkg/http"
	"github.com/quangdangfit/go-admin/pkg/utils"
)

type User struct {
	service services.IUserService
}

func NewUserAPI(service services.IUserService) *User {
	return &User{service: service}
}

func (u *User) validate(r schema.RegisterBodyParam) bool {
	return utils.Validate(
		[]utils.Validation{
			{Value: r.Username, Valid: "username"},
			{Value: r.Email, Valid: "email"},
			{Value: r.Password, Valid: "password"},
		})
}

func (u *User) checkPermission(id string, data map[string]interface{}) bool {
	return data["id"] == id
}

func (u *User) GetByID(c *gin.Context) {
	userID := c.Param("id")
	ctx := c.Request.Context()
	user, err := u.service.GetByID(ctx, userID)
	if err != nil {
		err = errors.Wrap(err, "API.GetByID")
		logger.Error("Failed to get user: ", err)
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), utils.ErrorGetDatabase))
		return
	}

	var res schema.User
	copier.Copy(&res, &user)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}

func (u *User) List(c *gin.Context) gohttp.Response {
	var queryParam schema.UserQueryParam
	if err := c.ShouldBindQuery(&queryParam); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	user, err := u.service.List(c, &queryParam)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	var res []schema.User
	copier.Copy(&res, &user)
	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  res,
	}
}
