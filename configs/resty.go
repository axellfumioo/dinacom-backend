package configs

import "github.com/go-resty/resty/v2"

var RestyClient *resty.Client

func InitRestyClient() *resty.Client {
	if RestyClient == nil {
		client := resty.New()
		RestyClient = client
	}

	return RestyClient
}
