package caregiver

import (
	"allodocteur-api/pkg/errors"
	"net/http"

	"github.com/jinzhu/gorm"
)

type CaregiverPersistance struct {
	DB *gorm.DB
}

func (caregiverPersistance CaregiverPersistance) getDB() *gorm.DB {
	return caregiverPersistance.DB
}

func (caregiverPersistance CaregiverPersistance) GetAllProfil() (*[]Profil, error) {
	var profil []Profil
	err := caregiverPersistance.DB.Preload("Specialities").Find(&profil).Error
	if err != nil {
		return nil, &errors.RequestError{
			StatusCode: http.StatusInternalServerError,
			ErrMessage: err.Error(),
		}
	}
	return &profil, err
}
func (caregiverPersistance CaregiverPersistance) GetProfilByUID(uid string) (*Profil, error) {
	var profil Profil
	err := caregiverPersistance.DB.Preload("Specialities").Where("id = ?", uid).Find(&profil).Error
	if err != nil {
		return nil, &errors.RequestError{
			StatusCode: http.StatusInternalServerError,
			ErrMessage: err.Error(),
		}
	}
	return &profil, nil
}
func (caregiverPersistance CaregiverPersistance) GetProfilByName(name string) (*Profil, error) {
	var profil Profil
	err := caregiverPersistance.DB.Preload("Specialities").Where("name = ?", name).Find(&profil).Error
	if err != nil {
		return nil, &errors.RequestError{
			StatusCode: http.StatusInternalServerError,
			ErrMessage: err.Error(),
		}
	}
	return &profil, nil
}
func (caregiverPersistance CaregiverPersistance) CreateProfil(profil *Profil) (*Profil, error) {
	err := caregiverPersistance.DB.Create(profil).Error
	if err != nil {
		return nil, &errors.RequestError{
			StatusCode: http.StatusInternalServerError,
			ErrMessage: err.Error(),
		}
	}
	return profil, nil
}
