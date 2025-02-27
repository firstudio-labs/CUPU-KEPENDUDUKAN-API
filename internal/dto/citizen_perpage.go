package dto

type CitizenResponse struct {
	Pagination `json:"pagination"`
	Citizens   []CitizensDTO `json:"citizens"`
}

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPage    int `json:"total_page"`
	TotalItems   int `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
	NextPage     int `json:"next_page"`
	PrevPage     int `json:"prev_page"`
}

type CitizensDTO struct {
	ID                     int
	NIK                    int64
	KK                     int64
	FullName               string
	Gender                 string
	BirthDate              string
	Age                    int
	BirthPlace             string
	Address                string
	ProvinceID             int
	DistrictID             int
	SubDistrictID          int
	VillageID              int
	RT                     string
	RW                     string
	PostalCode             int    `gorm:"null"`
	CitizenStatus          string `gorm:"type:enum('WNA', 'WNI')"`       // enum tag
	BirthCertificate       string `gorm:"type:enum('Ada', 'Tidak Ada')"` // enum tag
	BirthCertificateNo     string `gorm:"null"`
	BloodType              string `gorm:"type:enum('A', 'B', 'AB', 'O', 'A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-', 'Tidak Tahu')"`                                         // BloodType enum tag
	Religion               string `gorm:"type:enum('Islam', 'Kristen', 'Katholik', 'Hindu', 'Buddha', 'Kong Hu Cu', 'Lainya....')"`                                               // Religion enum tag
	MaritalStatus          string `gorm:"type:enum('Belum Kawin', 'Kawin Tercatat', 'Kawin Belum Tercatat', 'Cerai Hidup Tercatat', 'Cerai Hidup Belum Tercatat', 'Cerai Mati')"` // MaritalStatus enum tag
	MaritalCertificate     string `gorm:"type:enum('Ada', 'Tidak Ada');"`                                                                                                         // nikah
	MaritalCertificateNo   string `gorm:"null"`                                                                                                                                   // nikah
	MarriageDate           string
	DivorceCertificate     string `gorm:"type:enum('Ada', 'Tidak Ada')"` //cerai
	DivorceCertificateNo   string `gorm:"null"`                          //cerai
	DivorceCertificateDate string `gorm:"null"`                          //cerai
	FamilyStatus           string `gorm:"type:enum('KEPALA KELUARGA', 'ISTRI', 'ANAK', 'MERTUA', 'ORANG TUA', 'CUCU', 'FAMILI LAIN', 'LAINNYA');not null"`
	MentalDisorders        string `gorm:"type:enum('Ada', 'Tidak Ada');default:'Tidak Ada';"`
	Disabilities           string
	EducationStatus        string `gorm:"type:enum('Tidak/Belum Sekolah', 'Belum tamat SD/Sederajat', 'Tamat SD', 'SLTP/SMP/Sederajat', 'SLTA/SMA/Sederajat', 'Diploma I/II', 'Akademi/Diploma III/ Sarjana Muda', 'Diploma IV/ Strata I/ Strata II', 'Strata III', 'Lainya...')"` // EducationStatus enum tag
	JobTypeID              int    `gorm:"not null"`
	NIKMother              string `gorm:"size:255"`
	Mother                 string `gorm:"size:255"`
	NIKFather              string `gorm:"size:255"`
	Father                 string `gorm:"size:255"`
	Coordinate             string
}
