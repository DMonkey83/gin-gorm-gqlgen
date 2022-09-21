// Package repository ...
package repository

import (
	"strconv"

	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// CreateBank ...
func (db *Database) CreateBank(bank entity.Bank) {
	db.connection.Create(&bank)
}

// UpdateBank ...
func (db *Database) UpdateBank(id string, bank entity.Bank) {
	var ctx *gin.Context
	var bnk entity.Bank // Bank
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&bnk).Error; err != nil {
		panic("Failed to Find Bank")
	}

	// Validate input
	if err := ctx.ShouldBindJSON(&bank); err != nil {
		panic("Validation Failed")
	}

	db.connection.Model(&bnk).Save(&bank)
}

// DeleteBank ...
func (db *Database) DeleteBank(id string) {
	var ctx *gin.Context
	var bank entity.Bank
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&bank).Error; err != nil {
		panic("Failed to Find Bank")
	}
	db.connection.Delete(bank)
}

// FindBanks ...
func (db *Database) FindBanks() []entity.Bank {
	var bank []entity.Bank
	db.connection.Find(&bank)
	return bank
}

// FindBank ...
func (db *Database) FindBank(id string) entity.Bank {
	var bank entity.Bank
	ID, error := strconv.ParseUint(id, 10, 32)
	if error != nil {
		panic(error)
	}
	if err := db.connection.Where("id = ?", ID).First(&bank).Error; err != nil {
		panic("Failed to Find Bank")
	}
	db.connection.First(bank)
	return bank
}
