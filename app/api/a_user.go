package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/quangdangfit/gocommon/logger"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/pkg/errors"
	gohttp "github.com/quangdangfit/go-admin/pkg/http"
	"github.com/quangdangfit/go-admin/pkg/utils"
)

// UserAPI handle user api
type UserAPI struct {
	service interfaces.IUserService
}

// NewUserAPI return new UserAPI pointer
func NewUserAPI(service interfaces.IUserService) *UserAPI {
	return &UserAPI{service: service}
}

// GetByID get user by id
func (u *UserAPI) GetByID(c *gin.Context) {
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

// List user by query
func (u *UserAPI) List(c *gin.Context) gohttp.Response {
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
