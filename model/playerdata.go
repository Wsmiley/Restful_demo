package model

type Playerdata struct {
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
