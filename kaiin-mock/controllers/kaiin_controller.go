package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kaiin-mock/models"
	"kaiin-mock/repositories"
	"kaiin-mock/services"
)

type KaiinController struct {
	svc services.KaiinService
}

func NewKaiinController(svc services.KaiinService) *KaiinController {
	return &KaiinController{svc: svc}
}

func (ctl *KaiinController) GetKaiin(c *gin.Context) {
	kaiinNo := c.Param("kaiinNo")

	// 必要最小限の no-cache 系だけ付与
	c.Header("Cache-Control", "no-cache")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "-1")

	model, err := ctl.svc.GetKaiin(kaiinNo)
	if err != nil {
		if err == repositories.ErrNotFound {
			c.JSON(http.StatusNotFound, models.KaiinResponse{
				Model: nil,
				Errors: []models.ErrorItem{
					{Message: "kaiinNo not found", Fields: []string{"kaiinNo"}},
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.KaiinResponse{
			Model:  nil,
			Errors: []models.ErrorItem{{Message: "internal error"}},
		})
		return
	}

	c.JSON(http.StatusOK, models.KaiinResponse{Model: model})
}
