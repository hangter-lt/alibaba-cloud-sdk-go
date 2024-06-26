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

// DescribeVodDomainReqHitRateData invokes the vod.DescribeVodDomainReqHitRateData API synchronously
func (client *Client) DescribeVodDomainReqHitRateData(request *DescribeVodDomainReqHitRateDataRequest) (response *DescribeVodDomainReqHitRateDataResponse, err error) {
	response = CreateDescribeVodDomainReqHitRateDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVodDomainReqHitRateDataWithChan invokes the vod.DescribeVodDomainReqHitRateData API asynchronously
func (client *Client) DescribeVodDomainReqHitRateDataWithChan(request *DescribeVodDomainReqHitRateDataRequest) (<-chan *DescribeVodDomainReqHitRateDataResponse, <-chan error) {
	responseChan := make(chan *DescribeVodDomainReqHitRateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVodDomainReqHitRateData(request)
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

// DescribeVodDomainReqHitRateDataWithCallback invokes the vod.DescribeVodDomainReqHitRateData API asynchronously
func (client *Client) DescribeVodDomainReqHitRateDataWithCallback(request *DescribeVodDomainReqHitRateDataRequest, callback func(response *DescribeVodDomainReqHitRateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVodDomainReqHitRateDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeVodDomainReqHitRateData(request)
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

// DescribeVodDomainReqHitRateDataRequest is the request struct for api DescribeVodDomainReqHitRateData
type DescribeVodDomainReqHitRateDataRequest struct {
	*requests.RpcRequest
	StartTime  string `position:"Query" name:"StartTime"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	Interval   string `position:"Query" name:"Interval"`
}

// DescribeVodDomainReqHitRateDataResponse is the response struct for api DescribeVodDomainReqHitRateData
type DescribeVodDomainReqHitRateDataResponse struct {
	*responses.BaseResponse
	EndTime      string                                `json:"EndTime" xml:"EndTime"`
	RequestId    string                                `json:"RequestId" xml:"RequestId"`
	StartTime    string                                `json:"StartTime" xml:"StartTime"`
	DomainName   string                                `json:"DomainName" xml:"DomainName"`
	DataInterval string                                `json:"DataInterval" xml:"DataInterval"`
	Data         DataInDescribeVodDomainReqHitRateData `json:"Data" xml:"Data"`
}

// CreateDescribeVodDomainReqHitRateDataRequest creates a request to invoke DescribeVodDomainReqHitRateData API
func CreateDescribeVodDomainReqHitRateDataRequest() (request *DescribeVodDomainReqHitRateDataRequest) {
	request = &DescribeVodDomainReqHitRateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DescribeVodDomainReqHitRateData", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVodDomainReqHitRateDataResponse creates a response to parse from DescribeVodDomainReqHitRateData response
func CreateDescribeVodDomainReqHitRateDataResponse() (response *DescribeVodDomainReqHitRateDataResponse) {
	response = &DescribeVodDomainReqHitRateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
