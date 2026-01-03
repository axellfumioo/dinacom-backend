package request

import "mime/multipart"

type CreateFamilyRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	memberIds    []string
	FamilyAvatar *multipart.FileHeader
}

type UpdateFamilyRequest struct {
	Name        *string
	Description *string
}
