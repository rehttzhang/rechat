package service

import (
	"rechat/internal/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

//QuestionService ...
type QuestionService struct {
	DB *gorm.DB
}

//CreateQuestion 新增问题
func (q *QuestionService) CreateQuestion(question models.Question) error {
	err := q.DB.Table("question").Create(&question).Error
	if err != nil {
		zap.L().Error("Failed to create question", zap.Any("question", question))
		return err
	}
	return nil
}
