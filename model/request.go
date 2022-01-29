package model

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
