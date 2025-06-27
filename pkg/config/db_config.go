package config

import (
	env "allodocteur-api/pkg/environment"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbHost     = env.DatabaseHost
	dbName     = env.DatabaseName
	dbPort     = env.DatabasePort
	dbType     = env.DatabaseType
	dbUser     = env.DatabaseUser
	dbPassword = env.DatabasePassword
)

func DbConnection() *gorm.DB {
	db, err := gorm.Open(dbType, dbUser+":"+dbPassword+"@("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
