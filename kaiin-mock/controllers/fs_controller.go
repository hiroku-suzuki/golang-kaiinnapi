package controllers

import (
	"net/http"

	"kaiin-mock/models"
	"kaiin-mock/services"

	"github.com/gin-gonic/gin"
)

type FSController struct {
	svc *services.FSService
}

func NewFSController(svc *services.FSService) *FSController {
	return &FSController{svc: svc}
}

func (c *FSController) GetFS(ctx *gin.Context) {
	kaiinNo := ctx.Param("kaiinNo")

	// no-cache（必要分だけ）
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "-1")

	model, ok := c.svc.FindByKaiinNo(kaiinNo)
	if !ok {
		ctx.JSON(http.StatusNotFound, models.FSResponse{
			Model: nil,
			Errors: []models.ErrorItem{{
				Message: "kaiinNo not found",
				Fields:  []string{"kaiinNo"},
			}},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.FSResponse{Model: model})
}
