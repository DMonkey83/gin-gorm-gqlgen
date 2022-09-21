// Package repository ...
package repository

import (
	"strconv"

	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// CreateAccount ...
func (db *Database) CreateAccount(account entity.Account) {
	db.connection.Create(&account)
}

// UpdateAccount ...
func (db *Database) UpdateAccount(id string, account entity.Account) {
	var ctx *gin.Context
	var acc entity.Account //Account
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&acc).Error; err != nil {
		panic("Failed to Find Account")
	}

	// Validate input
	if err := ctx.ShouldBindJSON(&account); err != nil {
		panic("Validation Failed")
	}

	db.connection.Model(&acc).Save(&account)
}

// DeleteAccount ...
func (db *Database) DeleteAccount(id string) {
	var ctx *gin.Context
	var account entity.Account
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&account).Error; err != nil {
		panic("Failed to Find Transaction")
	}
	db.connection.Delete(account)
}

// FindAccounts ...
func (db *Database) FindAccounts() []entity.Account {
	var accounts []entity.Account
	db.connection.Find(&accounts)
	return accounts
}

// FindAccount ...
func (db *Database) FindAccount(id string) entity.Account {
	var account entity.Account
	ID, error := strconv.ParseUint(id, 10, 32)
	if error != nil {
		panic(error)
	}
	println("id", ID)
	if err := db.connection.Where("id = ?", ID).First(&account).Error; err != nil {
		panic("Failed to Find Transaction")
	}
	db.connection.First(&account)
	return account
}
