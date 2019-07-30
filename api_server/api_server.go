package api_server

import (
	"NBA-master/api_server/players"
	"NBA-master/api_server/teams"
	_ "NBA-master/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type APIServer struct {
	engine *gin.Engine
}

func (a *APIServer) registry() {
	APIServerInit(a.engine)
}

func (a *APIServer) Start() {
	a.registry()
	a.engine.Run(":5000")
}

func APIServerInit(r *gin.Engine) {
	// docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/v1/api")
	{
		teamsRegistry(v1)
		playersRegistry(v1)
	}
}

func teamsRegistry(r *gin.RouterGroup) {
	r.GET("/teams", teams.TeamListHandler)
	r.GET("/teams/:teamID", teams.TeamHandler)
}

func playersRegistry(r *gin.RouterGroup) {
	r.GET("/players", players.ShowAllPlayersHandler)
	r.GET("/players/:playerID", players.ShowPlayersHandler)
}

func New() *APIServer {
	return &APIServer{
		engine: gin.Default(),
	}
}
