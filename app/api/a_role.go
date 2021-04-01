package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/quangdangfit/gosdk/utils/logger"

	"github.com/quangdangfit/go-admin/app/interfaces"
	"github.com/quangdangfit/go-admin/app/schema"
	"github.com/quangdangfit/go-admin/pkg/utils"
)

// RoleAPI handle role api
type RoleAPI struct {
	service interfaces.IRoleService
}

// NewRoleAPI return new RoleAPI pointer
func NewRoleAPI(service interfaces.IRoleService) *RoleAPI {
	return &RoleAPI{service: service}
}

// CreateRole create new role
func (r *RoleAPI) CreateRole(c *gin.Context) {
	var item schema.RoleBodyParam
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(item)
	if err != nil {
		logger.Error("Request body is invalid: ", err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	ctx := c.Request.Context()
	user, err := r.service.CreateRole(ctx, &item)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, err.Error(), ""))
		return
	}

	var res schema.Role
	copier.Copy(&res, &user)
	c.JSON(http.StatusOK, utils.PrepareResponse(res, "OK", ""))
}
