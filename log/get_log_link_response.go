package log

type LogGetLinkResponse struct {
	ResponseMeta
	URL string `json:"data"` // todo ask mustafak
}
