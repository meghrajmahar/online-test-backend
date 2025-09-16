package resolvers

import (
	"context"
	"fmt"

	"github.com/meghraj/online-test-backend/internal/auth"
	"github.com/meghraj/online-test-backend/internal/graph/model"
	"github.com/meghraj/online-test-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// type mutationResolver struct{ *Resolver}

// func (r *Resolver) Mutation() generated.MutationResolver {return &mutationResolver{r}}

func (m *mutationResolver) Register(ctx context.Context, email, name, password string) (*model.AuthPayload, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u := models.User{Email: email, Name: name, Role: models.RoleAdmin, PasswordHash: string(hash)}
	if err := m.DB.Create(&u).Error; err != nil {
		return nil, err
	}
	token, err := auth.Sign(u.ID.String(), string(u.Role))
	if err != nil {
		return nil, err
	}
	return &model.AuthPayload{
		AccessToken: token,
		User:        &model.User{ID: u.ID.String(), Email: u.Email, Name: u.Name, Role: model.Role(u.Role)},
	}, nil
}

func (m *mutationResolver) Login(ctx context.Context, email, password string) (*model.AuthPayload, error) {
	var u models.User
	if err := m.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return nil, fmt.Errorf("Invalid credential")
	}

	token, err := auth.Sign(u.ID.String(), string(u.Role))

	if err != nil {
		return nil, err
	}
	return &model.AuthPayload{
		AccessToken: token,
		User:        &model.User{ID: u.ID.String(), Email: u.Email, Name: u.Name, Role: model.Role(u.Role)},
	}, nil
}
