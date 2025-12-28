package controllers

import "backend-dinakom/app/services"

type StravaController interface{}

type stravaController struct {
	stravaService services.StravaService
}

func NewStravaController(stravaService services.StravaService) StravaController {
	return &stravaController{stravaService: stravaService}
}
