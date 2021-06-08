package transport

type CreateUrlRequest struct {
	CustomAlias string `json:"customAlias"`
	OriginalUrl string `json:"originalUrl"`
}

func NewCreateUrlRequest(customAlias string, originalUrl string) *CreateUrlRequest {
	c := new(CreateUrlRequest)
	c.CustomAlias = customAlias
	c.OriginalUrl = originalUrl
	return c
}

func (r *CreateUrlRequest) GetCustomAlias() string {
	return r.CustomAlias
}
func (r *CreateUrlRequest) GetOriginalUrl() string {
	return r.OriginalUrl
}

type errorResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

type CreateUrlResponse struct {
	ShortUrl string `json:"shortUrl"`
}
