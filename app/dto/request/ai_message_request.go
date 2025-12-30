package request

import "mime/multipart"

type CreateMessageRequest struct {
	Content string
}

type CreateMessageWithMediaRequest struct {
	Content string
	UserID  string
	ChatID  string
	Image   *multipart.FileHeader
}
