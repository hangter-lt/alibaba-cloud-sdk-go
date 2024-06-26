package vod

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

// DescribeVodDomainRealTimeTrafficData invokes the vod.DescribeVodDomainRealTimeTrafficData API synchronously
func (client *Client) DescribeVodDomainRealTimeTrafficData(request *DescribeVodDomainRealTimeTrafficDataRequest) (response *DescribeVodDomainRealTimeTrafficDataResponse, err error) {
	response = CreateDescribeVodDomainRealTimeTrafficDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodDomainRealTimeTrafficDataWithChan invokes the vod.DescribeVodDomainRealTimeTrafficData API asynchronously
func (client *Client) DescribeVodDomainRealTimeTrafficDataWithChan(request *DescribeVodDomainRealTimeTrafficDataRequest) (<-chan *DescribeVodDomainRealTimeTrafficDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVodDomainRealTimeTrafficDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodDomainRealTimeTrafficData(request)
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

// DescribeVodDomainRealTimeTrafficDataWithCallback invokes the vod.DescribeVodDomainRealTimeTrafficData API asynchronously
func (client *Client) DescribeVodDomainRealTimeTrafficDataWithCallback(request *DescribeVodDomainRealTimeTrafficDataRequest, callback func(response *DescribeVodDomainRealTimeTrafficDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodDomainRealTimeTrafficDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodDomainRealTimeTrafficData(request)
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

// DescribeVodDomainRealTimeTrafficDataRequest is the request struct for api DescribeVodDomainRealTimeTrafficData
type DescribeVodDomainRealTimeTrafficDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeVodDomainRealTimeTrafficDataResponse is the response struct for api DescribeVodDomainRealTimeTrafficData
type DescribeVodDomainRealTimeTrafficDataResponse struct {
	*responses.BaseResponse
	EndTime                        string                         `json:"EndTime" xml:"EndTime"`
	StartTime                      string                         `json:"StartTime" xml:"StartTime"`
	RequestId                      string                         `json:"RequestId" xml:"RequestId"`
	DomainName                     string                         `json:"DomainName" xml:"DomainName"`
	DataInterval                   string                         `json:"DataInterval" xml:"DataInterval"`
	RealTimeTrafficDataPerInterval RealTimeTrafficDataPerInterval `json:"RealTimeTrafficDataPerInterval" xml:"RealTimeTrafficDataPerInterval"`
}

// CreateDescribeVodDomainRealTimeTrafficDataRequest creates a request to invoke DescribeVodDomainRealTimeTrafficData API
func CreateDescribeVodDomainRealTimeTrafficDataRequest() (request *DescribeVodDomainRealTimeTrafficDataRequest) {
	request = &DescribeVodDomainRealTimeTrafficDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodDomainRealTimeTrafficData", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodDomainRealTimeTrafficDataResponse creates a response to parse from DescribeVodDomainRealTimeTrafficData response
func CreateDescribeVodDomainRealTimeTrafficDataResponse() (response *DescribeVodDomainRealTimeTrafficDataResponse) {
	response = &DescribeVodDomainRealTimeTrafficDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
