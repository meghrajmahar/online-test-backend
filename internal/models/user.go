package models

import "github.com/google/uuid"

type Role string

const (
	RoleAdmin   Role = "ADMIN"
	RoleStudent Role = "STUDENT"
)

type User struct {
	ID           uuid.UUID `gorm:"type.uuid;default:gen_random_uuid();primaryKey"`
	Email        string    `gorm:"uniqueIndex;not null"`
	Name         string    `gorm:"not null"`
	Role         Role      `gorm:"type:text;not null"`
	PasswordHash string    `grom:"not null"`
	Timestamps
}
