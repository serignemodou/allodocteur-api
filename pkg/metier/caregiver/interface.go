package caregiver

import (
	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"
)

type ICaregiverPersistance interface {
	getDB() *gorm.DB
	GetAllProfil() (*[]Profil, error)
	GetProfilByUID(uid string) (*Profil, error)
	GetProfilByName(name string) (*Profil, error)
	CreateProfil(profil *Profil) (*Profil, error)
}

type ICaregiverService interface {
	GetAllProfil() (*[]Profil, error)
	GetProfilByUID(uid string) (*Profil, error)
	GetProfilByName(name string) (*Profil, error)
	CreateProfil(profil *Profil) (*Profil, error)
}

type ICaregiverEndpoint interface {
	//webServiceSpeciality(ws *restful.WebService)
	GetAllProfil(request *restful.Request, response *restful.Response)
	GetProfilByUID(request *restful.Request, response *restful.Response)
	GetProfilByName(request *restful.Request, response *restful.Response)
	CreateProfil(request *restful.Request, response *restful.Response)
}
