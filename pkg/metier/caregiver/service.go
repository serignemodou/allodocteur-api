package caregiver

import (
	"allodocteur-api/pkg/errors"
	"net/http"
)

type CaregiverService struct {
}

func (caregiverService CaregiverService) GetAllProfil() (*[]Profil, error) {
	profils, err := Persistance.GetAllProfil()
	if err != nil {
		return nil, err
	}
	return profils, nil
}

func (caregiverService CaregiverService) GetProfilByUID(uid string) (*Profil, error) {
	profil, err := Persistance.GetProfilByUID(uid)
	if err != nil {
		return nil, err
	}
	return profil, nil
}

func (caregiverService CaregiverService) GetProfilByName(name string) (*Profil, error) {
	profil, err := Persistance.GetProfilByName(name)
	if err != nil {
		return nil, err
	}
	return profil, nil
}

func (caregiverService CaregiverService) CreateProfil(profil *Profil) (*Profil, error) {
	profilExisted, err := caregiverService.GetProfilByUID(profil.ID)
	if err != nil {
		return nil, &errors.RequestError{
			StatusCode: http.StatusNotFound,
			ErrMessage: "caregiver profil not found",
		}
	}
	if profilExisted.Name == profil.Name {
		return nil, &errors.RequestError{
			StatusCode: http.StatusConflict,
			ErrMessage: "caregiver profil already exist",
		}
	}
	profilCreated, err1 := Persistance.CreateProfil(profil)
	if err1 != nil {
		println(err1.Error())
		return nil, err1
	}
	return profilCreated, nil
}

func (caregiverService CaregiverService) GetSpecialityByProfil(profilName string) (*[]Speciality, error) {
	_, err := caregiverService.GetProfilByName(profilName)
	if err != nil {
		return nil, &errors.RequestError{
			StatusCode: http.StatusNotFound,
			ErrMessage: "caregiver profil not found",
		}
	}
	specialities, err := Persistance.GetSpecialityByProfil(profilName)
	if err != nil {
		return nil, err
	}
	return specialities, nil
}

func (caregiverService CaregiverService) GetSpecialityByUID(uid string) (*Speciality, error) {
	speciality, err := Persistance.GetSpecialityByUID(uid)
	if err != nil {
		return nil, err
	}
	return speciality, nil
}

func (caregiverService CaregiverService) CreateSpeciality(speciality *Speciality) (*Speciality, error) {
	specialityExisted, err := caregiverService.GetSpecialityByUID(speciality.ID)
	if err != nil {
		return nil, &errors.RequestError{
			StatusCode: http.StatusNotFound,
			ErrMessage: "caregiver speciality not found",
		}
	}
	if specialityExisted.ID == speciality.ID {
		return nil, &errors.RequestError{
			StatusCode: http.StatusConflict,
			ErrMessage: "caregiver speciality already exist",
		}
	}
	_, errProfil := caregiverService.GetProfilByUID(speciality.Profil.ID)
	if errProfil != nil {
		return nil, &errors.RequestError{
			StatusCode: http.StatusNotFound,
			ErrMessage: "caregiver profil not found",
		}
	}
	specialityCreated, err := Persistance.CreateSpeciality(speciality)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return specialityCreated, nil
}
