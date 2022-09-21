// Package controller ...
package controller

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// FindUsers ...
func (c *Controller) FindUsers() []entity.User {
	return c.service.FindUsers()
}

// FindUser ...
func (c *Controller) FindUser(ctx *gin.Context) entity.User {
	id := ctx.Param("id")
  println("id: ", id)
	return c.service.FindUser(id)
}

// CreateUser ...
func (c *Controller) CreateUser(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.CreateUser(user)

	return nil
}

// UpdateUser ...
func (c *Controller) UpdateUser(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	c.service.UpdateUser(id, user)

	return nil
}

// DeleteUser ...
func (c *Controller) DeleteUser(ctx *gin.Context) error {
	id := ctx.Param("id")
	c.service.DeleteUser(id)

	return nil
}
