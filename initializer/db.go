package initializer

import (
	"fmt"
	"log"
	"os"

	"github.com/fajardodiaz/infosgroup-employee-management/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectToDatabase() {

	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, database)

	var err error
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		DSN:        dsn,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Migrate the models
	Db.AutoMigrate(&models.Employee{})
	Db.AutoMigrate(&models.Gender{})
	Db.AutoMigrate(&models.Position{})
	Db.AutoMigrate(&models.Project{})
	Db.AutoMigrate(&models.State{})
	Db.AutoMigrate(&models.Team{})
}
