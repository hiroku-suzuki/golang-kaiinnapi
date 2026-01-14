package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"kaiin-mock/models"
	"kaiin-mock/services"
)

type EngineRentalController struct {
	svc *services.EngineRentalService
}

func NewEngineRentalController(svc *services.EngineRentalService) *EngineRentalController {
	return &EngineRentalController{svc: svc}
}

func (c *EngineRentalController) GetByControlPanelKaiinNo(ctx *gin.Context) {
	kaiinNo := ctx.Param("kaiinNo")

	includeHistory := parseBoolQuery(ctx.Query("includeHistory"))
	includeReserve := parseBoolQuery(ctx.Query("includeReserve"))
	validDate := ctx.Query("validDate")

	m, err := c.svc.GetByControlPanelKaiinNo(kaiinNo, includeHistory, includeReserve, validDate)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}

	// よく使う no-cache ヘッダ（必要なものだけ）
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "-1")

	ctx.JSON(http.StatusOK, models.EngineRentalResponse{Model: m})
}

func parseBoolQuery(v string) bool {
	if v == "" {
		return false
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return b
}
