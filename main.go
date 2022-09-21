// Package main ...
package main

import (
	"io"
	"net/http"
	"os"

	"github.com/DMonkey83/go-gin-gorm/controller"
	"github.com/DMonkey83/go-gin-gorm/middlewares"
	"github.com/DMonkey83/go-gin-gorm/repository"
	"github.com/DMonkey83/go-gin-gorm/service"
	"github.com/gin-gonic/gin"
)

var (
	budgetRepo repository.Repository = repository.NewRepository()

	budgetService service.BudgetService = service.NewService(budgetRepo)
	loginService  service.LoginService  = service.NewLoginService()
	jwtService    service.JWTService    = service.NewJWTService()

	budgetController controller.BudgetController = controller.New(budgetService)
	loginController  controller.LoginController  = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())
	server.SetTrustedProxies([]string{"192.168.1.255"})

	// LoginController ...
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// Basic Auth middleware applies top "/api"
	apiRoutes := server.Group("/api")
	{

		apiRoutes.GET("/transactions", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindTransactions())
		})

		apiRoutes.GET("/transactions/:id", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindTransaction(ctx))
		})

		apiRoutes.POST("/transactions", func(ctx *gin.Context) {
			err := budgetController.CreateTransaction(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "transaction created"})
			}
		})

		apiRoutes.PATCH("/transactions/:id", func(ctx *gin.Context) {
			err := budgetController.UpdateTransaction(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "transaction updated"})
			}
		})

		apiRoutes.DELETE("/transactions/:id", func(ctx *gin.Context) {
			err := budgetController.DeleteTransaction(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "transaction deleted"})
			}
		})
		apiRoutes.GET("/accounts", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindAccounts())
		})

		apiRoutes.GET("/accounts/:id", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindAccount(ctx))
		})

		apiRoutes.POST("/accounts", func(ctx *gin.Context) {
			err := budgetController.CreateAccount(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "account created"})
			}
		})

		apiRoutes.PUT("/accounts/:id", func(ctx *gin.Context) {
			err := budgetController.UpdateAccount(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "account updated"})
			}
		})

		apiRoutes.DELETE("/accounts/:id", func(ctx *gin.Context) {
			err := budgetController.DeleteAccount(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "account deleted"})
			}
		})
		apiRoutes.GET("/users", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindUsers())
		})

		apiRoutes.GET("/users/:id", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindUser(ctx))
		})

		apiRoutes.POST("/users", func(ctx *gin.Context) {
			err := budgetController.CreateUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
			}
		})

		apiRoutes.PUT("/users/:id", func(ctx *gin.Context) {
			err := budgetController.UpdateUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user updated"})
			}
		})

		apiRoutes.DELETE("/users/:id", func(ctx *gin.Context) {
			err := budgetController.DeleteUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user deleted"})
			}
		})
		apiRoutes.GET("/categories", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindCategories())
		})

		apiRoutes.GET("/categories/:id", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindCategory(ctx))
		})

		apiRoutes.POST("/categories", func(ctx *gin.Context) {
			err := budgetController.CreateCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "category created"})
			}
		})

		apiRoutes.PUT("/categories/:id", func(ctx *gin.Context) {
			err := budgetController.UpdateCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "category updated"})
			}
		})

		apiRoutes.DELETE("/category/:id", func(ctx *gin.Context) {
			err := budgetController.DeleteCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "category deleted"})
			}
		})
		apiRoutes.GET("/subcategories", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindSubCategories())
		})

		apiRoutes.GET("/subcategories/:id", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindSubCategory(ctx))
		})

		apiRoutes.POST("/subcategories", func(ctx *gin.Context) {
			err := budgetController.CreateSubCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "subcategory created"})
			}
		})

		apiRoutes.PUT("/subcategories/:id", func(ctx *gin.Context) {
			err := budgetController.UpdateSubCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "subcategory updated"})
			}
		})

		apiRoutes.DELETE("/subcategories/:id", func(ctx *gin.Context) {
			err := budgetController.DeleteSubCategory(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "subcategory deleted"})
			}
		})
		apiRoutes.GET("/rpparty", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindRPParties())
		})

		apiRoutes.GET("/rpparty/:id", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindRPParty(ctx))
		})

		apiRoutes.POST("/rpparty", func(ctx *gin.Context) {
			err := budgetController.CreateRPParty(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "party created"})
			}
		})

		apiRoutes.PUT("/rpparty/:id", func(ctx *gin.Context) {
			err := budgetController.UpdateRPParty(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "party updated"})
			}
		})

		apiRoutes.DELETE("/rpparty/:id", func(ctx *gin.Context) {
			err := budgetController.DeleteRPParty(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "party deleted"})
			}
		})
		apiRoutes.GET("/banks", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindBanks())
		})

		apiRoutes.GET("/banks/:id", func(ctx *gin.Context) {
			ctx.JSON(200, budgetController.FindBank(ctx))
		})

		apiRoutes.POST("/banks", func(ctx *gin.Context) {
			err := budgetController.CreateBank(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "bank created"})
			}
		})

		apiRoutes.PUT("/banks/:id", func(ctx *gin.Context) {
			err := budgetController.UpdateBank(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "bank updated"})
			}
		})

		apiRoutes.DELETE("/banks/:id", func(ctx *gin.Context) {
			err := budgetController.DeleteBank(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "bank deleted"})
			}
		})
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
