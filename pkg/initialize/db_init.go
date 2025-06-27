package initialize

import (
	config "allodocteur-api/pkg/config"
	"allodocteur-api/pkg/metier/medecin"
	"log"
)

var DB = config.DbConnection()

func initDB() {
	DB.AutoMigrate(
		&medecin.Profil{},
		&medecin.Speciality{},
		&medecin.Caregiver{},
	)
	log.Println("Database schama created !")
}
