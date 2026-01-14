package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kaiin-mock/models"
	"kaiin-mock/services"
)

type HpToolPlanKeiyakuController struct {
	svc *services.HpToolPlanKeiyakuService
}

func NewHpToolPlanKeiyakuController(svc *services.HpToolPlanKeiyakuService) *HpToolPlanKeiyakuController {
	return &HpToolPlanKeiyakuController{svc: svc}
}

func (c *HpToolPlanKeiyakuController) GetHpToolPlanKeiyaku(ctx *gin.Context) {
	kaiinNo := ctx.Param("kaiinNo")

	m, err := c.svc.GetByKaiinNo(kaiinNo)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "-1")

	ctx.JSON(http.StatusOK, models.HpToolPlanKeiyakuResponse{Model: m})
}
