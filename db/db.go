package db

import (
	"fmt"

	"github.com/farmer-project/farmer/db/models"
	"github.com/fsouza/go-dockerclient"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const DB_TYPE = "mysql"

var (
	DbConnection gorm.DB
	DbServer     *docker.Container
)

func Connect() gorm.DB {
	if DbConnection.Error == nil {
		cs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			username,
			password,
			DbServer.NetworkSettings.IPAddress,
			"3306",
			db_name,
		)
		DbConnection, _ = gorm.Open(DB_TYPE, cs)
		DbConnection.LogMode(true)
	}

	return DbConnection
}

func Sync() {
	Connect()
	DbConnection.AutoMigrate(
		&models.Box{},
	)
}

func Close() error {
	if err := DbConnection.Close(); err != nil {
		return err
	}
	return nil
}
