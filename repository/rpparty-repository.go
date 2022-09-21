// Package repository ...
package repository

import (
	"strconv"

	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// CreateRPParty ...
func (db *Database) CreateRPParty(party entity.ReceivingPayingParty) {
	db.connection.Create(&party)
}

// UpdateRPParty ...
func (db *Database) UpdateRPParty(id string, party entity.ReceivingPayingParty) {
	var ctx *gin.Context
	var rpp entity.ReceivingPayingParty //RRParty
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&rpp).Error; err != nil {
		panic("Failed to Find the Party")
	}

	// Validate input
	if err := ctx.ShouldBindJSON(&party); err != nil {
		panic("Validation Failed")
	}

	db.connection.Model(&rpp).Save(&party)
}

// DeleteRPParty ...
func (db *Database) DeleteRPParty(id string) {
	var ctx *gin.Context
	var party entity.ReceivingPayingParty
	if err := db.connection.Where("id = ?", ctx.Param(id)).First(&party).Error; err != nil {
		panic("Failed to Find The Party")
	}
	db.connection.Delete(party)
}

// FindRPParties ...
func (db *Database) FindRPParties() []entity.ReceivingPayingParty {
	var parties []entity.ReceivingPayingParty
	db.connection.Find(&parties)
	return parties
}

// FindRPParty ...
func (db *Database) FindRPParty(id string) entity.ReceivingPayingParty {
	var party entity.ReceivingPayingParty
	ID, error := strconv.ParseUint(id, 10, 32)
	if error != nil {
		panic(error)
	}
	if err := db.connection.Where("id = ?", ID).First(&party).Error; err != nil {
		panic("Failed to Find Party")
	}
	db.connection.First(&party)
	return party
}
