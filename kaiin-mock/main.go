package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"kaiin-mock/controllers"
	"kaiin-mock/repositories"
	"kaiin-mock/services"
)

func main() {
	r := gin.Default()

	// api/AtbbKeiyaku/:kaiinNo
	repo := repositories.NewInMemoryAtbbKeiyakuRepository()
	svc := services.NewAtbbKeiyakuService(repo)
	ctrl := controllers.NewAtbbKeiyakuController(svc)

	// api/Kaiin/:kaiinNo
	kaiinRepo := repositories.NewInMemoryKaiinRepository()
	kaiinSvc := services.NewKaiinService(kaiinRepo)
	kaiinCtrl := controllers.NewKaiinController(kaiinSvc)

	// ルーティング
	api := r.Group("/api")
	{
		api.GET("/AtbbKeiyaku/:kaiinNo", ctrl.GetAtbbKeiyaku)
		api.GET("/Kaiin/:kaiinNo", kaiinCtrl.GetKaiin)
	}

	addr := ":8080"
	log.Printf("kaiin-mock server starting on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
