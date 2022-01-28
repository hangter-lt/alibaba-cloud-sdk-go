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

// DescribeLorneTasks invokes the clickhouse.DescribeLorneTasks API synchronously
func (client *Client) DescribeLorneTasks(request *DescribeLorneTasksRequest) (response *DescribeLorneTasksResponse, err error) {
	response = CreateDescribeLorneTasksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLorneTasksWithChan invokes the clickhouse.DescribeLorneTasks API asynchronously
func (client *Client) DescribeLorneTasksWithChan(request *DescribeLorneTasksRequest) (<-chan *DescribeLorneTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeLorneTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLorneTasks(request)
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

// DescribeLorneTasksWithCallback invokes the clickhouse.DescribeLorneTasks API asynchronously
func (client *Client) DescribeLorneTasksWithCallback(request *DescribeLorneTasksRequest, callback func(response *DescribeLorneTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLorneTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeLorneTasks(request)
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

// DescribeLorneTasksRequest is the request struct for api DescribeLorneTasks
type DescribeLorneTasksRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLorneTasksResponse is the response struct for api DescribeLorneTasks
type DescribeLorneTasksResponse struct {
	*responses.BaseResponse
	TotalCount  int          `json:"TotalCount" xml:"TotalCount"`
	PageSize    int          `json:"PageSize" xml:"PageSize"`
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	PageNumber  int          `json:"PageNumber" xml:"PageNumber"`
	TaskDetails []TaskDetail `json:"TaskDetails" xml:"TaskDetails"`
}

// CreateDescribeLorneTasksRequest creates a request to invoke DescribeLorneTasks API
func CreateDescribeLorneTasksRequest() (request *DescribeLorneTasksRequest) {
	request = &DescribeLorneTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeLorneTasks", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeLorneTasksResponse creates a response to parse from DescribeLorneTasks response
func CreateDescribeLorneTasksResponse() (response *DescribeLorneTasksResponse) {
	response = &DescribeLorneTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}