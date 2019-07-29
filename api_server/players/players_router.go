package players

import (
	initiator "NBA-master/init"
	"NBA-master/model"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

)

var (
	ErrorPlayerId    = errors.New("player id not allow")
	ErrorPlayerParam = errors.New("player param is not correct")
)

func ShowPlayersHandler(c *gin.Context) {
	id := c.Param("playerID")

	number, _ := strconv.Atoi(id)
	if number > 152 || number < 0 {
		c.AbortWithError(400, ErrorPlayerId)
		return
	}
	var player model.Playerdata
	if dbError := initiator.POSTGRES.Where("id=?", id).Find(&player).Error; dbError != nil {
		c.AbortWithError(400, dbError)
		return
	}
	c.JSON(http.StatusOK, player.Serializer())
}

type ListPlayerParm struct {
	Search string `form:"search"`
	Return string `form:"return"`
	Team   string `form:"team"`
	Name   string `form:"name"`
}

func ShowAllPlayersHandler(c *gin.Context) {

	var players []model.Playerdata
	initiator.POSTGRES.Find(&players)
	var result = make([]model.PlayerdataSerializer, len(players))
	for index, player := range players {
		result[index] = player.Serializer()
	}
	c.JSON(200, result)
}
