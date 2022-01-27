package util

import (
	"fmt"

	_config "sirclo/graphql/config"
	_entities "sirclo/graphql/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlDriver(config *_config.AppConfig) *gorm.DB {

	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database: ", err)
		panic(err)
	}

	DatabaseMigration(db)

	return db
}

func DatabaseMigration(db *gorm.DB) {
	// db.Migrator().DropTable(entities.Book{})
	// db.Migrator().DropTable(entities.Person{})
	db.AutoMigrate(_entities.User{})
	db.AutoMigrate(_entities.Book{})

}
