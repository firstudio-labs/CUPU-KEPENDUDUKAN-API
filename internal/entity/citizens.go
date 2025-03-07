package entity

import "time"

// Job represents the jobs table
type Job struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Code string `gorm:"unique" json:"code"`
	Name string `gorm:"not null" json:"name"`
} // IndonesiaProvince represents the indonesia_provinces table
type IndonesiaProvince struct {
	ID        uint64              `gorm:"column:id;primaryKey;autoIncrement"`
	Code      string              `gorm:"column:code;type:char(2);uniqueIndex"`
	Name      string              `gorm:"column:name;type:varchar(255)"`
	Meta      *string             `gorm:"column:meta;type:text"`
	CreatedAt *time.Time          `gorm:"column:created_at"`
	UpdatedAt *time.Time          `gorm:"column:updated_at"`
	Districts []IndonesiaDistrict `gorm:"foreignKey:ProvinceCode;references:Code"`
}

// IndonesiaDistrict represents the indonesia_districts table
type IndonesiaDistrict struct {
	ID           uint64                 `gorm:"column:id;primaryKey;autoIncrement"`
	Code         string                 `gorm:"column:code;type:char(4);uniqueIndex"`
	ProvinceCode string                 `gorm:"column:province_code;type:char(2)"`
	Name         string                 `gorm:"column:name;type:varchar(255)"`
	Meta         *string                `gorm:"column:meta;type:text"`
	CreatedAt    *time.Time             `gorm:"column:created_at"`
	UpdatedAt    *time.Time             `gorm:"column:updated_at"`
	Province     IndonesiaProvince      `gorm:"foreignKey:ProvinceCode;references:Code"`
	SubDistricts []IndonesiaSubDistrict `gorm:"foreignKey:DistrictCode;references:Code"`
}

// IndonesiaSubDistrict represents the indonesia_sub_districts table
type IndonesiaSubDistrict struct {
	ID           uint64             `gorm:"column:id;primaryKey;autoIncrement"`
	Code         string             `gorm:"column:code;type:char(7);uniqueIndex"`
	DistrictCode string             `gorm:"column:district_code;type:char(4)"`
	Name         string             `gorm:"column:name;type:varchar(255)"`
	Meta         *string            `gorm:"column:meta;type:text"`
	CreatedAt    *time.Time         `gorm:"column:created_at"`
	UpdatedAt    *time.Time         `gorm:"column:updated_at"`
	District     IndonesiaDistrict  `gorm:"foreignKey:DistrictCode;references:Code"`
	Villages     []IndonesiaVillage `gorm:"foreignKey:SubDistrictCode;references:Code"`
}

// IndonesiaVillage represents the indonesia_villages table
type IndonesiaVillage struct {
	ID              uint64               `gorm:"column:id;primaryKey;autoIncrement"`
	Code            string               `gorm:"column:code;type:char(10);uniqueIndex"`
	SubDistrictCode string               `gorm:"column:sub_district_code;type:char(7)"`
	Name            string               `gorm:"column:name;type:varchar(255)"`
	Meta            *string              `gorm:"column:meta;type:text"`
	CreatedAt       *time.Time           `gorm:"column:created_at"`
	UpdatedAt       *time.Time           `gorm:"column:updated_at"`
	SubDistrict     IndonesiaSubDistrict `gorm:"foreignKey:SubDistrictCode;references:Code"`
}

// TableName methods to specify custom table names
func (IndonesiaProvince) TableName() string {
	return "indonesia_provinces"
}

func (IndonesiaDistrict) TableName() string {
	return "indonesia_districts"
}

func (IndonesiaSubDistrict) TableName() string {
	return "indonesia_sub_districts"
}

func (IndonesiaVillage) TableName() string {
	return "indonesia_villages"
}

// Citizen represents the Citizen table
type Citizen struct {
	ID                     int    `gorm:"primaryKey;autoIncrement"`
	NIK                    int64  `gorm:"unique;not null"`
	KK                     int64  `gorm:"not null"`
	FullName               string `gorm:"not null;size:255"`
	Gender                 string `gorm:"type:enum('Laki-Laki', 'Perempuan')"` // Gender enum tag
	BirthDate              string `gorm:"not null"`
	Age                    int    `gorm:"not null"`
	BirthPlace             string `gorm:"not null"`
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
	//// Foreign Key Relations
	Job         Job                  `gorm:"foreignKey:JobTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Province    IndonesiaProvince    `gorm:"foreignKey:ProvinceID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	District    IndonesiaDistrict    `gorm:"foreignKey:DistrictID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	SubDistrict IndonesiaSubDistrict `gorm:"foreignKey:SubDistrictID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Village     IndonesiaVillage     `gorm:"foreignKey:VillageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

type FamilyStatus int

const (
	Children FamilyStatus = iota + 1
	HeadFamily
	Wife
	Parents
	ParentInLaw
	Grandchild
	OtherFamilyMember
)

func (fs FamilyStatus) ToString() string {
	switch fs {
	case Children:
		return "ANAK"

	case HeadFamily:
		return "KEPALA KELUARGA"

	case Wife:
		return "ISTRI"
	case Parents:
		return "ORANG TUA"
	case ParentInLaw:
		return "MERTUA"
	case Grandchild:
		return "CUCU"
	case OtherFamilyMember:
		return "FAMILI LAIN"

	default:
		return "LAINNYA"
	}
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
