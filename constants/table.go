package constants

type Role int

const (
	User Role = iota
	Admin
)

type CustomerType int

const (
	Debtor   CustomerType = iota // مشتری بدهکار
	Creditor                     // مشتری طلب کار
)
