package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TeamRanking struct {
	gorm.Model
	Team           string `gorm:"unique;not null"`
	Win            string
	Defeat         string
	Winrate        string
	Differencewins string
	Home_field     string
	Visiting_field string
	Division       string
	East           string
	Point          string
	Lose_point     string
	Net_victory    string
	WL_streak      string
}

type TeamRankingSerializer struct {
	ID             uint       `json:"id"`
	CreateAt       time.Time  `json:"create_at"`
	UpdateAt       time.Time  `json:"update_at"`
	DeleteAt       *time.Time `json:"delete_at"`
	Team           string     `json:"team"`
	Win            string     `json:"win"`
	Defeat         string     `json:"defeat"`
	Winrate        string     `json:"winrate"`
	Differencewins string     `json:"diffencewins"`
	Home_field     string     `json:"home_field"`
	Visiting_field string     `json:"visiting_field"`
	Division       string     `json:"division"`
	East           string     `json:"east"`
	Point          string     `json:"point"`
	Lose_point     string     `json:"lose_point"`
	Net_victory    string     `json:"net_victory"`
	WL_streak      string     `json:"wl_streak"`
}

func (t *TeamRanking) Serializer() TeamRankingSerializer {
	return TeamRankingSerializer{
		ID:             t.ID,
		CreateAt:       t.CreatedAt,
		UpdateAt:       t.UpdatedAt,
		DeleteAt:       t.DeletedAt,
		Team:           t.Team,
		Win:            t.Win,
		Defeat:         t.Defeat,
		Winrate:        t.Winrate,
		Differencewins: t.Differencewins,
		Home_field:     t.Home_field,
		Visiting_field: t.Visiting_field,
		Division:       t.Division,
		East:           t.East,
		Point:          t.Point,
		Lose_point:     t.Lose_point,
		Net_victory:    t.Net_victory,
		WL_streak:      t.WL_streak,
	}
}

// type TeamRanking struct {
// 	gorm.Model
// 	Team           string `gorm:"type: varchar(64); not null; column:team"`
// 	Win            string `gorm:"type: varchar(64); not null; column:win"`
// 	Defeat         string `gorm:"type: varchar(64); not null; column:defeat"`
// 	Winrate        string `gorm:"type: varchar(64); not null; column:winrate"`
// 	Differencewins string `gorm:"type: varchar(64); not null; column:diffenrencewins"`
// 	Home_field     string `gorm:"type: varchar(64); not null; column:home_field"`
// 	Visiting_field string `gorm:"type: varchar(64); not null; column:visting_field"`
// 	Division       string `gorm:"type: varchar(64); not null; column:division"`
// 	East           string `gorm:"type: varchar(64); not null; column:east"`
// 	Point          string `gorm:"type: varchar(64); not null; column:point"`
// 	Lose_point     string `gorm:"type: varchar(64); not null; column:lose_point"`
// 	Net_victory    string `gorm:"type: varchar(64); not null; column:net_victory"`
// 	WL_streak      string `gorm:"type: varchar(64); not null; column:wl_streak"`
// }
