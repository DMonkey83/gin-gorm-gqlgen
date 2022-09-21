// Package repository ...
package repository

import (
	"strconv"

	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// CreateSubCategory ...
func (db *Database) CreateSubCategory(subCategory entity.SubCategory) {
	db.connection.Create(&subCategory)
}

// UpdateSubCategory ...
func (db *Database) UpdateSubCategory(id string, subCategory entity.SubCategory) {
	var ctx *gin.Context
	var sub entity.SubCategory //Category
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&sub).Error; err != nil {
		panic("Failed to Find SubCategory")
	}

	// Validate input
	if err := ctx.ShouldBindJSON(&subCategory); err != nil {
		panic("Validation Failed")
	}

	db.connection.Model(&sub).Save(&subCategory)
}

// DeleteSubCategory ...
func (db *Database) DeleteSubCategory(id string) {
	var ctx *gin.Context
	var sub entity.SubCategory
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&sub).Error; err != nil {
		panic("Failed to Find Category")
	}
	db.connection.Delete(sub)
}

// FindSubCategories ...
func (db *Database) FindSubCategories() []entity.SubCategory {
	var subs []entity.SubCategory
	db.connection.Find(&subs)
	return subs
}

// FindSubCategory ...
func (db *Database) FindSubCategory(id string) entity.SubCategory {
	var sub entity.SubCategory
	ID, error := strconv.ParseUint(id, 10, 32)
	if error != nil {
		panic(error)
	}
	if err := db.connection.Where("id = ?", ID).First(&sub).Error; err != nil {
		panic("Failed to Find SubCategory")
	}
	db.connection.First(&sub)
	return sub
}
