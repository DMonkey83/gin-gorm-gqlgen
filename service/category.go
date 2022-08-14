package service

import "github.com/DMonkey83/go-gin-gorm/entity"

// CreateCategory ...
func (service *BudgetServiceRepo) CreateCategory(category entity.Category) entity.Category {
	service.budgetRepository.CreateCategory(category)
	return category
}

// FindCategories ...
func (service *BudgetServiceRepo) FindCategories() []entity.Category {
	return service.budgetRepository.FindCategories()
}

// FindCategory ...
func (service *BudgetServiceRepo) FindCategory(id string) entity.Category {
	return service.budgetRepository.FindCategory(id)
}

// UpdateCategory ...
func (service *BudgetServiceRepo) UpdateCategory(id string, category entity.Category) {
	service.budgetRepository.UpdateCategory(id, category)
}

// DeleteCategory ...
func (service *BudgetServiceRepo) DeleteCategory(id string) {
	service.budgetRepository.DeleteCategory(id)
}
