package repository

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// CreateTransaction ...
func (db *Database) CreateTransaction(transaction entity.Transaction) {
	db.connection.Create(&transaction)
}

// UpdateTransaction ...
func (db *Database) UpdateTransaction(id string, transaction entity.Transaction) {
	var ctx *gin.Context
	var trx entity.Transaction //transactions
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&trx).Error; err != nil {
		panic("Failed to Find Transaction")
	}

	// Validate input
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		panic("Validation Failed")
	}

	db.connection.Model(&trx).Save(&transaction)
}

// DeleteTransaction ...
func (db *Database) DeleteTransaction(id string) {
	var ctx *gin.Context
	var transaction entity.Transaction
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&transaction).Error; err != nil {
		panic("Failed to Find Transaction")
	}
	db.connection.Delete(transaction)
}

// FindTransactions ...
func (db *Database) FindTransactions() []entity.Transaction {
	var transactions []entity.Transaction
	db.connection.Find(&transactions)
	return transactions
}

// FindTransaction ...
func (db *Database) FindTransaction(id string) entity.Transaction {
	var trx entity.Transaction
	var ctx *gin.Context
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&trx).Error; err != nil {
		panic("Failed to Find Transaction")
	}
	db.connection.First(trx)
	return trx
}
