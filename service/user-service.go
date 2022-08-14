package service

import "github.com/DMonkey83/go-gin-gorm/entity"

// CreateUser ...
func (service *BudgetServiceRepo) CreateUser(user entity.User) entity.User {
	service.budgetRepository.CreateUser(user)
	return user
}

// FindUsers ...
func (service *BudgetServiceRepo) FindUsers() []entity.User {
	return service.budgetRepository.FindUsers()
}

// FindUser ...
func (service *BudgetServiceRepo) FindUser(id string) entity.User {
	return service.budgetRepository.FindUser(id)
}

// UpdateUser ...
func (service *BudgetServiceRepo) UpdateUser(id string, user entity.User) {
	service.budgetRepository.UpdateUser(id, user)
}

// DeleteUser ...
func (service *BudgetServiceRepo) DeleteUser(id string) {
	service.budgetRepository.DeleteUser(id)
}
