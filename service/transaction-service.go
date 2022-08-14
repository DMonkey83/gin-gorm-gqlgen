package service

import "github.com/DMonkey83/go-gin-gorm/entity"

// CreateTransaction ...
func (service *BudgetServiceRepo) CreateTransaction(tranaction entity.Transaction) entity.Transaction {
	service.budgetRepository.CreateTransaction(tranaction)
	return tranaction
}

// FindTransactions ...
func (service *BudgetServiceRepo) FindTransactions() []entity.Transaction {
	return service.budgetRepository.FindTransactions()
}

// FindTransaction ...
func (service *BudgetServiceRepo) FindTransaction(id string) entity.Transaction {
	return service.budgetRepository.FindTransaction(id)
}

// UpdateTransaction ...
func (service *BudgetServiceRepo) UpdateTransaction(id string, transaction entity.Transaction) {
	service.budgetRepository.UpdateTransaction(id, transaction)
}

// DeleteTransaction ...
func (service *BudgetServiceRepo) DeleteTransaction(id string) {
	service.budgetRepository.DeleteTransaction(id)
}
