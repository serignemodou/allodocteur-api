package medecin

import (
	"github.com/jinzhu/gorm"
)

var (
	Repo     IMedecinRepo
	Endpoint IMedecinEndpoint
	Service  IMedecinService
)

func InitConfig(db *gorm.DB) {
	Repo = MedecinRepo{
		DB: db,
	}
	Service = MedecinService{}
	Endpoint = MedecinEndpoint{}
}
