package repository

type Citizens interface {
	CreateCitizen()
	FindCitizenById()
	FindAllCitizenPerPage()
	UpdateCitizen()
	DeleteCitizen()
}
