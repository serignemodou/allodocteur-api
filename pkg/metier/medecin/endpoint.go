package medecin

import (
	"allodocteur-api/pkg/errors"
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
)

type MedecinEndpoint struct {
}

func (medecinEndpoint MedecinEndpoint) WebServiceMedecin(ws *restful.WebService) {
	tags := []string{"Medecin"}
	ws.Route(ws.POST("/medecin").
		To(medecinEndpoint.CreateSpeciliaty).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Create medecin speciality").
		Reads(Speciality{}).
		Returns(http.StatusCreated, "OK", Speciality{}).
		Returns(http.StatusBadRequest, "bad request", errors.RequestError{}))

	ws.Route(ws.GET("/medecin/speciality").
		To(medecinEndpoint.GetAllSpecialities).
		Doc("Get all medecin specialities").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(Speciality{}).
		Returns(http.StatusOK, "Ok", Speciality{}).
		Returns(http.StatusNotFound, "medecin speciality not found", errors.RequestError{}))
}

func (medecinEndpoint MedecinEndpoint) CreateSpeciliaty(request *restful.Request, response *restful.Response) {
	specilities := &Speciality{}
	err := request.ReadEntity(specilities)
	if err != nil {
		err = &errors.RequestError{
			StatusCode: http.StatusBadRequest,
			ErrMessage: err.Error(),
		}
		err.(*errors.RequestError).APIErrorHandler(request, response)
	}
	specilitiesCreated, err := Service.CreateSpeciliaty(specilities)
	if err != nil {
		err.(*errors.RequestError).APIErrorHandler(request, response)
	} else {
		err := response.WriteHeaderAndEntity(http.StatusCreated, specilitiesCreated)
		if err != nil {
			return
		}
	}
}

func (medecinEndpoint MedecinEndpoint) GetAllSpecialities(request *restful.Request, response *restful.Response) {
	specialities, err := Service.GetAllSpecialities()
	if err != nil {
		err.(*errors.RequestError).APIErrorHandler(request, response)
		return
	} else {
		err := response.WriteHeaderAndEntity(http.StatusOK, specialities)
		if err != nil {
			return
		}
	}
}
