package caregiver

import (
	"allodocteur-api/pkg/errors"
	"net/http"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
)

type CaregiverEndpoint struct {
}

func (caregiverEndpoint CaregiverEndpoint) WebServiceCaregiver(ws *restful.WebService) {
	tags := []string{"caregiver"}
	ws.Route(ws.GET("/caregiver/profil").
		To(caregiverEndpoint.GetAllProfil).
		Doc("Get all cavegiver profil").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(Profil{}).
		Returns(http.StatusOK, "Ok", Profil{}).
		Returns(http.StatusNotFound, "caregiver profil not found", errors.RequestError{}))

	ws.Route(ws.GET("/caregiver/profil/{uid-profil}").
		To(caregiverEndpoint.GetProfilByUID).
		Doc("Get caregiver profil by uid").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("uid-profil", "uid profil").DataType("string").Required(true)).
		Writes(Profil{}).
		Returns(http.StatusOK, "OK", Profil{}).
		Returns(http.StatusNotFound, "caregiver profil with uid not found", errors.RequestError{}))

	ws.Route(ws.GET("/caregiver/profil/{profil-name}").
		To(caregiverEndpoint.GetProfilByName).
		Doc("Get caregiver profil by name").
		Param(ws.PathParameter("profil-name", "profil name").DataType("string").Required(true)).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(Profil{}).
		Returns(http.StatusOK, "OK", Profil{}).
		Returns(http.StatusNotFound, "caregiver profil not found", errors.RequestError{}))

	ws.Route(ws.POST("/caregiver/profil").
		To(caregiverEndpoint.CreateProfil).
		Doc("Create caregiver profil").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(Profil{}).
		Returns(http.StatusOK, "OK", Profil{}).
		Returns(http.StatusBadRequest, "bad request error", errors.RequestError{}))

	ws.Route(ws.GET("/caregiver/speciality/{profil-name}").
		To(caregiverEndpoint.GetSpecialityByProfil).
		Doc("get caregiver speciality by profil name").
		Param(ws.PathParameter("profil-name", "caregiver profil name").DataType("string").Required(true)).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(Speciality{}).
		Returns(http.StatusOK, "OK", Speciality{}).
		Returns(http.StatusNotFound, "caregiver specialist not found", errors.RequestError{}))

	ws.Route(ws.GET("/caregiver/speciality/{speciality-uid}").
		To(caregiverEndpoint.GetSpecialityByUID).
		Doc("get caregiver speciality by uid").
		Param(ws.PathParameter("speciality-uid", "caregiver uid speciality").DataType("string").Required(true)).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(Speciality{}).
		Returns(http.StatusOK, "OK", Speciality{}).
		Returns(http.StatusNotFound, "caregiver speciality not found", errors.RequestError{}))

	ws.Route(ws.POST("/caregiver/speciality").
		To(caregiverEndpoint.CreateSpeciality).
		Doc("create caregiver speciality").
		Metadata(restfulspec.KeyOpenAPITags, tags).Reads(Speciality{}).
		Reads(Speciality{}).
		Returns(http.StatusOK, "OK", Speciality{}).
		Returns(http.StatusBadRequest, "bad request", errors.RequestError{}))
}

func (caregiverEndpoint CaregiverEndpoint) GetAllProfil(request *restful.Request, response *restful.Response) {
	profils, err := Service.GetAllProfil()
	if err != nil {
		err.(*errors.RequestError).APIErrorHandler(request, response)
	} else {
		err := response.WriteHeaderAndEntity(http.StatusOK, profils)
		if err != nil {
			return
		}
	}
}

func (caregiverEndpoint CaregiverEndpoint) GetProfilByUID(request *restful.Request, response *restful.Response) {
	uidProfil := request.PathParameter("uid-profil")
	profil, err := Service.GetProfilByUID(uidProfil)
	if err != nil {
		err.(*errors.RequestError).APIErrorHandler(request, response)
		return
	} else {
		err := response.WriteHeaderAndEntity(http.StatusOK, profil)
		if err != nil {
			return
		}
	}
}

func (caregiverEndpoint CaregiverEndpoint) GetProfilByName(request *restful.Request, response *restful.Response) {
	profilName := request.PathParameter("profil-name")
	profil, err := Service.GetProfilByName(profilName)
	if err != nil {
		err.(*errors.RequestError).APIErrorHandler(request, response)
		return
	} else {
		err := response.WriteHeaderAndEntity(http.StatusOK, profil)
		if err != nil {
			return
		}
	}
}

func (caregiverEndpoint CaregiverEndpoint) CreateProfil(request *restful.Request, response *restful.Response) {
	profil := &Profil{}
	err := request.ReadEntity(profil)
	if err != nil {
		err = &errors.RequestError{
			StatusCode: http.StatusBadRequest,
			ErrMessage: err.Error(),
		}
		err.(*errors.RequestError).APIErrorHandler(request, response)
	}
	profilCreated, err := Service.CreateProfil(profil)
	if err != nil {
		err.(*errors.RequestError).APIErrorHandler(request, response)
	} else {
		err := response.WriteHeaderAndEntity(http.StatusCreated, profilCreated)
		if err != nil {
			return
		}
	}
}

func (caregiverEndpoint CaregiverEndpoint) GetSpecialityByProfil(request *restful.Request, response *restful.Response) {
	profilName := request.PathParameter("profil-name")
	specialities, err := Service.GetSpecialityByProfil(profilName)
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

func (caregiverEndpoint CaregiverEndpoint) GetSpecialityByUID(request *restful.Request, response *restful.Response) {
	specialityUID := request.PathParameter("speciality-uid")
	speciality, err := Service.GetSpecialityByUID(specialityUID)
	if err != nil {
		err.(*errors.RequestError).APIErrorHandler(request, response)
		return
	} else {
		err := response.WriteHeaderAndEntity(http.StatusOK, speciality)
		if err != nil {
			return
		}
	}
}
func (caregiverEndpoint CaregiverEndpoint) CreateSpeciality(request *restful.Request, response *restful.Response) {
	speciality := Speciality{}
	err := request.ReadEntity(speciality)
	if err != nil {
		err = &errors.RequestError{
			StatusCode: http.StatusBadGateway,
			ErrMessage: err.Error(),
		}
		err.(*errors.RequestError).APIErrorHandler(request, response)
	}
	specialityCreated, err := Service.CreateSpeciality(speciality)
	if err != nil {
		err.(*errors.RequestError).APIErrorHandler(request, response)
	} else {
		err := response.WriteHeaderAndEntity(http.StatusOK, specialityCreated)
		if err != nil {
			return
		}
	}
}
