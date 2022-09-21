// Package repository ...
package repository

import (
	"strconv"

	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// CreateUser ...
func (db *Database) CreateUser(user entity.User) {
	db.connection.Create(&user)
}

// UpdateUser ...
func (db *Database) UpdateUser(id string, user entity.User) {
	var ctx *gin.Context
	var usr entity.User //user
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&usr).Error; err != nil {
		panic("Failed to Find User")
	}

	// Validate input
	if err := ctx.ShouldBindJSON(&user); err != nil {
		panic("Validation Failed")
	}

	db.connection.Model(&usr).Save(&user)
}

// DeleteUser ...
func (db *Database) DeleteUser(id string) {
	var ctx *gin.Context
	var user entity.User
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&user).Error; err != nil {
		panic("Failed to Find User")
	}
	db.connection.Delete(user)
}

// FindUsers ...
func (db *Database) FindUsers() []entity.User {
	var users []entity.User
	db.connection.Preload("Accounts").Find(&users)
	return users
}

// FindUser ...
func (db *Database) FindUser(id string) entity.User {
	var user entity.User
	ID, error := strconv.ParseUint(id, 10, 32)
	if error != nil {
		panic(error)
	}
	if err := db.connection.Where("id = ?", ID).First(&user).Error; err != nil {
		panic("Failed to Find User")
	}
	db.connection.Preload("Accounts").Preload("ReceivingPayingParties").First(&user)
	return user
}
