package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

func Connect() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root@tcp(baseurl)/namadatabase?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err:= gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	  fmt.Println("database tidak dapat terhubung")
	}
	
   fmt.Println("Database terkoneksi")
  
   DB = conn
}
