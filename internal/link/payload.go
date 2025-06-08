package link

type RequestLinkData struct {
	Url string `json:"url" validate:"required,url"`
}
