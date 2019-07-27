package domin

import (
	initiator "NBA-master/src/init"
	"NBA-master/src/model"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func PlayerDataPhase(doc *goquery.Document) error {
	doc.Find(".players_table tr").Each(func(i int, s *goquery.Selection) {
		var Player model.Playerdata
		Player.Name = s.Find(".left a").Text()
		Player.Team = s.Find("td[width='50'] a").Text()
		Player.Source = s.Find(".bg_b").Text()
		Player.Shot = s.Find("td").Eq(4).Text()
		Player.Hitrate = s.Find("td").Eq(5).Text()
		Player.Threeshot = s.Find("td").Eq(6).Text()
		Player.Threerate = s.Find("td").Eq(7).Text()
		Player.Penaltyshot = s.Find("td").Eq(8).Text()
		Player.Penaltyrate = s.Find("td").Eq(9).Text()
		Player.Session = s.Find("td").Eq(10).Text()
		Player.Time = s.Find("td").Eq(11).Text()
		fmt.Printf("Review %d：--%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n", i, Player.Name, Player.Team, Player.Source, Player.Shot, Player.Hitrate, Player.Threeshot, Player.Threerate, Player.Penaltyshot, Player.Penaltyrate, Player.Session, Player.Time)
		initiator.POSTGRES.Create(&Player)
	})
	return nil
}
