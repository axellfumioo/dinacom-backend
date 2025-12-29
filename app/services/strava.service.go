package services

import (
	"backend-dinakom/app/helpers"
	"backend-dinakom/app/types"
	"errors"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type StravaService interface {
	GetStravaProfile(UserID string) (*types.StravaProfileResponse, error)
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

	stravaResponse := &types.StravaProfileResponse{
		ID:            result.ID,
		Username:      result.Username,
		Firstname:     result.Firstname,
		Lastname:      result.Lastname,
		Profile:       result.Profile,
		ProfileMedium: result.ProfileMedium,
		State:         result.State,
		ResourceState: result.ResourceState,
		Sex:           result.Sex,
		CreatedAt:     result.CreatedAt,
	}
	return stravaResponse, nil
}
