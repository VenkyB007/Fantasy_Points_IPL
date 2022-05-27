package main

import (
	_ "fantasy_league_service/docs"
	"fantasy_league_service/handlers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Fantasy League Service API
// @version 1.0
// @description In this Service, Ranked Cap-holders and League Scores are shown
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9090
// @BasePath /
func main() {

	r := gin.Default()

	r.GET("/cap-holders", handlers.CapHodler)
	r.GET("/fantasy-scores", handlers.FantasyScores)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":9090")
}
