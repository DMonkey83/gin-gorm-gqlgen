package controller

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/gin-gonic/gin"
)

// FindSubCategories ...
func (c *Controller) FindSubCategories() []entity.SubCategory {
	return c.service.FindSubCategories()
}

// FindSubCategory ...
func (c *Controller) FindSubCategory(ctx *gin.Context) entity.SubCategory {
	id := ctx.Param("id")
	return c.service.FindSubCategory(id)
}

// CreateSubCategory
func (c *Controller) CreateSubCategory(ctx *gin.Context) error {
	var sub entity.SubCategory
	err := ctx.ShouldBindJSON(&sub)
	if err != nil {
		return err
	}
	c.service.CreateSubCategory(sub)

	return nil
}

// UpdateSubCategory
func (c *Controller) UpdateSubCategory(ctx *gin.Context) error {
	var sub entity.SubCategory
	err := ctx.ShouldBindJSON(&sub)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	c.service.UpdateSubCategory(id, sub)

	return nil
}

// DeleteSubCategory
func (c *Controller) DeleteSubCategory(ctx *gin.Context) error {
	id := ctx.Param("id")
	c.service.DeleteSubCategory(id)

	return nil
}
