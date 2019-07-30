package teams

import (
	initiator "NBA-master/init"
	"NBA-master/model"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrorTeam   = errors.New("not found record")
	ErrorTeamId = errors.New("id is not allow")
)

// TeamListHandler will list all the teams
// @Summaery List Team
// @Accept json
// @Tags teams
// @Security Bearer
// @Router /teams/ [get]
// @Produce  json
// @Success 200 {object} model.TeamRankingSerializer
func TeamListHandler(c *gin.Context) {
	var teams []model.TeamRanking
	initiator.POSTGRES.Find(&teams)
	var result = make([]model.TeamRankingSerializer, len(teams))
	for index, team := range teams {
		result[index] = team.Serializer()
	}
	c.JSON(200, result)
}

// TeamListHandler will list  the team
// @Summaery List Team
// @Accept json
// @Tags teams
// @Security Bearer
// @Router /teams/{id} [get]
// @Produce  json
// @Param teamID path string true "team id"
// @Success 200 {object} model.TeamRankingSerializer
func TeamHandler(c *gin.Context) {
	id := c.Param("teamID")

	number, _ := strconv.Atoi(id)

	if number <= 0 || number > 30 {
		c.AbortWithError(400, ErrorTeamId)
		return
	}

	var team model.TeamRanking
	if dbError := initiator.POSTGRES.Where("id = ?", id).First(&team).Error; dbError != nil {
		c.AbortWithError(400, ErrorTeam)
		return
	}

	c.JSON(200, team.Serializer())

}
