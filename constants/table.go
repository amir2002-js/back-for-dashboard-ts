package constants

// kyc Status of User
type UserKycStatue int

const (
	PendingKYCStatus UserKycStatue = iota
	BannedKYCStatus
	VerifiedKYCStatus
	ApprovedKYCStatus
	RejectedKYCStatus
)

// Doc type of KYC
type KYCDocType int

const (
	DocumentType1 KYCDocType = iota
	DocumentType2
	DocumentType3
	DocumentType4
)

// statuses of KYC
type KYCStatuses int

const (
	KYCStatus1 KYCStatuses = iota
	KYCStatus2
	KYCStatus3
	KYCStatus4
)

//role users
type Role int
const (
	User Role = iota
	Admin
)