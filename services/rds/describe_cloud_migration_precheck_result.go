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

// DescribeCloudMigrationPrecheckResult invokes the rds.DescribeCloudMigrationPrecheckResult API synchronously
func (client *Client) DescribeCloudMigrationPrecheckResult(request *DescribeCloudMigrationPrecheckResultRequest) (response *DescribeCloudMigrationPrecheckResultResponse, err error) {
	response = CreateDescribeCloudMigrationPrecheckResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudMigrationPrecheckResultWithChan invokes the rds.DescribeCloudMigrationPrecheckResult API asynchronously
func (client *Client) DescribeCloudMigrationPrecheckResultWithChan(request *DescribeCloudMigrationPrecheckResultRequest) (<-chan *DescribeCloudMigrationPrecheckResultResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudMigrationPrecheckResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudMigrationPrecheckResult(request)
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

// DescribeCloudMigrationPrecheckResultWithCallback invokes the rds.DescribeCloudMigrationPrecheckResult API asynchronously
func (client *Client) DescribeCloudMigrationPrecheckResultWithCallback(request *DescribeCloudMigrationPrecheckResultRequest, callback func(response *DescribeCloudMigrationPrecheckResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudMigrationPrecheckResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudMigrationPrecheckResult(request)
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

// DescribeCloudMigrationPrecheckResultRequest is the request struct for api DescribeCloudMigrationPrecheckResult
type DescribeCloudMigrationPrecheckResultRequest struct {
	*requests.RpcRequest
	DBInstanceName       string           `position:"Query" name:"DBInstanceName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TaskName             string           `position:"Query" name:"TaskName"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SourcePort           requests.Integer `position:"Query" name:"SourcePort"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	TaskId               requests.Integer `position:"Query" name:"TaskId"`
	SourceIpAddress      string           `position:"Query" name:"SourceIpAddress"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCloudMigrationPrecheckResultResponse is the response struct for api DescribeCloudMigrationPrecheckResult
type DescribeCloudMigrationPrecheckResultResponse struct {
	*responses.BaseResponse
	TotalSize  int                    `json:"TotalSize" xml:"TotalSize"`
	RequestId  string                 `json:"RequestId" xml:"RequestId"`
	PageNumber int64                  `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64                  `json:"PageSize" xml:"PageSize"`
	Items      []MigrateCloudTaskList `json:"Items" xml:"Items"`
}

// CreateDescribeCloudMigrationPrecheckResultRequest creates a request to invoke DescribeCloudMigrationPrecheckResult API
func CreateDescribeCloudMigrationPrecheckResultRequest() (request *DescribeCloudMigrationPrecheckResultRequest) {
	request = &DescribeCloudMigrationPrecheckResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeCloudMigrationPrecheckResult", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudMigrationPrecheckResultResponse creates a response to parse from DescribeCloudMigrationPrecheckResult response
func CreateDescribeCloudMigrationPrecheckResultResponse() (response *DescribeCloudMigrationPrecheckResultResponse) {
	response = &DescribeCloudMigrationPrecheckResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
