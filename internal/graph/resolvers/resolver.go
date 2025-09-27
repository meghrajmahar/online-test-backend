package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	services "github.com/meghraj/online-test-backend/internal/service"
	"gorm.io/gorm"
)

type Resolver struct {
	DB          *gorm.DB
	AuthService services.AuthService
	ExamService services.ExamService
}
