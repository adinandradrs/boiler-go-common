package base

import "database/sql"

type BaseEntity struct {
	IsDeleted   bool           `json:"is_deleted"`
	CreatedBy   sql.NullString `json:"created_by"`
	CreatedDate sql.NullTime   `json:"created_date"`
	UpdatedBy   sql.NullString `json:"updated_by"`
	UpdatedDate sql.NullTime   `json:"updated_date"`
}

type BaseService interface {
	Execute(input interface{}) (interface{}, error)
}

type DeleteRequest struct {
	Id           uint   `json:"id"`
	LoggedUserId uint   `json:"loggedUserId"`
	LoggedUser   string `json:"loggedUser"`
}

type FindByIdRequest struct {
	Id uint `json:"id"`
}

type SearchRequest struct {
	TextSearch string `json:"textSearch"`
	Start      uint   `json:"start"`
	Limit      uint   `json:"limit"`
	SortBy     string `json:"sortBy"`
	Sort       string `json:"sort"`
}

type SessionRequest struct {
	UserId       uint   `json:"userId" swaggerignore:"true"`
	Token        string `json:"token" swaggerignore:"true"`
	RefreshToken string `json:"refreshToken" swaggerignore:"true"`
	Username     string `json:"username" swaggerignore:"true"`
	ChannelId    string `json:"channelId" swaggerignore:"true"`
	ApiKey       string `json:"apiKey" swaggerignore:"true"`
}

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
