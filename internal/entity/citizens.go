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
	ID                     int       `gorm:"primaryKey;autoIncrement"`
	NIK                    int       `gorm:"unique;not null;size:16"`
	KK                     int       `gorm:"not null;size:16"`
	FullName               string    `gorm:"not null;size:255"`
	Gender                 string    `gorm:"type:enum('Laki-Laki', 'Perempuan')"` // Gender enum tag
	BirthDate              time.Time `gorm:"type:date;not null"`
	Age                    int       `gorm:"not null"`
	BirthPlace             string    `gorm:"not null"`
	Address                string
	ProvinceID             int
	DistrictID             int
	SubDistrictID          int
	VillageID              int
	RT                     string
	RW                     string
	PostalCode             int
	CitizenStatus          string `gorm:"type:enum('WNA', 'WNI')"`       // enum tag
	BirthCertificate       string `gorm:"type:enum('Ada', 'Tidak Ada')"` // enum tag
	BirthCertificateNo     int    `gorm:"null;size:16"`
	BloodType              string `gorm:"type:enum('A', 'B', 'AB', 'O', 'A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-', 'Tidak Tahu')"`                                         // BloodType enum tag
	Religion               string `gorm:"type:enum('Islam', 'Kristen', 'Katholik', 'Hindu', 'Buddha', 'Kong Hu Cu', 'Lainya....')"`                                               // Religion enum tag
	MaritalStatus          string `gorm:"type:enum('Belum Kawin', 'Kawin Tercatat', 'Kawin Belum Tercatat', 'Cerai Hidup Tercatat', 'Cerai Hidup Belum Tercatat', 'Cerai Mati')"` // MaritalStatus enum tag
	MaritalCertificate     string `gorm:"type:enum('Ada', 'Tidak Ada')"`                                                                                                          // nikah
	MaritalCertificateNo   int    `gorm:"null"`                                                                                                                                   // nikah
	MarriageDate           time.Time
	DivorceCertificate     string    `gorm:"type:enum('Ada', 'Tidak Ada')"` //cerai
	DivorceCertificateNo   int       `gorm:"not null"`                      //cerai
	DivorceCertificateDate time.Time `gorm:"not null"`                      //cerai
	FamilyStatusID         int       `gorm:"not null"`                      // status dalam keluarga
	MentalDisorders        string    `gorm:"type:enum('Ada', 'Tidak Ada')"`
	Disabilities           string
	EducationStatus        string `gorm:"type:enum('Tidak/Belum Sekolah', 'Belum tamat SD/Sederajat', 'Tamat SD', 'SLTP/SMP/Sederajat', 'SLTA/SMA/Sederajat', 'Diploma I/II', 'Akademi/Diploma III/ Sarjana Muda', 'Diploma IV/ Strata I/ Strata II', 'Strata III', 'Lainya...')"` // EducationStatus enum tag
	JobTypeID              int    `gorm:"not null"`
	NIKMother              string `gorm:"size:255"`
	Mother                 string `gorm:"size:255"`
	NIKFather              string `gorm:"size:255"`
	Father                 string `gorm:"size:255"`
	Coordinate             string
	// Foreign Key Relations
	FamilyStatus FamilyStatus `gorm:"foreignKey:FamilyStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Job          Job          `gorm:"foreignKey:JobTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Province     Province     `gorm:"foreignKey:ProvinceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	District     District     `gorm:"foreignKey:DistrictID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SubDistrict  SubDistrict  `gorm:"foreignKey:SubDistrictID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Village      Village      `gorm:"foreignKey:VillageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type AvailableStatus int

const (
	StatusAda AvailableStatus = iota + 1
	StatusTidakAda
)

// ToString method for AvailableStatus
func (s AvailableStatus) ToString() string {
	switch s {
	case StatusAda:
		return "Ada"
	case StatusTidakAda:
		return "Tidak Ada"
	default:
		return "Unknown"
	}
}

const (
	Man GenderOptions = iota + 1
	Girl
)

type GenderOptions int

// CitizenStatus options using iota
const (
	WNA CitizenStatusOption = iota + 1
	WNI
)

type CitizenStatusOption int

// BloodType options using iota
const (
	A BloodType = iota + 1
	B
	AB
	O
	APositive
	ANegative
	BPositive
	BNegative
	ABPositive
	ABNegative
	OPositive
	ONegative
	BloodTypeUnknown
)

type BloodType int

// Religion options using iota
const (
	Islam ReligionOption = iota + 1
	Kristen
	Katholik
	Hindu
	Buddha
	KongHuCu
	Etc
)

type ReligionOption int

// MaritalStatus options using iota
const (
	MaritalStatusUnmarried MaritalStatusOption = iota + 1
	MaritalStatusRegistered
	MaritalStatusUnregistered
	MaritalStatusRegisteredDivorce
	MaritalStatusUnregisteredDivorce
	MaritalStatusWidowed
)

type MaritalStatusOption int

// Disabilities options using iota
const (
	DisablitesFisik DisablitesStatus = iota + 1
	DisablitesBlind
	DisablitesMute
	DisablitesMentalAndSoul
	DisablitesMentalAndFisik
	DisablitesOther
)

type DisablitesStatus int

// EducationStatus options using iota
const (
	NotInSchool EducationStatusOption = iota + 1
	NotFinishedElementary
	CompletedElementary
	JuniorHighSchool
	SeniorHighSchool
	DiplomaIorII
	AcademyOrDiplomaIII
	DiplomaIVOrBachelorsOrMasters
	Doctorate
	OtherEducation
)

type EducationStatusOption int

// ToString methods for each type using iota

func (g GenderOptions) ToString() string {
	switch g {
	case Man:
		return "Laki-Laki"
	case Girl:
		return "Perempuan"
	default:
		return "Unknown"
	}
}

func (cs CitizenStatusOption) ToString() string {
	switch cs {
	case WNA:
		return "WNA"
	case WNI:
		return "WNI"
	default:
		return "Unknown"
	}
}

func (b BloodType) ToString() string {
	switch b {
	case A:
		return "A"
	case B:
		return "B"
	case AB:
		return "AB"
	case O:
		return "O"
	case APositive:
		return "A+"
	case ANegative:
		return "A-"
	case BPositive:
		return "B+"
	case BNegative:
		return "B-"
	case ABPositive:
		return "AB+"
	case ABNegative:
		return "AB-"
	case OPositive:
		return "O+"
	case ONegative:
		return "O-"
	case BloodTypeUnknown:
		return "Tidak Tahu"
	default:
		return "Unknown"
	}
}

func (r ReligionOption) ToString() string {
	switch r {
	case Islam:
		return "Islam"
	case Kristen:
		return "Kristen"
	case Katholik:
		return "Katholik"
	case Hindu:
		return "Hindu"
	case Buddha:
		return "Buddha"
	case KongHuCu:
		return "Kong Hu Cu"
	case Etc:
		return "Lainya..."
	default:
		return "Unknown"
	}
}

func (ms MaritalStatusOption) ToString() string {
	switch ms {
	case MaritalStatusUnmarried:
		return "Belum Kawin"
	case MaritalStatusRegistered:
		return "Kawin Tercatat"
	case MaritalStatusUnregistered:
		return "Kawin Belum Tercatat"
	case MaritalStatusRegisteredDivorce:
		return "Cerai Hidup Tercatat"
	case MaritalStatusUnregisteredDivorce:
		return "Cerai Hidup Belum Tercatat"
	case MaritalStatusWidowed:
		return "Cerai Mati"
	default:
		return "Unknown"
	}
}

func (d DisablitesStatus) ToString() string {
	switch d {
	case DisablitesFisik:
		return "Fisik"
	case DisablitesBlind:
		return "Netra/Buta"
	case DisablitesMute:
		return "Rungu/Wicara"
	case DisablitesMentalAndSoul:
		return "Mental/Jiwa"
	case DisablitesMentalAndFisik:
		return "Fisik dan Mental"
	case DisablitesOther:
		return "Lainnya"
	default:
		return "Unknown"
	}
}

func (e EducationStatusOption) ToString() string {
	switch e {
	case NotInSchool:
		return "Tidak/Belum Sekolah"
	case NotFinishedElementary:
		return "Belum tamat SD/Sederajat"
	case CompletedElementary:
		return "Tamat SD"
	case JuniorHighSchool:
		return "SLTP/SMP/Sederajat"
	case SeniorHighSchool:
		return "SLTA/SMA/Sederajat"
	case DiplomaIorII:
		return "Diploma I/II"
	case AcademyOrDiplomaIII:
		return "Akademi/Diploma III/ Sarjana Muda"
	case DiplomaIVOrBachelorsOrMasters:
		return "Diploma IV/ Strata I/ Strata II"
	case Doctorate:
		return "Strata III"
	case OtherEducation:
		return "Lainya..."
	default:
		return "Unknown"
	}
}
