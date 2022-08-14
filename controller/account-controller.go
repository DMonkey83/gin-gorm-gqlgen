package controller

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// FindAccounts ...
func (c *Controller) FindAccounts() []entity.Account {
	return c.service.FindAccounts()
}

// FindAccount ...
func (c *Controller) FindAccount(ctx *gin.Context) entity.Account {
	id := ctx.Param("id")
	return c.service.FindAccount(id)
}

// CreateAccount ...
func (c *Controller) CreateAccount(ctx *gin.Context) error {
	var account entity.Account
	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		return err
	}
	c.service.CreateAccount(account)

	return nil
}

// UpdateAccount ...
func (c *Controller) UpdateAccount(ctx *gin.Context) error {
	var account entity.Account
	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	c.service.UpdateAccount(id, account)

	return nil
}

// DeleteAccount ...
func (c *Controller) DeleteAccount(ctx *gin.Context) error {
	id := ctx.Param("id")
	c.service.DeleteAccount(id)

	return nil
}
