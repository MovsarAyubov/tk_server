package api

import (
	"fmt"
	v1 "tk_app_v1/src/api/v1"

	"github.com/gin-gonic/gin"
)

func StartRouter(port string) {
	// Router
	router := gin.Default()

	router.GET("/api/v1/workers", v1.GetWorkers)
	router.POST("/api/v1/worker", v1.CreateWorker)
	router.DELETE("/api/v1/workers", v1.DeleteWorker)
	router.GET("/api/v1/tk_info", v1.GetTKInfo)
	router.GET("/api/v1/periods", v1.GetPeriods)
	router.GET("/api/v1/works", v1.GetWorkByPeriod)

	router.Run(fmt.Sprintf("0.0.0.0:%s", port))
	fmt.Printf("Running and listening to port %s!", port)
}
