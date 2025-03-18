package entity

type Relocation struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	KKRequest int64
	///
	ReasonMoving             string `gorm:"type:enum('PEKERJAAN','KEAMANAN','KESEHATAN','PENDIDIKAN','KELUARGA LAIN','KEBUTUHAN LAIN','PERUMAHAN')"`
	ClassificationRelocation string `gorm:"type:enum('DALAM DESA/KELURAHAN','DESA/KELURAHAN','ANTAR KECAMATAN','ANTAR KABUPATEN','ANTAR PROVINSI')"`
	MovingDate               string
	RelocationType           string   `gorm:"type:enum('KEPALA KELUARGA DAN SELURUH', 'KEPALA KELUARGA', 'KEPALA KELUARGA DAN SEBAGIA','ANGOTA KELUARGA')"`
	StatusKKMove             string   `gorm:"type:enum('BUAT KK BARU','NUMPANG KK','SEMUA KELUARGA PINDAH','NO KK TETAP')"`
	NIKFamilyMove            []*int64 `gorm:"type:json"`
	StatusKKStay             string   `gorm:"type:enum('BUAT KK BARU','NUMPANG KK','SEMUA KELUARGA PINDAH','NO KK TETAP','NO-KK DAN KEPALA KELUARGA')"`
	NewProvinceID            int      `gorm:"default:null"`
	NewDistrictID            int      `gorm:"default:null"`
	NewSubDistrictID         int      `gorm:"default:null"`
	NewVillageID             int      `gorm:"default:null"`
	NewRT                    string
	NewRW                    string
	///
	NewKK             *int64   `gorm:"default:null"`
	NewHeadOfFamily   *int64   `gorm:"default:null"`
	NIKFamilyStay     []*int64 `gorm:"type:json"`
	ProvinceIDStay    *int     `gorm:"default:null"`
	DistrictIDStay    *int     `gorm:"default:null"`
	SubDistrictIDStay *int     `gorm:"default:null"`
	VillageIDStay     *int     `gorm:"default:null"`
	CreatedAt         int64    `gorm:"column:created_at;type:bigint;default:0"`
	///
	VerificationStatus bool
	UpdatedAt          *int64 `gorm:"column:updated_at;type:bigint;default:0"`
	DeletedAt          *int64 `gorm:"column:deleted_at;type:bigint;default:0"`

	//ONE TO ONE
	Approved Approved `gorm:"foreignKey:RelocationID"`
}

type Approved struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	RelocationID uint   `gorm:"not null"`
	ApprovedBy   string `gorm:"not null"`
	ApproveDate  string `gorm:"not null"`
}

type RelocationType int

const (
	HeadOfFamilyAndAll     RelocationType = 1
	HeadOfFamilyOnly       RelocationType = 2
	HeadOfFamilyAndPartial RelocationType = 3
	FamilyMemberOnly       RelocationType = 4
)

func (r RelocationType) ToString() string {
	switch r {
	case HeadOfFamilyAndAll:
		return "KEPALA KELUARGA DAN SELURUH"
	case HeadOfFamilyOnly:
		return "KEPALA KELUARGA"
	case HeadOfFamilyAndPartial:
		return "KEPALA KELUARGA DAN SEBAGIAN"
	case FamilyMemberOnly:
		return "ANGOTA KELUARGA"
	default:
		return "NULL"
	}
}

//UNTUK STATUS NO KK PINDAH DAN YG TIDAK PINDAH

type StatusKK int

const (
	SharedKK StatusKK = iota + 1
	CreateNewKK
	FullFamily
	NoChangeKK
	NoKKHeadFamily
)

func (s StatusKK) ToString() string {
	switch s {
	case CreateNewKK:
		return "BUAT KK BARU"
	case SharedKK:
		return "NUMPANG KK"
	case FullFamily:
		return "SEMUA KELUARGA PINDAH"
	case NoChangeKK:
		return "NO KK TETAP"
	case NoKKHeadFamily:
		return "NO KK DAN KEPALA KELUARGA"
	default:
		return "NULL"
	}
}

type ReasonsMoving int

const (
	ReasonJob ReasonsMoving = iota + 1
	ReasonEducation
	ReasonSecurity
	ReasonHealth
	ReasonsHousing
	ReasonOtherFamily
	ReasonOther
)

func (r ReasonsMoving) ToString() string {
	switch r {
	case ReasonJob:
		return "PEKERJAAN"
	case ReasonEducation:
		return "PENDIDIKAN"
	case ReasonSecurity:
		return "KEAMANAN"
	case ReasonHealth:
		return "KESEHATAN"
	case ReasonsHousing:
		return "PERUMAHAN"
	case ReasonOtherFamily:
		return "KELUARGA LAIN"
	case ReasonOther:
		return "KEBUTUHAN LAIN"
	default:
		return "NULL"
	}
}

type ClassificationRelocation int

const (
	CRInternalVillage = iota + 1
	CRInternalSubDistrict
	CRInternalDistrict
	CRInternalProvince
)

func (c ClassificationRelocation) ToString() string {
	switch c {
	case CRInternalVillage:
		return "DALAM DESA/KELURAHAN"
	case CRInternalSubDistrict:
		return "DESA/KELURAHAN"
	case CRInternalDistrict:
		return "ANTAR KECAMATAN"
	case CRInternalProvince:
		return "ANTAR KABUPATEN"
	default:
		return "NULL"
	}
}
