package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var BossDB *sql.DB
var BossGorm *gorm.DB

func InitBossDataDB(connectionString string) (*sql.DB, *gorm.DB) {
	fmt.Printf("connection string: %s\n", connectionString)
	newDB, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	if err = newDB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("The database is connected")
	BossDB = newDB

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: BossDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	BossGorm = gormDB
	return newDB, gormDB
}

func GetBossDB() *sql.DB {
	return BossDB
}
