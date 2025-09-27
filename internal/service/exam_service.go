package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/meghraj/online-test-backend/internal/graph/model"
	"github.com/meghraj/online-test-backend/internal/models"
	"gorm.io/gorm"
)

type ExamService interface {
	// Admin
	CreateExam(ctx context.Context, input model.CreateExamInput) (*model.Exam, error)
	ExamList(ctx context.Context, limit *int, offset *int, search *string, sort *string) (*model.ExamList, error)
	// CreateQuestion(ctx context.Context, input model.CreateQuestionInput) (*model.QuestionAdmin, error)
	// AddOptions(ctx context.Context, input model.AddOptionsInput) ([]*model.OptionAdmin, error)
	// SetCorrectOptions(ctx context.Context, input model.SetCorrectOptionsInput) (*model.QuestionAdmin, error)

	// Fetch
	// GetExamPublic(ctx context.Context, examID string) (*model.Exam, error)
	// GetExamAdmin(ctx context.Context, examID string) ([]*model.QuestionAdmin, error)

	// Student
	// StartAttempt(ctx context.Context, input model.StartAttemptInput) (*model.Attempt, error)
	// AnswerMCQ(ctx context.Context, input model.AnswerMCQInput) (bool, error)
	// AnswerText(ctx context.Context, input model.AnswerTextInput) (bool, error)
	// AnswerNumeric(ctx context.Context, input model.AnswerNumericInput) (bool, error)
	// SubmitAttempt(ctx context.Context, input model.SubmitAttemptInput) (*model.AttemptResult, error)

}

type examService struct {
	DB *gorm.DB
}
type ListParams struct {
	Limit  int
	Offset int
	Search string
	Sort   string
}
type Page[T any] struct {
	Items []T   `json:items`
	Total int64 `json:total`
}

func defLimit(p *int, def, max int) int {
	if p == nil || *p <= 0 {
		return def
	}
	if *p > max {
		return max
	}
	return *p
}
func defOffset(p *int, def int) int {
	if p == nil || *p < 0 {
		return def
	}
	return *p
}
func defString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
func parseSort(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	parts := strings.Split(strings.ToLower(s), ".")
	if len(parts) != 2 {
		return "", errors.New("invalid sort; use e.g. title.asc or created_at.desc")
	}
	col, dir := parts[0], parts[1]
	switch col {
	case "created_at", "updated_at", "title":
	default:
		return "", fmt.Errorf("invalid sort column: %s", col)
	}
	switch dir {
	case "asc", "desc":
	default:
		return "", fmt.Errorf("invalid sort direction: %s", dir)
	}
	return col + " " + strings.ToUpper(dir), nil
}

func NewExamService(db *gorm.DB) ExamService { return &examService{DB: db} }

func toF(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}
func toI(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

func (s *examService) CreateExam(ctx context.Context, in model.CreateExamInput) (*model.Exam, error) {
	e := models.Exam{
		Title:       in.Title,
		DurationMin: in.DurationMin,
		IsActive:    true,
	}
	if in.Description != nil && *in.Description != "" {
		e.Description = in.Description
	}
	if err := s.DB.WithContext(ctx).Create(&e).Error; err != nil {
		return nil, err
	}
	return &model.Exam{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
		DurationMin: e.DurationMin,
	}, nil
}

func (s *examService) ExamList(ctx context.Context, limit *int, offset *int, search *string, sort *string) (*model.ExamList, error) {
	lpLimit := defLimit(limit, 20, 100)
	lpOffset := defOffset(offset, 0)
	lpSearch := strings.TrimSpace(defString(search))
	lpSort := strings.ToLower(defString(sort))

	// order is a string (NOT *string)
	order, err := parseSort(lpSort)
	if err != nil {
		return nil, err
	}
	if order == "" {
		order = "created_at DESC"
	}

	q := s.DB.WithContext(ctx).Model(&models.Exam{})

	if lpSearch != "" {
		like := "%" + lpSearch + "%"
		q = q.Where(`title ILIKE ? OR COALESCE(description,'') ILIKE ?`, like, like)
	}

	var total int64

	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}

	var rows []models.Exam
	if err := q.Order(order).Limit(lpLimit).Offset(lpOffset).Find(&rows).Error; err != nil {
		return nil, err
	}

	items := make([]*model.Exam, len(rows))

	for i := range rows {
		e := rows[i]
		items[i] = &model.Exam{
			ID:          e.ID.String(),
			Title:       e.Title,
			Description: e.Description,
			DurationMin: e.DurationMin,
		}
	}

	return &model.ExamList{Items: items, Total: int(total)}, nil
}
