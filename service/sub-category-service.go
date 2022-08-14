package service

import "github.com/DMonkey83/go-gin-gorm/entity"

// CreateSubCategory ...
func (service *BudgetServiceRepo) CreateSubCategory(subCategory entity.SubCategory) entity.SubCategory {
	service.budgetRepository.CreateSubCategory(subCategory)
	return subCategory
}

// FindSubCategories ...
func (service *BudgetServiceRepo) FindSubCategories() []entity.SubCategory {
	return service.budgetRepository.FindSubCategories()
}

// FindSubCategory ...
func (service *BudgetServiceRepo) FindSubCategory(id string) entity.SubCategory {
	return service.budgetRepository.FindSubCategory(id)
}

// UpdateSubCategory ...
func (service *BudgetServiceRepo) UpdateSubCategory(id string, subCategory entity.SubCategory) {
	service.budgetRepository.UpdateSubCategory(id, subCategory)
}

// DeleteSubCategory ...
func (service *BudgetServiceRepo) DeleteSubCategory(id string) {
	service.budgetRepository.DeleteSubCategory(id)
}
