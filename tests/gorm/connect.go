package gorm

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	driverMySQL  = "mysql"
	driverPgSQL  = "postgres"
	driverSQLite = "sqlite"
)

var (
	dsn string
	dia gorm.Dialector
	db  *gorm.DB
	err error
)

func connect(driver string) *gorm.DB {
	if err = godotenv.Load("../.env"); err != nil {
		panic("`.env` file does not exist, please copy `.env.example` file to `.env` file")
	}

	switch driver {
	case driverMySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MySQL_DB_USERNAME"), os.Getenv("MySQL_DB_PASSWORD"), os.Getenv("MySQL_DB_HOST"), os.Getenv("MySQL_DB_PORT"), os.Getenv("MySQL_DB_DATABASE"))
		dia = mysql.Open(dsn)
	case driverPgSQL:
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("PgSQL_DB_HOST"), os.Getenv("PgSQL_DB_PORT"), os.Getenv("PgSQL_DB_USERNAME"), os.Getenv("PgSQL_DB_DATABASE"), os.Getenv("PgSQL_DB_PASSWORD"))
		dia = postgres.Open(dsn)
	case driverSQLite:
		dia = sqlite.Open(os.Getenv("SQLite_DB_DATABASE"))
	}
	db, err = gorm.Open(dia, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database, dsn: %q", dsn))
	}
	return db
}
