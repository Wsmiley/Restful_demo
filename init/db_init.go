package initiator

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var POSTGRES *gorm.DB

func init() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable password=weili")
	if err != nil {
		panic("连接数据库失败")
	}
	fmt.Println("connect database")
	POSTGRES = db
}
