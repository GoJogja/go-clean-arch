package addperson

type addPersonRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (ap *addPersonRequest) GetName() string {
	return ap.Name
}

func (ap *addPersonRequest) GetPhone() string {
	return ap.Phone
}
