package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/db_kiki?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuka koneksi ke database MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
