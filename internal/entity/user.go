package entity

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	NIK         string `gorm:"uniqueIndex;not null"`
	FullName    string
	Province    string
	District    string
	SubDistrict string
	Village     string
	FullAddress string
	Coordinate  string
	Roles       Roles
	Username    string
	Password    string
	CreatedAt   int64  `gorm:"autoCreateTime"`
	UpdatedAt   *int64 `gorm:"autoUpdateTime"`
	DeletedAt   *int64 `gorm:"index"` //soft delete

	Complaints  []Complaint  `gorm:"foreignKey:UserID"` // Relasi One-to-Many dengan Complaint
	SubsPackets []SubsPacket `gorm:"foreignKey:UserID"` // Relasi One-to-Many dengan SubsPacket
}

type Roles string

const (
	MasterAdmin       Roles = "admin"
	MasterTechnician  Roles = "technician"
	MasterRegion      Roles = "region"
	MasterCitizenData Roles = "citizen-data"
)

type PacketInternet struct {
	Code         string `gorm:"primaryKey;unique"`
	ProviderName string
	Source       string
	Packet       string
	Duration     int64 `gorm:"not null"`
	Price        int
}

type Complaint struct {
	ID                 uint `gorm:"primaryKey;autoIncrement"`
	UserID             uint
	PacketInternetCode string
	Village            string
	ComplaintMessage   string `gorm:"text"`
	TechnicianNote     *string
	Reply              *string
	Status             StatusComplaint
	CreatedAt          int64  `gorm:"autoCreateTime"`
	UpdatedAt          *int64 `gorm:"autoUpdateTime"`
	DeletedAt          *int64 `gorm:"index"` //soft delete

	User           User           `gorm:"foreignKey:UserID;reference:ID"`
	PacketInternet PacketInternet `gorm:"foreignKey:PacketInternetCode;references:Code"` // Perbaiki nama field di sini

}

type StatusComplaint string

const (
	StatusRejected StatusComplaint = "rejected"
	StatusAccepted StatusComplaint = "accepted"
)

type SubsPacket struct {
	ID                 uint `gorm:"primaryKey;autoIncrement"`
	UserID             uint
	PacketInternetCode string
	Lifetime           int64
	PaymentTime        int64
	Status             PaymentStatus
	CreatedAt          int64  `gorm:"autoCreateTime"`
	UpdatedAt          *int64 `gorm:"autoUpdateTime"`
	DeletedAt          *int64 `gorm:"index"` //soft delete

	User           User           `gorm:"foreignKey:UserID;reference:ID"`
	PacketInternet PacketInternet `gorm:"foreignKey:PacketInternet Code;reference:Code"`
}

type PaymentStatus string

const (
	PaymentStatusPaid   PaymentStatus = "paid"
	PaymentStatusUnPaid PaymentStatus = "unpaid"
)
