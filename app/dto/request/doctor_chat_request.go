package request

type CreateDoctorChatroomRequest struct {
	DoctorID string `json:"doctor_id" validate:"required"`
}

type CreateDoctorChatMessageRequest struct {
	RoomID  string `json:"room_id" validate:"required"`
	Message string `json:"message" validate:"required"`
}
