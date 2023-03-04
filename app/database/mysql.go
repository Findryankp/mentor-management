package database

import (
	"clean-arch/app/configs"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(app configs.AppConfig) *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", app.DBUSERNAME, app.DBPASS, app.DBHOST, app.DBPORT, app.DBNAME)

	DB, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	InitialMigration(DB)

	return DB
}

func InitialMigration(db *gorm.DB) {
	err := db.AutoMigrate(
	// &data.User{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
