package model

type RestResponse struct {
	Message string      `json:"message"`
	Result  bool        `json:"result"`
	Data    interface{} `json:"data"`
}

type StatusCodeResponse struct {
	Code     int
	Response RestResponse
}

func GetStatusCodeResponse(code int, value RestResponse) *StatusCodeResponse {
	return &StatusCodeResponse{
		Code:     code,
		Response: value,
	}
}

type ValidationResponse struct {
	Result bool `json:"result"`
}

type PaginationResponse struct {
	Content       interface{} `json:"content"`
	Number        int         `json:"number"`
	Size          int         `json:"size"`
	TotalElements int         `json:"total_elements"`
	TotalPages    int         `json:"total_pages"`
	Sort          string      `json:"sort"`
	SortBy        string      `json:"sort_by"`
}
