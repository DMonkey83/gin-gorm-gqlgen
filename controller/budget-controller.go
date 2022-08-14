package controller

import (
	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/DMonkey83/go-gin-gorm/service"
	"github.com/gin-gonic/gin"
)

// BudgetController ...
type BudgetController interface {
	FindTransactions() []entity.Transaction
	FindTransaction(ctx *gin.Context) entity.Transaction
	CreateTransaction(ctx *gin.Context) error
	UpdateTransaction(ctx *gin.Context) error
	DeleteTransaction(ctx *gin.Context) error
	FindAccounts() []entity.Account
	FindAccount(ctx *gin.Context) entity.Account
	CreateAccount(ctx *gin.Context) error
	UpdateAccount(ctx *gin.Context) error
	DeleteAccount(ctx *gin.Context) error
	FindBanks() []entity.Bank
	FindBank(ctx *gin.Context) entity.Bank
	CreateBank(ctx *gin.Context) error
	UpdateBank(ctx *gin.Context) error
	DeleteBank(ctx *gin.Context) error
	FindRPParties() []entity.ReceivingPayingParty
	FindRPParty(ctx *gin.Context) entity.ReceivingPayingParty
	CreateRPParty(ctx *gin.Context) error
	UpdateRPParty(ctx *gin.Context) error
	DeleteRPParty(ctx *gin.Context) error
	FindCategories() []entity.Category
	FindCategory(ctx *gin.Context) entity.Category
	CreateCategory(ctx *gin.Context) error
	UpdateCategory(ctx *gin.Context) error
	DeleteCategory(ctx *gin.Context) error
	FindSubCategories() []entity.SubCategory
	FindSubCategory(ctx *gin.Context) entity.SubCategory
	CreateSubCategory(ctx *gin.Context) error
	UpdateSubCategory(ctx *gin.Context) error
	DeleteSubCategory(ctx *gin.Context) error
	FindUsers() []entity.User
	FindUser(ctx *gin.Context) entity.User
	CreateUser(ctx *gin.Context) error
	UpdateUser(ctx *gin.Context) error
	DeleteUser(ctx *gin.Context) error
}

// Controller ...
type Controller struct {
	service service.BudgetService
}

// New ...
func New(service service.BudgetService) BudgetController {
	return &Controller{
		service: service,
	}
}
