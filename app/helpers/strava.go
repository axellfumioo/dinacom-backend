package helpers

import (
	"backend-dinakom/app/types/response"
	"backend-dinakom/configs"
	"net/url"
)

func StravaRefreshTokenHandle(refreshToken string) (*string, error) {
	client := configs.RestyClient
	data := url.Values{
		"client_id":     {configs.AppConfig.Strava.CLIENT_ID},
		"client_secret": {configs.AppConfig.Strava.CLIENT_KEY},
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	}
	var responseToken *response.StravaTokenResponse

	_, err := client.R().SetBody(&data).SetResult(&responseToken).Post("https://www.strava.com/api/v3/oauth/token")
	if err != nil {
		return nil, err
	}
	return &responseToken.AccessToken, nil
}