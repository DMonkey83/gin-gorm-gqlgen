package controller

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// FindRPParties
func (c *Controller) FindRPParties() []entity.ReceivingPayingParty {
	return c.service.FindRPParties()
}

// FindRPParty
func (c *Controller) FindRPParty(ctx *gin.Context) entity.ReceivingPayingParty {
	id := ctx.Param("id")
	return c.service.FindRPParty(id)
}

// CreateRPParty
func (c *Controller) CreateRPParty(ctx *gin.Context) error {
	var party entity.ReceivingPayingParty
	err := ctx.ShouldBindJSON(&party)
	if err != nil {
		return err
	}
	c.service.CreateRPParty(party)

	return nil
}

// UpdateRPParty
func (c *Controller) UpdateRPParty(ctx *gin.Context) error {
	var party entity.ReceivingPayingParty
	err := ctx.ShouldBindJSON(&party)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	c.service.UpdateRPParty(id, party)

	return nil
}

// DeleteRPParty
func (c *Controller) DeleteRPParty(ctx *gin.Context) error {
	id := ctx.Param("id")
	c.service.DeleteRPParty(id)

	return nil
}
