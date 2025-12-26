package services

import "gorm.io/gorm"

type AIChatService interface {
}

type aIChatService struct {
	db *gorm.DB
}

func NewAIChatService(db *gorm.DB) AIChatService {
	return &aIChatService{db: db}
}
