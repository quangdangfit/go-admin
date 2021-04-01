package api

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/logger"
	"github.com/quangdangfit/gocommon/validation"

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
// @Param body body schema.LoginBodyParams true "Body"
// @Success 200 {object} schema.BaseResponse
// @Router /login [post]
func (a *AuthAPI) Login(c *gin.Context) gohttp.Response {
	var params schema.LoginBodyParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	tokenInfo, err := a.service.Login(c, &params)
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
// @Param body body schema.RegisterBodyParams true "Body"
// @Success 200 {object} schema.BaseResponse
// @Router /register [post]
func (a *AuthAPI) Register(c *gin.Context) gohttp.Response {
	var params schema.RegisterBodyParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	ctx := c.Request.Context()
	tokenInfo, err := a.service.Register(ctx, &params)
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
// @Param body body schema.RefreshBodyParams true "Body"
// @Success 200 {object} schema.BaseResponse
// @Router /refresh [post]
func (a *AuthAPI) Refresh(c *gin.Context) gohttp.Response {
	var params schema.RefreshBodyParams
	if err := c.ShouldBindJSON(&params); err != nil {
		logger.Error(err.Error())
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	validator := validation.New()
	if err := validator.ValidateStruct(params); err != nil {
		return gohttp.Response{
			Error: errors.InvalidParams.New(),
		}
	}

	ctx := c.Request.Context()
	tokenInfo, err := a.service.Refresh(ctx, &params)
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
