package extservices

import (
	"backend-dinakom/configs"
	"fmt"
)

type FonnteResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func FetchFonnteAPI(userName string, phoneNumber string) error {
	restyClient := configs.RestyClient
	var message = fmt.Sprintf("%s, Jangan sampe kelupaan!! kamu belum scan makanan hari ini😉", userName)

	var result FonnteResponse
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetAuthToken("suKrp5PHdeGjNev54CCa").
		SetFormData(map[string]string{
			"target":  phoneNumber,
			"message": message,
		}).
		SetResult(&result).
		Post("https://api.fonnte.com/send")

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf(
			"failed send WA: status=%d body=%s",
			resp.StatusCode(),
			resp.String(),
		)
	}

	return nil
}
