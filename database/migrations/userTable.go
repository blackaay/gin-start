package migrations

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255"`
	Phone    string `gorm:"size:20;not null;unique"`
	Email    string `gorm:"size:100;not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func MigrateUserTable(db *gorm.DB) {
	db.Migrator().CreateTable(&User{})
}
