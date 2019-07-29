package main

import (
	"NBA-master/api_server"
	"NBA-master/config"
	"NBA-master/domin"
	"NBA-master/download"
	initiator "NBA-master/init"
	"NBA-master/model"
	"strconv"
)

func main() {
	StartTable()
	defer initiator.POSTGRES.Close()

	PushDataIntoDB()

	api_server.New().Start()
}

func StartTable() {
	initiator.POSTGRES.AutoMigrate(&model.Playerdata{},
		&model.TeamRanking{})
}

func PushDataIntoDB() {
	for i := 1; i <= 4; i++ {
		playerURL := config.PlayerDataURL + strconv.Itoa(i)
		docPlayer, _ := download.Downloder(playerURL)
		domin.PlayerDataPhase(docPlayer)
	}
	teamerURL := config.TeamURL
	docteamer, _ := download.Downloder(teamerURL)
	domin.TeamPhase(docteamer)
}
