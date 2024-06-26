package live

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeLiveStreamPushMetricDetailData invokes the live.DescribeLiveStreamPushMetricDetailData API synchronously
func (client *Client) DescribeLiveStreamPushMetricDetailData(request *DescribeLiveStreamPushMetricDetailDataRequest) (response *DescribeLiveStreamPushMetricDetailDataResponse, err error) {
	response = CreateDescribeLiveStreamPushMetricDetailDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamPushMetricDetailDataWithChan invokes the live.DescribeLiveStreamPushMetricDetailData API asynchronously
func (client *Client) DescribeLiveStreamPushMetricDetailDataWithChan(request *DescribeLiveStreamPushMetricDetailDataRequest) (<-chan *DescribeLiveStreamPushMetricDetailDataResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamPushMetricDetailDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamPushMetricDetailData(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeLiveStreamPushMetricDetailDataWithCallback invokes the live.DescribeLiveStreamPushMetricDetailData API asynchronously
func (client *Client) DescribeLiveStreamPushMetricDetailDataWithCallback(request *DescribeLiveStreamPushMetricDetailDataRequest, callback func(response *DescribeLiveStreamPushMetricDetailDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamPushMetricDetailDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamPushMetricDetailData(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeLiveStreamPushMetricDetailDataRequest is the request struct for api DescribeLiveStreamPushMetricDetailData
type DescribeLiveStreamPushMetricDetailDataRequest struct {
	*requests.RpcRequest
	NextPageToken string           `position:"Query" name:"NextPageToken"`
	StartTime     string           `position:"Query" name:"StartTime"`
	AppName       string           `position:"Query" name:"AppName"`
	StreamName    string           `position:"Query" name:"StreamName"`
	DomainName    string           `position:"Query" name:"DomainName"`
	EndTime       string           `position:"Query" name:"EndTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamPushMetricDetailDataResponse is the response struct for api DescribeLiveStreamPushMetricDetailData
type DescribeLiveStreamPushMetricDetailDataResponse struct {
	*responses.BaseResponse
	DomainName       string                                                   `json:"DomainName" xml:"DomainName"`
	EndTime          string                                                   `json:"EndTime" xml:"EndTime"`
	NextPageToken    string                                                   `json:"NextPageToken" xml:"NextPageToken"`
	PageSize         int                                                      `json:"PageSize" xml:"PageSize"`
	RequestId        string                                                   `json:"RequestId" xml:"RequestId"`
	StartTime        string                                                   `json:"StartTime" xml:"StartTime"`
	StreamDetailData StreamDetailDataInDescribeLiveStreamPushMetricDetailData `json:"StreamDetailData" xml:"StreamDetailData"`
}

// CreateDescribeLiveStreamPushMetricDetailDataRequest creates a request to invoke DescribeLiveStreamPushMetricDetailData API
func CreateDescribeLiveStreamPushMetricDetailDataRequest() (request *DescribeLiveStreamPushMetricDetailDataRequest) {
	request = &DescribeLiveStreamPushMetricDetailDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamPushMetricDetailData", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveStreamPushMetricDetailDataResponse creates a response to parse from DescribeLiveStreamPushMetricDetailData response
func CreateDescribeLiveStreamPushMetricDetailDataResponse() (response *DescribeLiveStreamPushMetricDetailDataResponse) {
	response = &DescribeLiveStreamPushMetricDetailDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
