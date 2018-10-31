package log

import (
	"net/url"
)

// LogListDownloadablesRequest is a struct of request to download customerLog.
type LogListDownloadablesRequest struct {
	After  string
	ApiKey string
}

// GetApiKey returns api key.
func (r *LogListDownloadablesRequest) GetApiKey() string {
	return r.ApiKey
}

// GenerateUrl generates url to API endpoint.
func (r *LogListDownloadablesRequest) GenerateUrl() (string, url.Values, error) {
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
