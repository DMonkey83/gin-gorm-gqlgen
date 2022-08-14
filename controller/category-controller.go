package controller

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// FindCategories ...
func (c *Controller) FindCategories() []entity.Category {
	return c.service.FindCategories()
}

// FindCategory ...
func (c *Controller) FindCategory(ctx *gin.Context) entity.Category {
	id := ctx.Param("id")
	return c.service.FindCategory(id)
}

// CreateCategory
func (c *Controller) CreateCategory(ctx *gin.Context) error {
	var category entity.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		return err
	}
	c.service.CreateCategory(category)

	return nil
}

// UpdateCategory
func (c *Controller) UpdateCategory(ctx *gin.Context) error {
	var category entity.Category
	err := ctx.ShouldBindJSON(&category)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	c.service.UpdateCategory(id, category)

	return nil
}

// DeleteCategory
func (c *Controller) DeleteCategory(ctx *gin.Context) error {
	id := ctx.Param("id")
	c.service.DeleteCategory(id)

	return nil
}
