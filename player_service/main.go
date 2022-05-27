package main

import (
	"playerservice/handlers"

	_ "github.com/lib/pq"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "playerservice/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Player Service API
// @version 1.0
// @description In this Service, You can Post, Get, Delete Player details
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8989
// @BasePath /

func main() {

	r := gin.Default()

	r.POST("/player", handlers.CreatePlayer)
	r.POST("/player/:id/score", handlers.PostScores)

	r.GET("/players", handlers.GetPlayers)
	r.GET("/players/scores", handlers.GetScores)

	r.DELETE("/player/:id", handlers.DeletePlayer)
	r.DELETE("player/:id/score", handlers.DeletePlayerScores)

	r.GET("/players/flp", handlers.GetflpPlayers)
	r.GET("/players/flpscores", handlers.GetflpScores)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8989")
}
