package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/quangdangfit/gocommon/utils/logger"

	"go-admin/app/schema"
	"go-admin/app/services"
	"go-admin/pkg/utils"
)

type Auth struct {
	service services.IAuthService
}

func NewAuthAPI(service services.IAuthService) *Auth {
	return &Auth{service: service}
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
func (a *Auth) Login(c *gin.Context) {
	var item schema.LoginBodyParam
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	tokenInfo, err := a.service.Login(ctx, &item)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.PrepareResponse(tokenInfo, "OK", ""))
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
func (a *Auth) Register(c *gin.Context) {
	var item schema.RegisterBodyParam
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	tokenInfo, err := a.service.Register(ctx, &item)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.PrepareResponse(tokenInfo, "OK", ""))
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
func (a *Auth) Refresh(c *gin.Context) {
	var bodyParam schema.RefreshBodyParam
	if err := c.ShouldBindJSON(&bodyParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(bodyParam); err != nil {
		logger.Error("Body is invalid: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	tokenInfo, err := a.service.Refresh(ctx, &bodyParam)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.PrepareResponse(tokenInfo, "OK", ""))
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
func (a *Auth) Logout(c *gin.Context) {
	err := a.service.Logout(c)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	c.JSON(http.StatusOK, utils.PrepareResponse(nil, "OK", ""))
}
