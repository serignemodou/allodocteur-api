package caregiver

import (
	"github.com/jinzhu/gorm"
)

var (
	Persistance ICaregiverPersistance
	Endpoint    ICaregiverEndpoint
	Service     ICaregiverService
)

func InitConfig(db *gorm.DB) {
	Persistance = CaregiverPersistance{
		DB: db,
	}
	Service = CaregiverService{}
	Endpoint = CaregiverEndpoint{}
}
