package my

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Migrate() {
	db, er := gorm.Open("sqlite3", "data.sqlite3")
	if er != nil {
		fmt.Print(er)
		return
	}

	defer db.Close()
	db.AutoMigrate(&User{}, &Group{}, &Post{}, &Comment{})
}
