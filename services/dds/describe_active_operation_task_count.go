package dds

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

// DescribeActiveOperationTaskCount invokes the dds.DescribeActiveOperationTaskCount API synchronously
func (client *Client) DescribeActiveOperationTaskCount(request *DescribeActiveOperationTaskCountRequest) (response *DescribeActiveOperationTaskCountResponse, err error) {
	response = CreateDescribeActiveOperationTaskCountResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeActiveOperationTaskCountWithChan invokes the dds.DescribeActiveOperationTaskCount API asynchronously
func (client *Client) DescribeActiveOperationTaskCountWithChan(request *DescribeActiveOperationTaskCountRequest) (<-chan *DescribeActiveOperationTaskCountResponse, <-chan error) {
	responseChan := make(chan *DescribeActiveOperationTaskCountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeActiveOperationTaskCount(request)
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

// DescribeActiveOperationTaskCountWithCallback invokes the dds.DescribeActiveOperationTaskCount API asynchronously
func (client *Client) DescribeActiveOperationTaskCountWithCallback(request *DescribeActiveOperationTaskCountRequest, callback func(response *DescribeActiveOperationTaskCountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeActiveOperationTaskCountResponse
		var err error
		defer close(result)
		response, err = client.DescribeActiveOperationTaskCount(request)
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

// DescribeActiveOperationTaskCountRequest is the request struct for api DescribeActiveOperationTaskCount
type DescribeActiveOperationTaskCountRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeActiveOperationTaskCountResponse is the response struct for api DescribeActiveOperationTaskCount
type DescribeActiveOperationTaskCountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	NeedPop   int    `json:"NeedPop" xml:"NeedPop"`
	TaskCount int    `json:"TaskCount" xml:"TaskCount"`
}

// CreateDescribeActiveOperationTaskCountRequest creates a request to invoke DescribeActiveOperationTaskCount API
func CreateDescribeActiveOperationTaskCountRequest() (request *DescribeActiveOperationTaskCountRequest) {
	request = &DescribeActiveOperationTaskCountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeActiveOperationTaskCount", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeActiveOperationTaskCountResponse creates a response to parse from DescribeActiveOperationTaskCount response
func CreateDescribeActiveOperationTaskCountResponse() (response *DescribeActiveOperationTaskCountResponse) {
	response = &DescribeActiveOperationTaskCountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
