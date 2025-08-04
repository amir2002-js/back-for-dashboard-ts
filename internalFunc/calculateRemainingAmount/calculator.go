package calculateRemainingAmount

import (
	"gorm.io/gorm"
	"paysee2/layers/models"
)

func CalculatorAmount(db *gorm.DB, idCustomer int) (int, error) {
	var customer models.Customer
	err := db.Preload("Payment").First(&customer, idCustomer).Error
	if err != nil {
		return 0, err
	}
	totalPaid := 0
	for _, value := range customer.Payment {
		totalPaid += int(value.Amount)
	}

	remainingAmount := int(customer.Totality) - totalPaid

	return remainingAmount, nil
}
