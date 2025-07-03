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
	_, err := caregiverService.GetProfilByUID(profil.ID)
	if err != nil {
		return nil, &errors.RequestError{
			StatusCode: http.StatusBadRequest,
			ErrMessage: "error getting caregiver by uid",
		}
	}
	profilCreated, err1 := Persistance.CreateProfil(profil)
	if err1 != nil {
		println(err1.Error())
		return nil, err1
	}
	return profilCreated, nil
}
