package service

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/DMonkey83/go-gin-gorm/repository"
)

// BudgetService ...
type BudgetService interface {
	CreateTransaction(transaction entity.Transaction) entity.Transaction
	FindTransactions() []entity.Transaction
	FindTransaction(id string) entity.Transaction
	UpdateTransaction(id string, transaction entity.Transaction)
	DeleteTransaction(id string)
	CreateAccount(account entity.Account) entity.Account
	FindAccounts() []entity.Account
	FindAccount(id string) entity.Account
	UpdateAccount(id string, account entity.Account)
	DeleteAccount(id string)
	CreateUser(user entity.User) entity.User
	FindUsers() []entity.User
	FindUser(id string) entity.User
	UpdateUser(id string, user entity.User)
	DeleteUser(id string)
	CreateBank(bank entity.Bank) entity.Bank
	FindBanks() []entity.Bank
	FindBank(id string) entity.Bank
	UpdateBank(id string, user entity.Bank)
	DeleteBank(id string)
	CreateRPParty(party entity.ReceivingPayingParty) entity.ReceivingPayingParty
	FindRPParties() []entity.ReceivingPayingParty
	FindRPParty(id string) entity.ReceivingPayingParty
	UpdateRPParty(id string, user entity.ReceivingPayingParty)
	DeleteRPParty(id string)
	CreateCategory(party entity.Category) entity.Category
	FindCategories() []entity.Category
	FindCategory(id string) entity.Category
	UpdateCategory(id string, user entity.Category)
	DeleteCategory(id string)
	CreateSubCategory(party entity.SubCategory) entity.SubCategory
	FindSubCategories() []entity.SubCategory
	FindSubCategory(id string) entity.SubCategory
	UpdateSubCategory(id string, user entity.SubCategory)
	DeleteSubCategory(id string)
}

// BudgetServiceRepo ...
type BudgetServiceRepo struct {
	budgetRepository repository.Repository
}

// NewService ...
func NewService(repo repository.Repository) BudgetService {
	return &BudgetServiceRepo{
		budgetRepository: repo,
	}
}
