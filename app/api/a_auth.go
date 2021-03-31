package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/quangdangfit/gosdk/utils/logger"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/pkg/errors"
	gohttp "github.com/quangdangfit/go-admin/pkg/http"
)

// AuthAPI handle authentication api
type AuthAPI struct {
	service interfaces.IAuthService
}

// NewAuthAPI return new AuthAPI
func NewAuthAPI(service interfaces.IAuthService) *AuthAPI {
	return &AuthAPI{service: service}
}

// Login godoc
// @Tags Auth
// @Summary api login
// @Description api login
// @Accept  json
// @Produce json
// @Param body body schema.LoginBodyParam true "Body"
// @Success 200 {object} schema.BaseResponse
// @Router /login [post]
func (a *AuthAPI) Login(c *gin.Context) gohttp.Response {
	var item schema.LoginBodyParam
	if err := c.ShouldBindJSON(&item); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	ctx := c.Request.Context()
	tokenInfo, err := a.service.Login(ctx, &item)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  tokenInfo,
	}
}

// Register godoc
// @Tags Auth
// @Summary api register
// @Description api register
// @Accept  json
// @Produce json
// @Param body body schema.RegisterBodyParam true "Body"
// @Success 200 {object} schema.BaseResponse
// @Router /register [post]
func (a *AuthAPI) Register(c *gin.Context) gohttp.Response {
	var item schema.RegisterBodyParam
	if err := c.ShouldBindJSON(&item); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	ctx := c.Request.Context()
	tokenInfo, err := a.service.Register(ctx, &item)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  tokenInfo,
	}
}

// Refresh godoc
// @Tags Auth
// @Summary api refresh token
// @Description api refresh token
// @Accept  json
// @Produce json
// @Param body body schema.RefreshBodyParam true "Body"
// @Success 200 {object} schema.BaseResponse
// @Router /refresh [post]
func (a *AuthAPI) Refresh(c *gin.Context) gohttp.Response {
	var bodyParam schema.RefreshBodyParam
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	validate := validator.New()
	if err := validate.Struct(bodyParam); err != nil {
		logger.Error("Body is invalid: ", err)
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	ctx := c.Request.Context()
	tokenInfo, err := a.service.Refresh(ctx, &bodyParam)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
		Data:  tokenInfo,
	}
}

// Logout godoc
// @Tags Auth
// @Summary api logout
// @Description api logout
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} schema.BaseResponse
// @Router /logout [post]
func (a *AuthAPI) Logout(c *gin.Context) gohttp.Response {
	err := a.service.Logout(c)
	if err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: err,
		}
	}

	return gohttp.Response{
		Error: errors.Success.New(),
	}
}
