package caregiver

type Profil struct {
	ID           string       `json:"id" gorm:"primaryKey"`
	Name         string       `json:"name"`
	Specialities []Speciality `json:"specilities" gorm:"foreignKey:SpecialityID"`
}

type Speciality struct {
	ID     string `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"unique;not null"`
	Profil Profil `json:"profil_id" `
}

type Caregiver struct {
	ID            string      `json:"id" gorm:"primaryKey"`
	LastName      string      `json:"lastName" gorm:"not null"`
	FirstName     string      `json:"firstName"`
	Profil        Profil      `json:"profil" gorm:"foreignKey:ProfilID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Speciality    *Speciality `json:"speciality" gorm:"foreignKey:SpecialityID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Address       string      `json:"adresse"`
	ContactNumber string      `json:"contactNumber"`
	Email         string      `json:"email"`
	Description   string      `json:"description"`
}
