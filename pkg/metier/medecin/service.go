package medecin

type MedecinService struct {
}

func (medecinService MedecinService) CreateSpeciliaty(specility *Speciality) (*Speciality, error) {
	return nil, nil
}

func (medecinService MedecinService) GetAllSpecialities() (*[]Speciality, error) {
	specialities, err := Repo.GetAllSpecialities()
	if err != nil {
		return nil, err
	}
	return specialities, nil
}
