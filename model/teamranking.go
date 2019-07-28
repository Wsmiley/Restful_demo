package model

type TeamRanking struct {
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
