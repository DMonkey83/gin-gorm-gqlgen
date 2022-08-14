package service_test

import (
	"fmt"

	"github.com/DMonkey83/go-gin-gorm/entity"
	"github.com/DMonkey83/go-gin-gorm/repository"
	"github.com/DMonkey83/go-gin-gorm/service"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	title    = "Account Title"
	accName  = "89009889"
	sortCode = "90-90-90"
	balance  = "9876"
)

var testAccount entity.Account = entity.Account{
	Title:     title,
	AccNumber: accName,
	SortCode:  sortCode,
	Balance:   balance,
}

var _ = Describe("AccountService", func() {

	var (
		budgetRepository repository.Repository
		budgetService    service.BudgetService
	)

	BeforeEach(func() {
		budgetRepository = repository.NewRepository()
		budgetService = service.NewService(budgetRepository)
		createdAccount := budgetService.CreateAccount(testAccount)
		fmt.Println(createdAccount, "created")
	})

	Describe("Fetching all existing Accounts", func() {
		Context("If there is an account in the database", func() {
			It("should return at least one element", func() {
				accountList := budgetService.FindAccounts()
				fmt.Println("acc list", accountList[0])
				Expect(accountList).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				allAccounts := budgetService.FindAccounts()
				lastCreatedAccount := allAccounts[len(allAccounts)-1]
				fmt.Println("acc list", lastCreatedAccount.ID)

				Expect(lastCreatedAccount.Title).Should(Equal(title))
			})

		})
	})

})
