package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kaiin-mock/models"
	"kaiin-mock/services"
)

type BukkakuController struct {
	svc *services.BukkakuService
}

func NewBukkakuController(svc *services.BukkakuService) *BukkakuController {
	return &BukkakuController{svc: svc}
}

func (c *BukkakuController) GetBukkaku(ctx *gin.Context) {
	kaiinNo := ctx.Param("kaiinNo")

	m, err := c.svc.GetByKaiinNo(kaiinNo)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "-1")

	ctx.JSON(http.StatusOK, models.BukkakuResponse{Model: m})
}