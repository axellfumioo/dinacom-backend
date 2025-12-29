package services

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/types"
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type StravaService interface {
	GetStravaProfile(UserID string) (*types.StravaProfileResponse, error)
	GetStravaActivities(UserID string, page int, pageSize int) ([]types.StravaActivityResponse, error)
	GetStravaActivityByID(ID int, UserID string) (*types.StravaActivityResponse, error)
}

type stravaService struct {
	db          *gorm.DB
	restyClient *resty.Client
}

func NewStravaService(db *gorm.DB, restyClient *resty.Client) StravaService {
	return &stravaService{db: db, restyClient: restyClient}
}

func (s *stravaService) GetStravaProfile(UserID string) (*types.StravaProfileResponse, error) {
	stravaAccessToken, err := helpers.GetValidStravaToken(s.db, UserID)
	if err != nil {
		return nil, err
	}

	var result types.StravaProfileResponse
	if _, err := s.restyClient.R().SetAuthToken(*stravaAccessToken).SetResult(&result).Get("https://www.strava.com/api/v3/athlete"); err != nil {
		return nil, errors.New("failed to get strava profile")
	}

	return &result, nil
}

func (s *stravaService) GetStravaActivities(UserID string, page int, pageSize int) ([]types.StravaActivityResponse, error) {
	stravaAccessToken, err := helpers.GetValidStravaToken(s.db, UserID)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://www.strava.com/api/v3/athlete/activities?page=%d&per_page=%d", page, pageSize)
	var result []types.StravaActivityResponse
	if _, err := s.restyClient.R().SetAuthToken(*stravaAccessToken).SetResult(&result).Get(url); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *stravaService) GetStravaActivityByID(ID int, UserID string) (*types.StravaActivityResponse, error) {
	stravaAccessToken, err := helpers.GetValidStravaToken(s.db, UserID)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://www.strava.com/api/v3/activities/%d?include_all_efforts=false", ID)
	var result types.StravaActivityResponse
	if _, err := s.restyClient.R().SetAuthToken(*stravaAccessToken).SetResult(&result).Get(url); err != nil {
		return nil, err
	}

	return &result, nil
}
