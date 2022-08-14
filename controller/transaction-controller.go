package controller

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// FindTransactions ...
func (c *Controller) FindTransactions() []entity.Transaction {
	return c.service.FindTransactions()
}

// FindTransaction ...
func (c *Controller) FindTransaction(ctx *gin.Context) entity.Transaction {
	id := ctx.Param("id")
	return c.service.FindTransaction(id)
}

// CreateTransaction ...
func (c *Controller) CreateTransaction(ctx *gin.Context) error {
	var transaction entity.Transaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		return err
	}
	c.service.CreateTransaction(transaction)

	return nil
}

// UpdateTransaction ...
func (c *Controller) UpdateTransaction(ctx *gin.Context) error {
	var transaction entity.Transaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	c.service.UpdateTransaction(id, transaction)

	return nil
}

// DeleteTransaction ...
func (c *Controller) DeleteTransaction(ctx *gin.Context) error {
	id := ctx.Param("id")
	c.service.DeleteTransaction(id)

	return nil
}
