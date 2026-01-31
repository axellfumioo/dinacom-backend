package extservices

import (
	"backend-dinakom/configs"
	"fmt"
)

type FonnteResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func FetchFonnteAPI(userName, phoneNumber string) error {
	restyClient := configs.RestyClient
	message := fmt.Sprintf("%s, Jangan sampe kelupaan!! kamu belum scan makanan hari ini😉", userName)

	var result FonnteResponse

	resp, err := restyClient.R().
		SetHeader("Authorization", "suKrp5PHdeGjNev54CCa"). // ganti TOKEN dengan token asli lo
		SetMultipartFormData(map[string]string{
			"target":      phoneNumber,
			"message":     message,
			"delay":       "2",
			"countryCode": "62",
		}).
		SetResult(&result).
		Post("https://api.fonnte.com/send")

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("failed send WA: status=%d body=%s",
			resp.StatusCode(),
			resp.String(),
		)
	}

	fmt.Println("WA Response:", resp.String())
	return nil
}
