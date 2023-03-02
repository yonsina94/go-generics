package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yonsina94/go-generics/commons/service"
)

type BaseController[TModel any] struct {
	Service *service.BaseService[TModel]
}

func (ctr *BaseController[TModel]) getAll(c *gin.Context) {
	if models, err := ctr.Service.GetAll(); err == nil {
		c.JSON(http.StatusOK, models)
	} else {
		c.Error(err)
	}
}

func (ctr *BaseController[TModel]) getById(c *gin.Context) {

	if id, err := strconv.ParseUint(c.Param("id"), 10, 10); err == nil {
		if model, err := ctr.Service.GetByID(uint(id)); err == nil {
			c.JSON(http.StatusOK, model)
		} else {
			c.Error(err)
		}
	} else {
		c.Error(err)
	}
}

func (ctr *BaseController[TModel]) save(c *gin.Context) {
	var models []TModel
	if err := c.ShouldBindJSON(&models); err == nil {
		if inserted, err := ctr.Service.Save(models...); err == nil {
			c.JSON(http.StatusOK, map[string]any{"inserted": inserted})
		} else {
			c.Error(err)
		}
	} else {
		c.Error(err)
	}
}

func (ctr *BaseController[TModel]) deleteById(c *gin.Context) {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 10); err == nil {
		if deleted, err := ctr.Service.DeleteById(uint(id)); err == nil {
			c.JSON(http.StatusOK, map[string]any{"deleted": deleted})
		} else {
			c.Error(err)
		}
	} else {
		c.Error(err)
	}
}

func (ctr *BaseController[TModel]) AssingEndpoints(router *gin.RouterGroup) {
	router.GET("/", ctr.getAll).GET("/:id", ctr.getById).POST("/", ctr.save).DELETE("/:id", ctr.deleteById)
}
