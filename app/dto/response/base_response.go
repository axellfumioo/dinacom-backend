package response

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message" example:"data retrieved/created/updated/deleted successfully"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success    bool   `json:"success" example:"false"`
	Message    string `json:"message" example:"Internal Server Error"`
	StatusCode int    `json:"status_code"`
}

type PaginationResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}
