package log

type LogListDownloadablesResponse struct {
	ResponseMeta
	Downloadables []string `json:"data"` // todo ask mustafak
}
