package request

type CreateDoctorChatroomRequest struct {
	DoctorID string `json:"doctor_id" validate:"required"`
}
