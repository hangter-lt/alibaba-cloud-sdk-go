package dcdn

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

// DescribeDcdnWafLogs invokes the dcdn.DescribeDcdnWafLogs API synchronously
func (client *Client) DescribeDcdnWafLogs(request *DescribeDcdnWafLogsRequest) (response *DescribeDcdnWafLogsResponse, err error) {
	response = CreateDescribeDcdnWafLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnWafLogsWithChan invokes the dcdn.DescribeDcdnWafLogs API asynchronously
func (client *Client) DescribeDcdnWafLogsWithChan(request *DescribeDcdnWafLogsRequest) (<-chan *DescribeDcdnWafLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnWafLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnWafLogs(request)
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

// DescribeDcdnWafLogsWithCallback invokes the dcdn.DescribeDcdnWafLogs API asynchronously
func (client *Client) DescribeDcdnWafLogsWithCallback(request *DescribeDcdnWafLogsRequest, callback func(response *DescribeDcdnWafLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnWafLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnWafLogs(request)
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

// DescribeDcdnWafLogsRequest is the request struct for api DescribeDcdnWafLogs
type DescribeDcdnWafLogsRequest struct {
	*requests.RpcRequest
	LogType    string           `position:"Query" name:"LogType"`
	DomainName string           `position:"Query" name:"DomainName"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	EndTime    string           `position:"Query" name:"EndTime"`
	StartTime  string           `position:"Query" name:"StartTime"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeDcdnWafLogsResponse is the response struct for api DescribeDcdnWafLogs
type DescribeDcdnWafLogsResponse struct {
	*responses.BaseResponse
	RequestId        string                                 `json:"RequestId" xml:"RequestId"`
	DomainLogDetails []DomainLogDetailInDescribeDcdnWafLogs `json:"DomainLogDetails" xml:"DomainLogDetails"`
}

// CreateDescribeDcdnWafLogsRequest creates a request to invoke DescribeDcdnWafLogs API
func CreateDescribeDcdnWafLogsRequest() (request *DescribeDcdnWafLogsRequest) {
	request = &DescribeDcdnWafLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnWafLogs", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnWafLogsResponse creates a response to parse from DescribeDcdnWafLogs response
func CreateDescribeDcdnWafLogsResponse() (response *DescribeDcdnWafLogsResponse) {
	response = &DescribeDcdnWafLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
