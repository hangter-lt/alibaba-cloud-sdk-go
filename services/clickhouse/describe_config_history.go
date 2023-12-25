package clickhouse

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

// DescribeConfigHistory invokes the clickhouse.DescribeConfigHistory API synchronously
func (client *Client) DescribeConfigHistory(request *DescribeConfigHistoryRequest) (response *DescribeConfigHistoryResponse, err error) {
	response = CreateDescribeConfigHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeConfigHistoryWithChan invokes the clickhouse.DescribeConfigHistory API asynchronously
func (client *Client) DescribeConfigHistoryWithChan(request *DescribeConfigHistoryRequest) (<-chan *DescribeConfigHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeConfigHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeConfigHistory(request)
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

// DescribeConfigHistoryWithCallback invokes the clickhouse.DescribeConfigHistory API asynchronously
func (client *Client) DescribeConfigHistoryWithCallback(request *DescribeConfigHistoryRequest, callback func(response *DescribeConfigHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeConfigHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeConfigHistory(request)
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

// DescribeConfigHistoryRequest is the request struct for api DescribeConfigHistory
type DescribeConfigHistoryRequest struct {
	*requests.RpcRequest
	StartTime   string `position:"Query" name:"StartTime"`
	DBClusterId string `position:"Query" name:"DBClusterId"`
	EndTime     string `position:"Query" name:"EndTime"`
}

// DescribeConfigHistoryResponse is the response struct for api DescribeConfigHistory
type DescribeConfigHistoryResponse struct {
	*responses.BaseResponse
	RequestId          string              `json:"RequestId" xml:"RequestId"`
	ConfigHistoryItems []ConfigHistoryItem `json:"ConfigHistoryItems" xml:"ConfigHistoryItems"`
}

// CreateDescribeConfigHistoryRequest creates a request to invoke DescribeConfigHistory API
func CreateDescribeConfigHistoryRequest() (request *DescribeConfigHistoryRequest) {
	request = &DescribeConfigHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeConfigHistory", "service", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeConfigHistoryResponse creates a response to parse from DescribeConfigHistory response
func CreateDescribeConfigHistoryResponse() (response *DescribeConfigHistoryResponse) {
	response = &DescribeConfigHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
