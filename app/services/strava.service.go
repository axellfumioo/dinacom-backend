package services

import (
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type StravaService interface{}

type stravaService struct {
	db          *gorm.DB
	restyClient *resty.Client
}

func NewStravaService(db *gorm.DB, restyClient *resty.Client) StravaService {
	return &stravaService{db: db, restyClient: restyClient}
}
