package request

import "mime/multipart"

type UploadAvatarRequest struct {
	UserID string                `json:"user_id"`
	Avatar *multipart.FileHeader `json:"avatar"`
}
