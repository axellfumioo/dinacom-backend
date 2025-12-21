package request

import "mime/multipart"

type ScanFoodRequest struct {
	Image multipart.FileHeader ``
}
