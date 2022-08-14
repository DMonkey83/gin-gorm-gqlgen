package repository

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// CreateCategory ...
func (db *Database) CreateCategory(category entity.Category) {
	db.connection.Create(&category)
}

// UpdateCategory ...
func (db *Database) UpdateCategory(id string, category entity.Category) {
	var ctx *gin.Context
	var cat entity.Category //Category
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&cat).Error; err != nil {
		panic("Failed to Find Category")
	}

	// Validate input
	if err := ctx.ShouldBindJSON(&category); err != nil {
		panic("Validation Failed")
	}

	db.connection.Model(&cat).Save(&category)
}

// DeleteCategory ...
func (db *Database) DeleteCategory(id string) {
	var ctx *gin.Context
	var category entity.Category
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&category).Error; err != nil {
		panic("Failed to Find Category")
	}
	db.connection.Delete(category)
}

// FindCategories ...
func (db *Database) FindCategories() []entity.Category {
	var categories []entity.Category
	db.connection.Find(&categories)
	return categories
}

// FindCategory ...
func (db *Database) FindCategory(id string) entity.Category {
	var category entity.Category
	var ctx *gin.Context
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&category).Error; err != nil {
		panic("Failed to Find Category")
	}
	db.connection.First(category)
	return category
}
