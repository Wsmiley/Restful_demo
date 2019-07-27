package domin

import (
	initiator "NBA-master/src/init"
	"NBA-master/src/model"

	"github.com/PuerkitoBio/goquery"
)

func TeamPhase(doc *goquery.Document) error {
	doc.Find(".players_table tr.color_font1.bg_a~tr").Each(func(i int, s *goquery.Selection) {
		var Teamer model.TeamRanking
		Teamer.Team = s.Find(".left a").Text()
		Teamer.Win = s.Find("td").Eq(2).Text()
		Teamer.Defeat = s.Find("td").Eq(3).Text()
		Teamer.Winrate = s.Find("td").Eq(4).Text()
		Teamer.Differencewins = s.Find("td").Eq(5).Text()
		Teamer.Homae_field = s.Find("td").Eq(6).Text()
		Teamer.Visiting_field = s.Find("td").Eq(7).Text()
		Teamer.Division = s.Find("td").Eq(8).Text()
		Teamer.East = s.Find("td").Eq(9).Text()
		Teamer.Point = s.Find("td").Eq(10).Text()
		Teamer.Lose_point = s.Find("td").Eq(11).Text()
		Teamer.Net_victory = s.Find("td").Eq(12).Text()
		Teamer.WL_streak = s.Find("td").Eq(13).Text()

		initiator.POSTGRES.Create(&Teamer)
	})
	return nil
}
