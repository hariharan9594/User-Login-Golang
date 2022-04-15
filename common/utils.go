package common

import (
	"fmt"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"UserAuth/models"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	Db  *gorm.DB
	Log *logrus.Logger
)

type configuration struct {
	DbUser, DbPwd, Database, DbHost string
}

var Serviceconfig configuration

func initConfig() {
	Serviceconfig.DbUser = os.Getenv("DBUSER")
	Serviceconfig.DbPwd = os.Getenv("DBPWD")
	Serviceconfig.Database = "User"
	Serviceconfig.DbHost = "localhost"
}

func createDb() {
	if Db == nil {
		var err error
		db_config := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
			Serviceconfig.DbHost, Serviceconfig.DbUser, Serviceconfig.Database, Serviceconfig.DbPwd)
		Db, err = gorm.Open("postgres", db_config)
		if err != nil {
			Log.Panicf("Unable to connect to the %s database: %s", Serviceconfig.Database, err.Error())
		}
		Log.Debugf("Successfully connected to database '%s'", Serviceconfig.Database)
	}

	//Migrations without queries and using struct in models.go
	Db.AutoMigrate(&models.User{})

}

func CreateLog() {
	if Log == nil {
		Log = logrus.New()
		Log.SetLevel(logrus.DebugLevel)
		Log.SetFormatter(&nested.Formatter{
			HideKeys:    false,
			FieldsOrder: []string{"handler", "issue"},
			NoColors:    true,
		})
	}
}
