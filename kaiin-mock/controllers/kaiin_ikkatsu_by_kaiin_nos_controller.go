package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kaiin-mock/models"
	"kaiin-mock/services"
)

type KaiinIkkatsuByKaiinNosController struct {
	svc *services.KaiinIkkatsuByKaiinNosService
}

func NewKaiinIkkatsuByKaiinNosController(svc *services.KaiinIkkatsuByKaiinNosService) *KaiinIkkatsuByKaiinNosController {
	return &KaiinIkkatsuByKaiinNosController{svc: svc}
}

func (c *KaiinIkkatsuByKaiinNosController) GetByKaiinNos(ctx *gin.Context) {
	kaiinNos := ctx.Query("kaiinNos")
	if kaiinNos == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "kaiinNos is required"})
		return
	}

	items, err := c.svc.GetByKaiinNos(kaiinNos)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}

	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "-1")

	// model: [] を返したいので nil を避ける
	if items == nil {
		items = []models.KaiinIkkatsuItem{}
	}

	ctx.JSON(http.StatusOK, models.KaiinIkkatsuByKaiinNosResponse{Model: items})
}
