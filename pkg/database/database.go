package database

import (
	"boilerplate/pkg/configs"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	postgresv4 "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dsn := "root:paramadaksa@tcp(127.0.0.1:3306)/db41?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuka koneksi ke database MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectAndMigratePostgres() (rersult *gorm.DB, err error) {

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		configs.Database.User, configs.Database.Password, configs.Database.Host, configs.Database.Port, configs.Database.Name))
	if err != nil {
		return
	}

	// 2. Gunakan koneksi untuk golang-migrate
	driver, err := postgresv4.WithInstance(db, &postgresv4.Config{})
	if err != nil {
		return
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", configs.App.MigrationsPath), // Path ke file migrasi
		"postgres", driver)
	if err != nil {
		return
	}

	// Jalankan migrasi
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return
	}

	// 3. Gunakan koneksi yang sama dengan GORM
	return gorm.Open(postgres.New(postgres.Config{
		Conn: db, // Berikan koneksi `*sql.DB` yang sudah dibuka tadi ke GORM
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(configs.Database.PrintQuery)),
	})

}
