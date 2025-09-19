package services

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	DB        *gorm.DB
	DummyHash []byte
}

func NewAuthService(db *gorm.DB) AuthService {
	dh, _ := bcrypt.GenerateFromPassword([]byte("dummy-password"), bcrypt.DefaultCost)
	return &AuthServiceImpl{DB: db, DummyHash: dh}
}
