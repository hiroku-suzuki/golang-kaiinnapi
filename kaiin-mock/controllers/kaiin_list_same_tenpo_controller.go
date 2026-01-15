package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kaiin-mock/models"
	"kaiin-mock/services"
)

type KaiinListSameTenpoKaiinNoController struct {
	svc *services.KaiinListSameTenpoKaiinNoService
}

func NewKaiinListSameTenpoKaiinNoController(svc *services.KaiinListSameTenpoKaiinNoService) *KaiinListSameTenpoKaiinNoController {
	return &KaiinListSameTenpoKaiinNoController{svc: svc}
}

func (c *KaiinListSameTenpoKaiinNoController) GetSameTenpoKaiinNo(ctx *gin.Context) {
	kaiinNo := ctx.Query("kaiinNo")
	if kaiinNo == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "kaiinNo is required"})
		return
	}

	list, err := c.svc.GetByKaiinNo(kaiinNo)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "-1")

	ctx.JSON(http.StatusOK, models.KaiinListSameTenpoKaiinNoResponse{Model: list})
}
