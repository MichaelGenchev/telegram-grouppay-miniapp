package storage

import (
	"context"

	"github.com/MichaelGenchev/telegram-grouppay-miniapp/internal/models"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByTelegramID(ctx context.Context, telegramID int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

// GroupRepository defines the interface for group data operations
type GroupRepository interface {
	CreateGroup(ctx context.Context, group *models.Group) error
	GetGroup(ctx context.Context, id int64) (*models.Group, error)
	GetUserGroups(ctx context.Context, userID int64) ([]models.Group, error)
	UpdateGroup(ctx context.Context, group *models.Group) error
	DeleteGroup(ctx context.Context, id int64) error
}

// ExpenseRepository defines the interface for expense data operations
type ExpenseRepository interface {
	CreateExpense(ctx context.Context, expense *models.Expense) error
	GetExpense(ctx context.Context, id int64) (*models.Expense, error)
	GetGroupExpenses(ctx context.Context, groupID int64) ([]models.Expense, error)
	UpdateExpense(ctx context.Context, expense *models.Expense) error
	DeleteExpense(ctx context.Context, id int64) error
}

// ParticipantRepository defines the interface for participant data operations
type ParticipantRepository interface {
	CreateParticipant(ctx context.Context, participant *models.Participant) error
	GetExpenseParticipants(ctx context.Context, expenseID int64) ([]models.Participant, error)
	UpdateParticipant(ctx context.Context, participant *models.Participant) error
	DeleteParticipant(ctx context.Context, id int64) error
}

// SettlementRepository defines the interface for settlement data operations
type SettlementRepository interface {
	CreateSettlement(ctx context.Context, settlement *models.Settlement) error
	GetSettlement(ctx context.Context, id int64) (*models.Settlement, error)
	GetGroupSettlements(ctx context.Context, groupID int64) ([]models.Settlement, error)
	UpdateSettlement(ctx context.Context, settlement *models.Settlement) error
}

// Storage aggregates all repository interfaces
type Storage interface {
	UserRepository
	GroupRepository
	ExpenseRepository
	ParticipantRepository
	SettlementRepository
}
