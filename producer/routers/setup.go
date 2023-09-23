package routers

import (
	"github.com/AJackTi/rabbitmq-go-example/producer/controllers"
	"github.com/gin-gonic/gin"
)

// Function to setup routers and router groups
func SetupRouters(app *gin.Engine) {
	v1 := app.Group("/v1")
	{
		v1.GET("/ping", controllers.Ping)
		v1.POST("/publish/example", controllers.Example)
	}
	// Standalone route example
	// app.GET("/ping", controllers.Ping)
}
