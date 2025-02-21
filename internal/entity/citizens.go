package entity

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

// Citizen represents the Citizen table
type Citizen struct {
	ID                int    `gorm:"primaryKey;autoIncrement"`
	NIK               string `gorm:"unique;not null;size:16"`
	KK                string `gorm:"not null;size:16"`
	FullName          string `gorm:"not null;size:255"`
	BirthPlaceCode    string `gorm:"not null;size:20"`
	GenderID          int    `gorm:"not null"`
	ReligionID        int    `gorm:"not null"`
	StatusID          int    `gorm:"not null"`
	FamilyStatusID    int    `gorm:"not null"`
	EducationStatusID int    `gorm:"not null"`
	JobTypeID         int    `gorm:"not null"`
	Mother            string `gorm:"size:255"`
	Father            string `gorm:"size:255"`
	FamilyHead        string `gorm:"not null;size:255"`
	ProvinceCode      string `gorm:"not null;size:20"`
	DistrictCode      string `gorm:"not null;size:20"`
	SubDistrictCode   string `gorm:"not null;size:20"`
	VillageCode       string `gorm:"not null;size:20"`

	// Foreign Key Relations
	BirthPlace      Province        `gorm:"foreignKey:BirthPlaceCode;references:ProvinceCode;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Gender          Gender          `gorm:"foreignKey:GenderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Religion        Region          `gorm:"foreignKey:ReligionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Status          Status          `gorm:"foreignKey:StatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	FamilyStatus    FamilyStatus    `gorm:"foreignKey:FamilyStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	EducationStatus EducationStatus `gorm:"foreignKey:EducationStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Job             Job             `gorm:"foreignKey:JobTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Province        Province        `gorm:"foreignKey:ProvinceCode;references:ProvinceCode;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	District        District        `gorm:"foreignKey:DistrictCode;references:DistrictCode;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SubDistrict     SubDistrict     `gorm:"foreignKey:SubDistrictCode;references:SubDistrictCode;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Village         Village         `gorm:"foreignKey:VillageCode;references:VillageCode;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}
