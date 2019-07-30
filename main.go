package main

import (
	"NBA-master/api_server"
	initiator "NBA-master/init"
	"NBA-master/model"
)

// @title NBA-data API Demo
// @version 1.0
// @description This is a server for NBA-data.
// @contact.name API Support
// @contact.email wsmiley29@gmail.com

// @host 127.0.0.1:5000
// @BasePath /v1/api
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
	// for i := 1; i <= 4; i++ {
	// 	playerURL := config.PlayerDataURL + strconv.Itoa(i)
	// 	docPlayer, _ := download.Downloder(playerURL)
	// 	domin.PlayerDataPhase(docPlayer)
	// }
	// teamerURL := config.TeamURL
	// docteamer, _ := download.Downloder(teamerURL)
	// domin.TeamPhase(docteamer)
}
