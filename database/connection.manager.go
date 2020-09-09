package database

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	pq "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var InstanceDB *gorm.DB
var instanceDBMigrate *sql.DB

//InitDB start connection with DB
func InitDB() {
	dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		os.Getenv("ORDENES_DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("ORDENES_DB_USER"), os.Getenv("ORDENES_DB_NAME"), os.Getenv("ORDENES_DB_PASS"))
	fmt.Println("db connected ...")
	sqlDB, err := sql.Open("postgres", dsn)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	InstanceDB = gormDB
	instanceDBMigrate = sqlDB
}

func RunMigrations() {
	driver, err := pq.WithInstance(instanceDBMigrate, &pq.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err := m.Up(); err != nil {
		log.Println(err)
	} else {
		log.Println("Database migrated")
	}
	if err != nil {
		log.Fatalf("migration failed... %v", err)
	}
}
