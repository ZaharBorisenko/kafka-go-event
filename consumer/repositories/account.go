package repositories

import (
	"gorm.io/gorm"
	"log"
)

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(account BankAccount) error
	Delete(id string) error
	FindAll() (accounts []BankAccount, err error)
	FindByID(id string) (account BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	if err := db.Table("Zakhar_Bank").AutoMigrate(&BankAccount{}); err != nil {
		log.Printf("Failed to auto-migrate: %v", err)
	}
	return accountRepository{db: db}
}

func (a accountRepository) Save(account BankAccount) error {
	return a.db.Table("Zakhar_Bank").Save(&account).Error
}

func (a accountRepository) Delete(id string) error {
	return a.db.Table("Zakhar_Bank").Where("id = ?", id).Delete(&BankAccount{}).Error
}

func (a accountRepository) FindAll() ([]BankAccount, error) {
	var accounts []BankAccount
	err := a.db.Table("Zakhar_Bank").Find(&accounts).Error
	return accounts, err

}

func (a accountRepository) FindByID(id string) (BankAccount, error) {
	var account BankAccount
	err := a.db.Table("Zakhar_Bank").Where("id = ?", id).First(&account).Error
	return account, err
}
