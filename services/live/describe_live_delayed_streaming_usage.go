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

// DescribeLiveDelayedStreamingUsage invokes the live.DescribeLiveDelayedStreamingUsage API synchronously
func (client *Client) DescribeLiveDelayedStreamingUsage(request *DescribeLiveDelayedStreamingUsageRequest) (response *DescribeLiveDelayedStreamingUsageResponse, err error) {
	response = CreateDescribeLiveDelayedStreamingUsageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDelayedStreamingUsageWithChan invokes the live.DescribeLiveDelayedStreamingUsage API asynchronously
func (client *Client) DescribeLiveDelayedStreamingUsageWithChan(request *DescribeLiveDelayedStreamingUsageRequest) (<-chan *DescribeLiveDelayedStreamingUsageResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDelayedStreamingUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDelayedStreamingUsage(request)
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

// DescribeLiveDelayedStreamingUsageWithCallback invokes the live.DescribeLiveDelayedStreamingUsage API asynchronously
func (client *Client) DescribeLiveDelayedStreamingUsageWithCallback(request *DescribeLiveDelayedStreamingUsageRequest, callback func(response *DescribeLiveDelayedStreamingUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDelayedStreamingUsageResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDelayedStreamingUsage(request)
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

// DescribeLiveDelayedStreamingUsageRequest is the request struct for api DescribeLiveDelayedStreamingUsage
type DescribeLiveDelayedStreamingUsageRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	StreamName string           `position:"Query" name:"StreamName"`
	SplitBy    string           `position:"Query" name:"SplitBy"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Interval   string           `position:"Query" name:"Interval"`
	Region     string           `position:"Query" name:"Region"`
}

// DescribeLiveDelayedStreamingUsageResponse is the response struct for api DescribeLiveDelayedStreamingUsage
type DescribeLiveDelayedStreamingUsageResponse struct {
	*responses.BaseResponse
	EndTime   string    `json:"EndTime" xml:"EndTime"`
	RequestId string    `json:"RequestId" xml:"RequestId"`
	StartTime string    `json:"StartTime" xml:"StartTime"`
	DelayData DelayData `json:"DelayData" xml:"DelayData"`
}

// CreateDescribeLiveDelayedStreamingUsageRequest creates a request to invoke DescribeLiveDelayedStreamingUsage API
func CreateDescribeLiveDelayedStreamingUsageRequest() (request *DescribeLiveDelayedStreamingUsageRequest) {
	request = &DescribeLiveDelayedStreamingUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDelayedStreamingUsage", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDelayedStreamingUsageResponse creates a response to parse from DescribeLiveDelayedStreamingUsage response
func CreateDescribeLiveDelayedStreamingUsageResponse() (response *DescribeLiveDelayedStreamingUsageResponse) {
	response = &DescribeLiveDelayedStreamingUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
