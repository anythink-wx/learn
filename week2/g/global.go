package g

import (
	"gorm.io/gorm"
	"learn/week2/db"
)

var Db *gorm.DB

func Init() {
	Db = db.Init()
}
