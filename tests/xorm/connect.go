package xorm

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

const (
	driverMySQL  = "mysql"
	driverPgSQL  = "postgres"
	driverSQLite = "sqlite"
)

var (
	dsn string
	db  *xorm.Engine
	err error
)

func connect(driver string) *xorm.Engine {
	if err = godotenv.Load("../.env"); err != nil {
		panic("`.env` file does not exist, please copy `.env.example` file to `.env` file")
	}

	switch driver {
	case driverMySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MySQL_DB_USERNAME"), os.Getenv("MySQL_DB_PASSWORD"), os.Getenv("MySQL_DB_HOST"), os.Getenv("MySQL_DB_PORT"), os.Getenv("MySQL_DB_DATABASE"))
		db, err = xorm.NewEngine("mysql", dsn)
	case driverPgSQL:
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("PgSQL_DB_HOST"), os.Getenv("PgSQL_DB_PORT"), os.Getenv("PgSQL_DB_USERNAME"), os.Getenv("PgSQL_DB_DATABASE"), os.Getenv("PgSQL_DB_PASSWORD"))
		db, err = xorm.NewEngine("postgres", dsn)
	case driverSQLite:
		db, err = xorm.NewEngine("sqlite3", os.Getenv("SQLite_DB_DATABASE"))
	}
	if err != nil {
		panic(fmt.Sprintf("failed to connect database, dsn: %q", dsn))
	}
	return db
}
