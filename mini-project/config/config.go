package config

import (
	"fmt"
	"mini-project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "admin",
		DB_Password: "miniproject",
		DB_Port:     "3306",
		DB_Host:     "database-mini-project.cfhwnsmwbh2n.ap-southeast-2.rds.amazonaws.com",
		DB_Name:     "mini_project",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	// DB, err = gorm.Open(mysql.Open("root:@tcp(host.docker.internal)/mini-project?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.Debtor{}, &model.Debt{})
	DB.Migrator().CreateConstraint(&model.Debtor{}, "Debts")
}
