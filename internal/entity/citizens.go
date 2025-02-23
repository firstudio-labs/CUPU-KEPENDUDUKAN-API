package entity

import "time"

// FamilyStatus represents the family_status table
type FamilyStatus struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50"`
}

// Job represents the jobs table
type Job struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50"`
}

// Province represents the province table
type Province struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	ProvinceCode string `gorm:"unique;size:20"`
	Name         string `gorm:"size:60"`
	Coordinate   string `gorm:"size:100"`
}

// District represents the district table
type District struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	DistrictCode string `gorm:"unique;size:20"`
	Name         string `gorm:"size:60"`
	Coordinate   string `gorm:"size:100"`
}

// SubDistrict represents the sub_district table
type SubDistrict struct {
	ID              int    `gorm:"primaryKey;autoIncrement"`
	SubDistrictCode string `gorm:"unique;size:20"`
	Name            string `gorm:"size:60"`
	Coordinate      string `gorm:"size:100"`
}

// Village represents the village table
type Village struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	VillageCode string `gorm:"unique;size:20"`
	Name        string `gorm:"size:60"`
	Coordinate  string `gorm:"size:100"`
}

//type Coordinate struct {
//	Longitude string
//	Latitude  string
//}

// Citizen represents the Citizen table
type Citizen struct {
	ID                     int           `gorm:"primaryKey;autoIncrement"`
	NIK                    int           `gorm:"unique;not null;size:16"`
	KK                     int           `gorm:"not null;size:16"`
	FullName               string        `gorm:"not null;size:255"`
	Gender                 GenderOptions `gorm:"type:enum('Laki-Laki', 'Perempuan')"` // Gender enum tag
	BirthDate              time.Time     `gorm:"type:date;not null"`
	Age                    int           `gorm:"not null"`
	BirthPlace             string        `gorm:"not null"`
	Address                string
	ProvinceID             int
	DistrictID             int
	SubDistrictID          int
	VillageID              int
	RT                     string
	RW                     string
	PostalCode             int
	CitizenStatus          CitizenStatusOption `gorm:"type:enum('WNA', 'WNI')"` // CitizenStatus enum tag
	BirthCertificate       AvailableStatus     `gorm:"type:enum('Ada', 'Tidak Ada')"`
	BirthCertificateNo     int                 `gorm:"null;size:16"`
	BloodType              BloodType           `gorm:"type:enum('A', 'B', 'AB', 'O', 'A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-', 'Tidak Tahu')"`                                         // BloodType enum tag
	Religion               ReligionOption      `gorm:"type:enum('Islam', 'Kristen', 'Katholik', 'Hindu', 'Buddha', 'Kong Hu Cu', 'Lainya....')"`                                               // Religion enum tag
	MaritalStatus          MaritalStatusOption `gorm:"type:enum('Belum Kawin', 'Kawin Tercatat', 'Kawin Belum Tercatat', 'Cerai Hidup Tercatat', 'Cerai Hidup Belum Tercatat', 'Cerai Mati')"` // MaritalStatus enum tag
	MaritalCertificate     AvailableStatus     `gorm:"type:enum('Ada', 'Tidak Ada')"`                                                                                                          // nikah
	MaritalCertificateNo   int                 `gorm:"null"`                                                                                                                                   // nikah
	MarriageDate           time.Time
	DivorceCertificate     AvailableStatus `gorm:"type:enum('Ada', 'Tidak Ada')"` //cerai
	DivorceCertificateNo   int             `gorm:"not null"`                      //cerai
	DivorceCertificateDate time.Time       `gorm:"not null"`                      //cerai
	FamilyStatusID         int             `gorm:"not null"`                      // status dalam keluarga
	MentalDisorders        AvailableStatus `gorm:"type:enum('Ada', 'Tidak Ada')"`
	Disabilities           DisablitesStatus
	EducationStatus        EducationStatusOption `gorm:"type:enum('Tidak/Belum Sekolah', 'Belum tamat SD/Sederajat', 'Tamat SD', 'SLTP/SMP/Sederajat', 'SLTA/SMA/Sederajat', 'Diploma I/II', 'Akademi/Diploma III/ Sarjana Muda', 'Diploma IV/ Strata I/ Strata II', 'Strata III', 'Lainya...')"` // EducationStatus enum tag
	JobTypeID              int                   `gorm:"not null"`
	NIKMother              string                `gorm:"size:255"`
	Mother                 string                `gorm:"size:255"`
	NIKFather              string                `gorm:"size:255"`
	Father                 string                `gorm:"size:255"`
	Coordinate             string
	// Foreign Key Relations
	FamilyStatus FamilyStatus `gorm:"foreignKey:FamilyStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Job          Job          `gorm:"foreignKey:JobTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Province     Province     `gorm:"foreignKey:ProvinceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	District     District     `gorm:"foreignKey:DistrictID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SubDistrict  SubDistrict  `gorm:"foreignKey:SubDistrictID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Village      Village      `gorm:"foreignKey:VillageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type AvailableStatus string

const (
	StatusAvailable    AvailableStatus = "Ada"
	StatusNotAvailable AvailableStatus = "Tidak Ada"
)

type DisablitesStatus string

const (
	DisablitesFisik          DisablitesStatus = "Fisik"
	DisablitesBlind          DisablitesStatus = "Netra/Buta"
	DisablitesMute           DisablitesStatus = "Rungu/Wicara"
	DisablitesMentalAndSoul  DisablitesStatus = "Mental/Jiwa"
	DisablitesMentalAndFisik DisablitesStatus = "Fisik dan Mental"
	DisablitesOther          DisablitesStatus = "Lainnya"
)

type GenderOptions string

const (
	Man  GenderOptions = "Laki-Laki"
	Girl GenderOptions = "Perempuan"
)

type CitizenStatusOption string

const (
	WNA CitizenStatusOption = "WNA"
	WNI CitizenStatusOption = "WNI"
)

type BloodType string

const (
	A                BloodType = "A"
	B                BloodType = "B"
	AB               BloodType = "AB"
	O                BloodType = "O"
	APositive        BloodType = "A+"
	ANegative        BloodType = "A-"
	BPositive        BloodType = "B+"
	BNegative        BloodType = "B-"
	ABPositive       BloodType = "AB+"
	ABNegative       BloodType = "AB-"
	OPositive        BloodType = "O+"
	ONegative        BloodType = "O-"
	BloodTypeUnknown BloodType = "Tidak Tahu"
)

type ReligionOption string

const (
	Islam    ReligionOption = "Islam"
	Kristen  ReligionOption = "Kristen"
	Katholik ReligionOption = "Katholik"
	Hindu    ReligionOption = "Hindu"
	Buddha   ReligionOption = "Buddha"
	KongHuCu ReligionOption = "Kong Hu Cu"
	Etc      ReligionOption = "Lainya...."
)

type MaritalStatusOption string

const (
	Unmarried            MaritalStatusOption = "Belum Kawin"
	RegisteredMarriage   MaritalStatusOption = "Kawin Tercatat"
	UnregisteredMarriage MaritalStatusOption = "Kawin Belum Tercatat"
	RegisteredDivorce    MaritalStatusOption = "Cerai Hidup Tercatat"
	UnregisteredDivorce  MaritalStatusOption = "Cerai Hidup Belum Tercatat"
	Widowed              MaritalStatusOption = "Cerai Mati"
)

type EducationStatusOption string

const (
	NotInSchool                   EducationStatusOption = "Tidak/Belum Sekolah"
	NotFinishedElementary         EducationStatusOption = " Belum tamat SD/Sederajat"
	CompletedElementary           EducationStatusOption = "Tamat SD"
	JuniorHighSchool              EducationStatusOption = "SLTP/SMP/Sederajat"
	SeniorHighSchool              EducationStatusOption = "SLTA/SMA/Sederaja"
	DiplomaIorII                  EducationStatusOption = "Diploma I/II"
	AcademyOrDiplomaIII           EducationStatusOption = "Akademi/Diploma III/ Sarjana Muda"
	DiplomaIVOrBachelorsOrMasters EducationStatusOption = "Diploma IV/ Strata I/ Strata II"
	Doctorate                     EducationStatusOption = "Strata III"
	OtherEducation                EducationStatusOption = "Lainya..."
)
