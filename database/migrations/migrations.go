package migrations

import (
	"github.com/blackaay/gin-start/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Migrate() {
	dsn := configs.GetDbDsn()
	log.Printf("dsn: %s", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
