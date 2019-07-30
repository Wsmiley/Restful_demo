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

// TeamListHandler will list  the player
// @Summaery List player
// @Tags Players
// @Accept json
// @Router /players/{id} [get][get]
// @Produce  json
// @Param playerID path string true "player id"
// @Success 200 {object} model.PlayerdataSerializer
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

// TeamListHandler will list all the players
// @Summaery List players
// @Accept json
// @Tags Players
// @Router /players/ [get]
// @Produce  json
// @Success 200 {object} model.PlayerdataSerializer
func ShowAllPlayersHandler(c *gin.Context) {

	var players []model.Playerdata
	initiator.POSTGRES.Find(&players)
	var result = make([]model.PlayerdataSerializer, len(players))
	for index, player := range players {
		result[index] = player.Serializer()
	}
	c.JSON(200, result)
}
