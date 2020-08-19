package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
// @Tags Private
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
// @Tags Private
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
