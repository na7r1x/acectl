package main

import (
	"github.com/gin-gonic/gin"
	"github.com/na7r1x/acectl/internal/core/service/brokersrv"
	brokerhdl "github.com/na7r1x/acectl/internal/handlers/brokerhdl/http"
	"github.com/na7r1x/acectl/internal/repositories/brokerrepo"
)

func main() {
	brokerRepository := brokerrepo.NewSqliteRepo("acectl.db")
	brokerService := brokersrv.New(brokerRepository)
	brokerHandler := brokerhdl.NewHttpHandler(brokerService)

	router := gin.New()
	router.POST("/register", brokerHandler.Register)
	router.GET("/broker", brokerHandler.List)
	router.GET("/broker/:brokerId", brokerHandler.Get)
	router.DELETE("/broker/:brokerId", brokerHandler.Unregister)
	router.GET("/broker/:brokerId/status", brokerHandler.Status)
	router.GET("/broker/:brokerId/stop", brokerHandler.Stop)
	router.GET("/broker/:brokerId/start", brokerHandler.Start)

	router.Run(":8080")
}
