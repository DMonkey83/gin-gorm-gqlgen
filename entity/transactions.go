// Package entity ...
package entity

import (
	"time"

	"gorm.io/gorm"
)

// User struct  ...
type User struct {
	gorm.Model
	UserName               string                 `json:"username" binding:"required" gorm:"UNIQUE"`
	HashedPassword         string                 `json:"password" binding:"required" `
	FirstName              string                 `json:"first_name" gorm:"type:varchar(32)"`
	LastName               string                 `json:"last_name" binding:"required" gorm:"type:varchar(32)"`
	Age                    int8                   `json:"age" binding:"gte=1,lte=120"`
	Email                  string                 `json:"email" validate:"required, email" gorm:"type:varchar(256)"`
	Accounts               []Account              `json:"accounts" gorm:"foreignkey:UserID"`
	ReceivingPayingParties []ReceivingPayingParty `json:"receiving_paying_parties" gorm:"foreignkey:UserID"`
}

// Bank struct ...
type Bank struct {
	gorm.Model
	Title    string    `json:"title"`
	Accounts []Account `json:"accounts" gorm:"foreignkey:BankID"`
}

// Category struct ...
type Category struct {
	gorm.Model
	Title         string        `json:"title"`
	SubCategory   SubCategory   `json:"sub_category" gorm:"foreignkey:CategoryID"`
	SubCategories []SubCategory `json:"sub_categories" gorm:"foreignkey:CategoryID"`
	Transactions  []Transaction `json:"transactions" gorm:"foreignkey:CategoryID"`
}

// Account struct ...
type Account struct {
	gorm.Model
	Title        string        `json:"title"`
	AccNumber    string        `json:"acc_number"`
	SortCode     string        `json:"sort_code"`
	Transactions []Transaction `json:"transactions" gorm:"foreignkey:AccountID"`
	Balance      string        `json:"balance" binding:"required" gorm:"type:money"`
	UserID       uint          `json:"user_id"`
	BankID       uint          `json:"bank_id"`
}

// Transaction struct ...
type Transaction struct {
	gorm.Model
	Title      string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	TxType     TxType    `json:"tx_type" gorm:"type:integer"`
	Comment    string    `json:"comment" binding:"max=20" gorm:"type:varchar(100)"`
	CategoryID uint      `json:"category_id"`
	AccountID  uint      `json:"account_id"`
	ReceiverID uint      `json:"-"`
	Amount     string    `json:"amount" binding:"required" gorm:"type:money"`
	TxDate     time.Time `json:"tx_date"`
	Repeatable bool      `json:"repeatable"`
	EndDate    time.Time `json:"end_date"`
}

// SubCategory struct ...
type SubCategory struct {
	gorm.Model
	Title      string `json:"title"`
	CategoryID uint   `json:"-"`
}

// ReceivingPayingParty struct ...
type ReceivingPayingParty struct {
	gorm.Model
	Title        string        `json:"title"`
	Transactions []Transaction `json:"transactions" gorm:"foreignkey:ReceiverID"`
	UserID       uint          `json:"user_id"`
}

// TxType ...
type TxType int

// TxType enum
const (
	Expense TxType = iota + 1
	Income
	Transfer
)
