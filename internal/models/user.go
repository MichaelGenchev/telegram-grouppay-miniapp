package models

import "time"

// User represents a user in the system
type User struct {
	ID           int64     `json:"id" db:"id"`
	TelegramID   int64     `json:"telegram_id" db:"telegram_id"`
	Username     string    `json:"username" db:"username"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	LanguageCode string    `json:"language_code" db:"language_code"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// Group represents an expense group
type Group struct {
	ID          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedBy   int64     `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Expense represents a shared expense
type Expense struct {
	ID          int64     `json:"id" db:"id"`
	GroupID     int64     `json:"group_id" db:"group_id"`
	Description string    `json:"description" db:"description"`
	Amount      int64     `json:"amount" db:"amount"` // Amount in cents
	Currency    string    `json:"currency" db:"currency"`
	PaidBy      int64     `json:"paid_by" db:"paid_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Participant represents a user's participation in an expense
type Participant struct {
	ID        int64 `json:"id" db:"id"`
	ExpenseID int64 `json:"expense_id" db:"expense_id"`
	UserID    int64 `json:"user_id" db:"user_id"`
	Share     int64 `json:"share" db:"share"` // Share in cents
}

// Settlement represents a settlement between users
type Settlement struct {
	ID        int64     `json:"id" db:"id"`
	GroupID   int64     `json:"group_id" db:"group_id"`
	FromUser  int64     `json:"from_user" db:"from_user"`
	ToUser    int64     `json:"to_user" db:"to_user"`
	Amount    int64     `json:"amount" db:"amount"` // Amount in cents
	Currency  string    `json:"currency" db:"currency"`
	Status    string    `json:"status" db:"status"` // pending, completed, cancelled
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
