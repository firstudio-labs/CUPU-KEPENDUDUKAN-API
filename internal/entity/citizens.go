package entity

import "time"

// Gender represents the gender table
type Gender struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50"`
}

// Region represents the region table
type Region struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50"`
}

// Status represents the status table
type Status struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50"`
}

// FamilyStatus represents the family_status table
type FamilyStatus struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50"`
}

// EducationStatus represents the education_status table
type EducationStatus struct {
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

type CitizenStatus struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50"`
}

type BloodType struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50"`
}

// MaritalStatus SATATUS PERNIKAHAN
type MaritalStatus struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;size:50"`
}

// Citizen represents the Citizen table
type Citizen struct {
	ID                     int       `gorm:"primaryKey;autoIncrement"`
	NIK                    string    `gorm:"unique;not null;size:16"`
	KK                     string    `gorm:"not null;size:16"`
	FullName               string    `gorm:"not null;size:255"`
	GenderID               int       `gorm:"not null"`
	BirthDate              time.Time `gorm:"type:date;not null"`
	Age                    int       `gorm:"not null"`
	BirthPlaceID           string    `gorm:"not null"`
	Address                string
	ProvinceID             int `gorm:"not null;size:20"`
	DistrictID             int `gorm:"not null;size:20"`
	SubDistrictID          int `gorm:"not null;size:20"`
	VillageID              int `gorm:"not null;size:20"`
	RT                     int
	RW                     int
	PostalCode             int
	CitizenStatusID        int
	BirthCertificateNo     int `gorm:"null;size:16"`
	BloodTypeID            int
	ReligionID             int             `gorm:"not null"`
	MaritalStatusID        int             `gorm:"null"` // nikah
	MaritalCertificate     AvailableStatus `gorm:"null"` // nikah
	MaritalCertificateNo   int             `gorm:"null"` // nikah
	MarriageDate           time.Time
	DivorceCertificate     AvailableStatus `gorm:"not null"` //cerai
	DivorceCertificateNo   int             `gorm:"not null"` //cerai
	DivorceCertificateDate time.Time       `gorm:"not null"` //cerai
	FamilyStatusID         int             `gorm:"not null"`
	MentalDisorders        AvailableStatus
	Disabilities           DisablitesStatus
	StatusID               int    `gorm:"not null"`
	EducationStatusID      int    `gorm:"not null"`
	JobTypeID              int    `gorm:"not null"`
	NIKMother              string `gorm:"size:255"`
	Mother                 string `gorm:"size:255"`
	NIKFather              string `gorm:"size:255"`
	Father                 string `gorm:"size:255"`
	FamilyHead             string `gorm:"not null;size:255"`

	// Foreign Key Relations
	BirthPlace      Province        `gorm:"foreignKey:BirthPlaceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Gender          Gender          `gorm:"foreignKey:GenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Religion        Region          `gorm:"foreignKey:ReligionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Status          Status          `gorm:"foreignKey:StatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	FamilyStatus    FamilyStatus    `gorm:"foreignKey:FamilyStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	EducationStatus EducationStatus `gorm:"foreignKey:EducationStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Job             Job             `gorm:"foreignKey:JobTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Province        Province        `gorm:"foreignKey:ProvinceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	District        District        `gorm:"foreignKey:DistrictID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SubDistrict     SubDistrict     `gorm:"foreignKey:SubDistrictID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Village         Village         `gorm:"foreignKey:VillageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CitizenStatus   CitizenStatus   `gorm:"foreignKey:CitizenStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	BloodType       BloodType       `gorm:"foreignKey:BloodTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	MaritalStatus   MaritalStatus   `gorm:"foreignKey:MaritalStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type AvailableStatus string

const (
	StatusAvailable    AvailableStatus = "available"
	StatusNotAvailable AvailableStatus = "not_available"
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
