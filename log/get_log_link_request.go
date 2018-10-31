package log

import (
	"net/url"
)

// LogGetLinkRequest is a struct of request to list new log.
type LogGetLinkRequest struct {
	LogFile string
	ApiKey  string
}

// GetApiKey returns api key.
func (r *LogGetLinkRequest) GetApiKey() string {
	return r.ApiKey
}

func (r *LogGetLinkRequest) GetLogFile() string {
	return r.LogFile
}

// GenerateUrl generates url to API endpoint.
func (r *LogGetLinkRequest) GenerateUrl() (string, url.Values, error) {
	// todo kagan fix here
	//baseUrl, params, err := r.ScheduleIdentifier.GenerateUrl()
	//
	//if err != nil {
	//	return "", nil, err
	//}
	//
	//baseUrl += "/overrides"
	//return baseUrl, params, nil
	return "", nil, nil
}
