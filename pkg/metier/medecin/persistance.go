package medecin

import "github.com/jinzhu/gorm"

type MedecinRepo struct {
	DB *gorm.DB
}

func (medecinRepo MedecinRepo) getDB() *gorm.DB {
	return medecinRepo.DB
}

func (medecinRepo MedecinRepo) GetAllSpecialities() (*[]Speciality, error) {
	return nil, nil
}

func (medecinRepo MedecinRepo) CreateSpeciliaty(speciality *Speciality) (*Speciality, error) {
	return nil, nil
}
