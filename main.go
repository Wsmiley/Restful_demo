package main

import (
	"NBA-master/src/config"
	"NBA-master/src/domin"
	"NBA-master/src/download"
	initiator "NBA-master/src/init"
	"NBA-master/src/model"
	"fmt"
	"strconv"
)

func main() {
	StartTable()
	PushDataIntoDB()
	fmt.Println("end")
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
