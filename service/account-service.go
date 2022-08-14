package service

import "github.com/DMonkey83/go-gin-gorm/entity"

// CreateAccount ...
func (service *BudgetServiceRepo) CreateAccount(account entity.Account) entity.Account {
	service.budgetRepository.CreateAccount(account)
	return account
}

// FindAccounts ...
func (service *BudgetServiceRepo) FindAccounts() []entity.Account {
	return service.budgetRepository.FindAccounts()
}

// FindAccount ...
func (service *BudgetServiceRepo) FindAccount(id string) entity.Account {
	return service.budgetRepository.FindAccount(id)
}

// UpdateAccount ...
func (service *BudgetServiceRepo) UpdateAccount(id string, account entity.Account) {
	service.budgetRepository.UpdateAccount(id, account)
}

// DeleteAccount ...
func (service *BudgetServiceRepo) DeleteAccount(id string) {
	service.budgetRepository.DeleteAccount(id)
}
