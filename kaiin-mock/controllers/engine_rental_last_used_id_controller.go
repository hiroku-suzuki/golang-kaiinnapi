package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kaiin-mock/models"
	"kaiin-mock/services"
)

type EngineRentalLastUsedEngineRentalIdController struct {
	svc *services.EngineRentalLastUsedEngineRentalIdService
}

func NewEngineRentalLastUsedEngineRentalIdController(svc *services.EngineRentalLastUsedEngineRentalIdService) *EngineRentalLastUsedEngineRentalIdController {
	return &EngineRentalLastUsedEngineRentalIdController{svc: svc}
}

func (c *EngineRentalLastUsedEngineRentalIdController) GetLastUsedEngineRentalId(ctx *gin.Context) {
	kaiinNo := ctx.Param("kaiinNo")

	v, err := c.svc.GetByKaiinNo(kaiinNo)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "-1")

	ctx.JSON(http.StatusOK, models.EngineRentalLastUsedEngineRentalIdResponse{Model: v})
}
