// Package controller ...
package controller

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// FindBanks ...
func (c *Controller) FindBanks() []entity.Bank {
	return c.service.FindBanks()
}

// FindBank ...
func (c *Controller) FindBank(ctx *gin.Context) entity.Bank {
	id := ctx.Param("id")
	return c.service.FindBank(id)
}

// CreateBank ...
func (c *Controller) CreateBank(ctx *gin.Context) error {
	var bank entity.Bank
	err := ctx.ShouldBindJSON(&bank)
	if err != nil {
		return err
	}
	c.service.CreateBank(bank)

	return nil
}

// UpdateBank ...
func (c *Controller) UpdateBank(ctx *gin.Context) error {
	var bank entity.Bank
	err := ctx.ShouldBindJSON(&bank)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	c.service.UpdateBank(id, bank)

	return nil
}

// DeleteBank ...
func (c *Controller) DeleteBank(ctx *gin.Context) error {
	id := ctx.Param("id")
	c.service.DeleteBank(id)

	return nil
}
