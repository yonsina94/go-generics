package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yonsina94/go-generics/commons/controller"
)

type UserController struct {
	controller.BaseController[User]
}

func (ctr *UserController) getByEmail(c *gin.Context) {
	if users, err := ctr.Service.GetBy(User{Email: c.Param("email")}); err == nil {
		c.JSON(http.StatusOK, users)
	} else {
		c.Error(err)
	}
}

func (ctr *UserController) getByUserName(c *gin.Context) {
	if users, err := ctr.Service.GetBy(User{UserName: c.Param("username")}); err == nil {
		c.JSON(http.StatusOK, users)
	} else {
		c.Error(err)
	}
}

func (ctr *UserController) AssingEndpoints(router *gin.RouterGroup) {
	ctr.BaseController.AssingEndpoints(router)

	router.GET("/by/email/:email", ctr.getByEmail).GET("/by/username/:username", ctr.getByUserName)
}

func Instance(service *UserService) *UserController {
	return &UserController{
		BaseController: controller.BaseController[User]{
			Service: &service.BaseService,
		},
	}
}
