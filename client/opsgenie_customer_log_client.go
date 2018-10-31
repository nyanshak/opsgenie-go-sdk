package client

import (
	"github.com/opsgenie/opsgenie-go-sdk/customerlog"
)

// OpsGenieLogClient is the data type to make Customer Log rule API requests.
type OpsGenieLogClient struct {
	RestClient
}

// SetOpsGenieClient sets the embedded OpsGenieClient type of the OpsGenieScheduleRotationV2Client.
func (cli *OpsGenieLogClient) SetOpsGenieClient(ogCli OpsGenieClient) {
	cli.OpsGenieClient = ogCli
}

func (cli *OpsGenieLogClient) GetLink(req customerlog.CustomerLogGetLinkRequest) (*customerlog.CustomerLogGetLinkResponse, error) {
	var response customerlog.CustomerLogGetLinkResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (cli *OpsGenieLogClient) DownloadableList(req customerlog.CustomerLogListDownloadablesRequest) (*customerlog.CustomerLogListDownloadablesResponse, error) {
	var response customerlog.CustomerLogListDownloadablesResponse
	err := cli.sendGetRequest(&req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
