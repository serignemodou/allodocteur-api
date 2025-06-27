package medecin

import (
	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"
)

type IMedecinRepo interface {
	getDB() *gorm.DB
	CreateSpeciliaty(specility *Speciality) (*Speciality, error)
	GetAllSpecialities() (*[]Speciality, error)
}

type IMedecinService interface {
	CreateSpeciliaty(speciality *Speciality) (*Speciality, error)
	GetAllSpecialities() (*[]Speciality, error)
}

type IMedecinEndpoint interface {
	//webServiceSpeciality(ws *restful.WebService)
	CreateSpeciliaty(request *restful.Request, response *restful.Response)
	GetAllSpecialities(request *restful.Request, response *restful.Response)
}
