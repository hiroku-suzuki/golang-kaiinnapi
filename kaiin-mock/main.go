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

	// /api/FactSheetKeiyaku/:kaiinNo
	fsRepo := repositories.NewInMemoryFSRepository()
	fsSvc := services.NewFSService(fsRepo)
	fsCtrl := controllers.NewFSController(fsSvc)

	// api/KaiintenShokaiKeiyaku/ByKaiinNo/:kaiinNo
	kskRepo := repositories.NewInMemoryKaiintenShokaiKeiyakuRepository()
	kskSvc := services.NewKaiintenShokaiKeiyakuService(kskRepo)
	kskCtrl := controllers.NewKaiintenShokaiKeiyakuController(kskSvc)

	// ルーティング
	api := r.Group("/api")
	{
		api.GET("/AtbbKeiyaku/:kaiinNo", ctrl.GetAtbbKeiyaku)
		api.GET("/Kaiin/:kaiinNo", kaiinCtrl.GetKaiin)
		api.GET("/FactSheetKeiyaku/:kaiinNo", fsCtrl.GetFS)
		api.GET("/KaiintenShokaiKeiyaku/ByKaiinNo/:kaiinNo", kskCtrl.GetKSK)
	}

	addr := ":8080"
	log.Printf("kaiin-mock server starting on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
