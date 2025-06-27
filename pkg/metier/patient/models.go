package patient

type Patient struct {
	ID            string `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Adresse       string `json:"adresse"`
	Email         string `json:"email"`
	ContactNumber string `json:"contactNumber"`
	Birthday      string `json:"birthday"`
}
