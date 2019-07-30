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
		panic("Database connection failed")
	}
	fmt.Println("Database connection success")
	POSTGRES = db
}
