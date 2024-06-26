package rds

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

// DescribeErrorLogs invokes the rds.DescribeErrorLogs API synchronously
func (client *Client) DescribeErrorLogs(request *DescribeErrorLogsRequest) (response *DescribeErrorLogsResponse, err error) {
	response = CreateDescribeErrorLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeErrorLogsWithChan invokes the rds.DescribeErrorLogs API asynchronously
func (client *Client) DescribeErrorLogsWithChan(request *DescribeErrorLogsRequest) (<-chan *DescribeErrorLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeErrorLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeErrorLogs(request)
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

// DescribeErrorLogsWithCallback invokes the rds.DescribeErrorLogs API asynchronously
func (client *Client) DescribeErrorLogsWithCallback(request *DescribeErrorLogsRequest, callback func(response *DescribeErrorLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeErrorLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeErrorLogs(request)
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

// DescribeErrorLogsRequest is the request struct for api DescribeErrorLogs
type DescribeErrorLogsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeErrorLogsResponse is the response struct for api DescribeErrorLogs
type DescribeErrorLogsResponse struct {
	*responses.BaseResponse
	PageNumber       int                      `json:"PageNumber" xml:"PageNumber"`
	RequestId        string                   `json:"RequestId" xml:"RequestId"`
	PageRecordCount  int                      `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int                      `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Items            ItemsInDescribeErrorLogs `json:"Items" xml:"Items"`
}

// CreateDescribeErrorLogsRequest creates a request to invoke DescribeErrorLogs API
func CreateDescribeErrorLogsRequest() (request *DescribeErrorLogsRequest) {
	request = &DescribeErrorLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeErrorLogs", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeErrorLogsResponse creates a response to parse from DescribeErrorLogs response
func CreateDescribeErrorLogsResponse() (response *DescribeErrorLogsResponse) {
	response = &DescribeErrorLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
