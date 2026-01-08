package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kaiin-mock/models"
	"kaiin-mock/repositories"
	"kaiin-mock/services"
)

type AtbbKeiyakuController struct {
	svc services.AtbbKeiyakuService
}

func NewAtbbKeiyakuController(svc services.AtbbKeiyakuService) *AtbbKeiyakuController {
	return &AtbbKeiyakuController{svc: svc}
}

func (ctl *AtbbKeiyakuController) GetAtbbKeiyaku(c *gin.Context) {
	kaiinNo := c.Param("kaiinNo")

	// レスポンスヘッダ（提示例に“寄せる”ために付与）
	c.Header("cache-control", "no-cache")
	c.Header("pragma", "no-cache")
	c.Header("expires", "-1")
	c.Header("server", "Gin-gonic")
	// c.Header("x-aspnet-version", "x.x.xxxxx")
	// c.Header("x-powered-by", "ASP.NET")
	c.Header("content-type", "application/json; charset=utf-8")

	model, err := ctl.svc.GetAtbbKeiyaku(kaiinNo)
	if err != nil {
		if err == repositories.ErrNotFound {
			c.JSON(http.StatusNotFound, models.AtbbKeiyakuResponse{
				Model: nil,
				Errors: []models.ErrorItem{
					{Message: "kaiinNo not found", Fields: []string{"kaiinNo"}},
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.AtbbKeiyakuResponse{
			Model: nil,
			Errors: []models.ErrorItem{
				{Message: "internal error"},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.AtbbKeiyakuResponse{
		Model: model,
	})
}
