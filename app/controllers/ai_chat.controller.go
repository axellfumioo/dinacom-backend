package controllers

import "backend-dinakom/app/services"

type AIChatController interface{}

type aiChatController struct {
	AIChatService services.AIChatService
}

func NewAIChatController (AIChatService services.AIChatService) AIChatController {
	return &aiChatController{ AIChatService: AIChatService }
}