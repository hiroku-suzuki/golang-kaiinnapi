package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kaiin-mock/models"
	"kaiin-mock/services"
)

type KaiintenShokaiKeiyakuController struct {
	svc *services.KaiintenShokaiKeiyakuService
}

func NewKaiintenShokaiKeiyakuController(svc *services.KaiintenShokaiKeiyakuService) *KaiintenShokaiKeiyakuController {
	return &KaiintenShokaiKeiyakuController{svc: svc}
}

func (c *KaiintenShokaiKeiyakuController) GetKSK(ctx *gin.Context) {
	kaiinNo := ctx.Param("kaiinNo")

	// no-cache（必要分だけ）
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "-1")

	model, ok := c.svc.FindByKaiinNo(kaiinNo)
	if !ok {
		ctx.JSON(http.StatusNotFound, models.KaiintenShokaiKeiyakuByKaiinNoResponse{
			Model: nil,
			Errors: []models.ErrorItem{{
				Message: "kaiinNo not found",
				Fields:  []string{"kaiinNo"},
			}},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.KaiintenShokaiKeiyakuByKaiinNoResponse{Model: model})
}
