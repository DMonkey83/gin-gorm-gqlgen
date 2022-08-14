package repository

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Repository ...
type Repository interface {
	CreateTransaction(transaction entity.Transaction)
	UpdateTransaction(id string, transaction entity.Transaction)
	DeleteTransaction(id string)
	FindTransactions() []entity.Transaction
	FindTransaction(id string) entity.Transaction
	CreateAccount(account entity.Account)
	UpdateAccount(id string, account entity.Account)
	DeleteAccount(id string)
	FindAccounts() []entity.Account
	FindAccount(id string) entity.Account
	CreateUser(user entity.User)
	UpdateUser(id string, user entity.User)
	DeleteUser(id string)
	FindUsers() []entity.User
	FindUser(id string) entity.User
	CreateBank(bank entity.Bank)
	UpdateBank(id string, bank entity.Bank)
	DeleteBank(id string)
	FindBanks() []entity.Bank
	FindBank(id string) entity.Bank
	CreateRPParty(party entity.ReceivingPayingParty)
	UpdateRPParty(id string, user entity.ReceivingPayingParty)
	DeleteRPParty(id string)
	FindRPParties() []entity.ReceivingPayingParty
	FindRPParty(id string) entity.ReceivingPayingParty
	CreateCategory(category entity.Category)
	UpdateCategory(id string, user entity.Category)
	DeleteCategory(id string)
	FindCategories() []entity.Category
	FindCategory(id string) entity.Category
	CreateSubCategory(user entity.SubCategory)
	UpdateSubCategory(id string, user entity.SubCategory)
	DeleteSubCategory(id string)
	FindSubCategories() []entity.SubCategory
	FindSubCategory(id string) entity.SubCategory
}

// Database ...
type Database struct {
	connection *gorm.DB
}

// NewRepository ...
func NewRepository() Repository {
	dsn := "host=localhost user=evilnis password=Lon19ska83 dbname=go-gin-gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&entity.Transaction{}, &entity.User{}, &entity.Bank{}, &entity.Account{}, &entity.ReceivingPayingParty{}, &entity.Category{}, &entity.SubCategory{})
	return &Database{
		connection: db,
	}
}
