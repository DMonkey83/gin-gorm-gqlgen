package controller

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// FindAccounts ...
func (c *Controller) FindBanks() []entity.Bank {
	return c.service.FindBanks()
}

// FinBank ...
func (c *Controller) FindBank(ctx *gin.Context) entity.Bank {
	id := ctx.Param("id")
	return c.service.FindBank(id)
}

// CreateAccount ...
func (c *Controller) CreateBank(ctx *gin.Context) error {
	var bank entity.Bank
	err := ctx.ShouldBindJSON(&bank)
	if err != nil {
		return err
	}
	c.service.CreateBank(bank)

	return nil
}

// UpdateAccount ...
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

// DeleteAccount ...
func (c *Controller) DeleteBank(ctx *gin.Context) error {
	id := ctx.Param("id")
	c.service.DeleteBank(id)

	return nil
}
