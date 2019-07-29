package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Playerdata struct {
	gorm.Model
	Name        string `gorm:"unique;NOT NULL"`
	Team        string `gorm:"NOT NULL"`
	Source      string `gorm:"NOT NULL"`
	Shot        string `gorm:"NOT NULL"`
	Hitrate     string `gorm:"NOT NULL"`
	Threeshot   string `gorm:"NOT NULL"`
	Threerate   string `gorm:"NOT NULL"`
	Penaltyshot string `gorm:"NOT NULL"`
	Penaltyrate string `gorm:"NOT NULL"`
	Session     string `gorm:"NOT NULL"`
	Time        string `gorm:"NOT NULL"`
}

type PlayerdataSerializer struct {
	ID          uint       `json:"id"`
	CreateAt    time.Time  `json:"create_at"`
	UpdateAt    time.Time  `json:"update_at"`
	DeleteAt    *time.Time `json:"delete_at"`
	Name        string     `json:"name"`
	Team        string     `json:"team"`
	Source      string     `json:"source"`
	Shot        string     `json:"shot"`
	Hitrate     string     `json:"hitrate"`
	Threeshot   string     `json:"threeshot"`
	Threerate   string     `json:"threerate"`
	Penaltyshot string     `json:"penaltyshot"`
	Penaltyrate string     `json:"penaltyrate"`
	Session     string     `json:"session"`
	Time        string     `json:"time"`
}

func (p *Playerdata) Serializer() PlayerdataSerializer {
	return PlayerdataSerializer{
		ID:          p.ID,
		CreateAt:    p.CreatedAt,
		UpdateAt:    p.UpdatedAt,
		DeleteAt:    p.DeletedAt,
		Name:        p.Name,
		Team:        p.Team,
		Source:      p.Source,
		Shot:        p.Shot,
		Hitrate:     p.Hitrate,
		Threeshot:   p.Threeshot,
		Threerate:   p.Threerate,
		Penaltyshot: p.Penaltyshot,
		Penaltyrate: p.Penaltyrate,
		Session:     p.Session,
		Time:        p.Time,
	}
}

// type Playerdata struct {
// 	gorm.Model
// 	Name        string `gorm:"type: varchar(64); not null; column:name"`
// 	Team        string `gorm:"type: varchar(64); not null; column:team"`
// 	Source      string `gorm:"type: varchar(64); not null; column:source"`
// 	Shot        string `gorm:"type: varchar(64); not null; column:shot"`
// 	Hitrate     string `gorm:"type: varchar(64); not null; column:hitrate"`
// 	Threeshot   string `gorm:"type: varchar(64); not null; column:threeshot"`
// 	Threerate   string `gorm:"type: varchar(64); not null; column:threerate"`
// 	Penaltyshot string `gorm:"type: varchar(64); not null; column:"penaltyshot`
// 	Penaltyrate string `gorm:"type: varchar(64); not null; column:penaltyrate"`
// 	Session     string `gorm:"type: varchar(64); not null; column:session"`
// 	Time        string `gorm:"type: varchar(64); not null; column:time"`
// }
