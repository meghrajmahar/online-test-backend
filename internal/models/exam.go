package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

const (
	QTypeMCQSingle = "MCQ_SINGLE"
	QTypeMCQMulti  = "MCQ_MULTI"
	QTypeTrueFalse = "TRUE_FALSE"
	QTypeShortText = "SHORT_TEXT"
	QTypeNumeric   = "NUMERIC"
)

type Exam struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string    `gorm:"not null"`
	Description *string
	DurationMin int  `gorm:"not null;check:duration_min > 0"`
	IsActive    bool `gorm:"not null;default:true"`
	Questions   []Question
	Timestamps
}

type Question struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ExamID         uuid.UUID `gorm:"type:uuid;index;not null"`
	Statement      string    `gorm:"not null"`
	QuestionType   string    `gorm:"not null"`
	Points         float64   `gorm:"not null;default:1"`
	NegativePoints float64   `gorm:"not null;default:0"`
	Position       int       `gorm:"not null;default:0"`
	Explanation    *string
	Metadata       datatypes.JSON
	Options        []Option `gorm:"constraint:OnDelete:CASCADE"`
	Timestamps
}

type Option struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	QuestionID uuid.UUID `gorm:"type:uuid;index;not null"`
	Text       string    `gorm:"not null"`
	IsCorrect  bool      `gorm:"not null;default:false"`
	Position   int       `gorm:"not null;default:0"`
	Timestamps
}

// type Attempt struct {
// 	ID          uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
// 	ExamID      uuid.UUID       `gorm:"type:uuid;index;not null"`
// 	UserID      uuid.UUID       `gorm:"type:uuid;index;not null"`
// 	StartedAt   time.Time       `gorm:"autoCreateTime"`
// 	SubmittedAt *time.Time
// 	TotalScore  *float64
// 	Status      string          `gorm:"not null;default:IN_PROGRESS"`
// 	Answers     []AttemptAnswer `gorm:"constraint:OnDelete:CASCADE"`
// }

// type AttemptAnswer struct {
// 	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
// 	AttemptID      uuid.UUID `gorm:"type:uuid;index;not null"`
// 	QuestionID     uuid.UUID `gorm:"type:uuid;index;not null"`
// 	AnswerText     *string
// 	NumericAnswer  *float64
// 	IsCorrect      *bool
// 	ScoreAwarded   *float64
// 	AnsweredAt     time.Time `gorm:"autoCreateTime"`
// 	SelectedOptions []Option `gorm:"many2many:attempt_answer_options;joinForeignKey:AttemptAnswerID;joinReferences:OptionID"`
// }

// type AttemptAnswerOption struct {
// 	AttemptAnswerID uuid.UUID `gorm:"type:uuid;primaryKey"`
// 	OptionID        uuid.UUID `gorm:"type:uuid;primaryKey"`
// }
