package service

import "github.com/DMonkey83/go-gin-gorm/entity"

// CreateRPParty ...
func (service *BudgetServiceRepo) CreateRPParty(party entity.ReceivingPayingParty) entity.ReceivingPayingParty {
	service.budgetRepository.CreateRPParty(party)
	return party
}

// FindRPParties ...
func (service *BudgetServiceRepo) FindRPParties() []entity.ReceivingPayingParty {
	return service.budgetRepository.FindRPParties()
}

// FindRPParty ...
func (service *BudgetServiceRepo) FindRPParty(id string) entity.ReceivingPayingParty {
	return service.budgetRepository.FindRPParty(id)
}

// UpdateRPParty ...
func (service *BudgetServiceRepo) UpdateRPParty(id string, party entity.ReceivingPayingParty) {
	service.budgetRepository.UpdateRPParty(id, party)
}

// DeleteRPParty ...
func (service *BudgetServiceRepo) DeleteRPParty(id string) {
	service.budgetRepository.DeleteRPParty(id)
}
