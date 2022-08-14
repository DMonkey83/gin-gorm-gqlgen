package service

import "github.com/DMonkey83/go-gin-gorm/entity"

// CreateBank ...
func (service *BudgetServiceRepo) CreateBank(bank entity.Bank) entity.Bank {
	service.budgetRepository.CreateBank(bank)
	return bank
}

// FindBanks ...
func (service *BudgetServiceRepo) FindBanks() []entity.Bank {
	return service.budgetRepository.FindBanks()
}

// FindBank ...
func (service *BudgetServiceRepo) FindBank(id string) entity.Bank {
	return service.budgetRepository.FindBank(id)
}

// UpdateBank ...
func (service *BudgetServiceRepo) UpdateBank(id string, bank entity.Bank) {
	service.budgetRepository.UpdateBank(id, bank)
}

// DeleteBank ...
func (service *BudgetServiceRepo) DeleteBank(id string) {
	service.budgetRepository.DeleteBank(id)
}
