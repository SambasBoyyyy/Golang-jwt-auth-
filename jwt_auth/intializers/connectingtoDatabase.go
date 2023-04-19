package intializers

import (
	

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func GetDB() *gorm.DB {
	db, err := ConnectToMySQL()
	if err != nil {
		// handle error
	}
	
	if err != nil {
		// handle error
	}
	
	return db
}



func ConnectToMySQL() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}



