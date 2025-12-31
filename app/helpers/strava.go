package helpers

import (
	"backend-dinakom/app/models"
	"backend-dinakom/configs"
	"backend-dinakom/external/types"
	"errors"
	"time"

	"gorm.io/gorm"
)

func GetValidStravaToken(db *gorm.DB, UserID string) (*string, error) {
	var token models.StravaToken
	if err := db.First(&token, "user_id = ?", UserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("token not found")
		}
		return nil, err
	}

	if time.Now().Unix() >= token.ExpiresAt {
		newToken, err := StravaRefreshTokenHandle(token.RefreshToken)
		if err != nil {
			return nil, errors.New("failed to refresh strava access_token")
		}

		if err := db.Model(&token).Updates(map[string]interface{}{
			"access_token":  newToken.AccessToken,
			"refresh_token": newToken.RefreshToken,
			"expires_at":    newToken.ExpiresAt,
		}).Error; err != nil {
			return nil, err
		}

	}

	return &token.AccessToken, nil
}

func StravaRefreshTokenHandle(refreshToken string) (*types.StravaTokenResponse, error) {
	client := configs.RestyClient
	data := map[string]string{
		"client_id":     configs.AppConfig.Strava.CLIENT_ID,
		"client_secret": configs.AppConfig.Strava.CLIENT_KEY,
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	var response *types.StravaTokenResponse
	_, err := client.R().
		SetFormData(data).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(&response).
		Post("https://www.strava.com/api/v3/oauth/token")
	if err != nil {
		return nil, err
	}
	return response, nil
}
