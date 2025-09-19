package services

import (
	"context"
	"fmt"

	"github.com/meghraj/online-test-backend/internal/auth"
	"github.com/meghraj/online-test-backend/internal/graph/model"
	"github.com/meghraj/online-test-backend/internal/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// type mutationResolver struct{ *Resolver}

// func (r *Resolver) Mutation() generated.MutationResolver {return &mutationResolver{r}}
type AuthService interface {
	Login(ctx context.Context, email, password string) (*model.AuthPayload, error)
	Register(ctx context.Context, email, name, password string) (*model.AuthPayload, error)
}

func (m *AuthServiceImpl) Register(ctx context.Context, email, name, password string) (*model.AuthPayload, error) {
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

func (m *AuthServiceImpl) Login(ctx context.Context, email, password string) (*model.AuthPayload, error) {
	if m == nil || m.DB == nil {
		return nil, fmt.Errorf("auth service DB not configured")
	}
	var u models.User
	if err := m.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, &gqlerror.Error{
			Message: "Email not exist.",
			Extensions: map[string]any{
				"code": "EMAIL_NOT_EXIST",
			},
		}
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return nil, &gqlerror.Error{
			Message: "Incorrect password.",
			Extensions: map[string]any{
				"code": "INVALID_PASSWORD",
			},
		}
	}

	token, err := auth.Sign(u.ID.String(), string(u.Role))

	if err != nil {
		return nil, &gqlerror.Error{
			Message: "Something went wrong",
			Extensions: map[string]any{
				"code": "SOMETHING_WRONG",
			},
		}
	}
	return &model.AuthPayload{
		AccessToken: token,
		User:        &model.User{ID: u.ID.String(), Email: u.Email, Name: u.Name, Role: model.Role(u.Role)},
	}, nil
}
