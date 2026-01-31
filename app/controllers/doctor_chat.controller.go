package controllers

type DoctorChatController interface {
	
}

type doctorChatController struct {

}

func newDoctorChatController() DoctorChatController {
	return &doctorChatController{}
}

