package services

import (
	"backend-dinakom/app/helpers"
	"errors"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type StravaService interface {
	GetStravaProfile(UserID string) (any, error)
}

type stravaService struct {
	db          *gorm.DB
	restyClient *resty.Client
}

func NewStravaService(db *gorm.DB, restyClient *resty.Client) StravaService {
	return &stravaService{db: db, restyClient: restyClient}
}

func (s *stravaService) GetStravaProfile(UserID string) (any, error) {
	stravaAccessToken, err := helpers.GetValidStravaToken(s.db, UserID)
	if err != nil {
		return nil, err
	}

	var result any
	if _, err := s.restyClient.R().SetAuthToken(*stravaAccessToken).SetResult(&result).Get("https://www.strava.com/api/v3/athlete"); err != nil {
		return nil, errors.New("failed to get strava profile")
	}

	return &result, nil
}
