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

// DescribeHostWebShell invokes the rds.DescribeHostWebShell API synchronously
func (client *Client) DescribeHostWebShell(request *DescribeHostWebShellRequest) (response *DescribeHostWebShellResponse, err error) {
	response = CreateDescribeHostWebShellResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHostWebShellWithChan invokes the rds.DescribeHostWebShell API asynchronously
func (client *Client) DescribeHostWebShellWithChan(request *DescribeHostWebShellRequest) (<-chan *DescribeHostWebShellResponse, <-chan error) {
	responseChan := make(chan *DescribeHostWebShellResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHostWebShell(request)
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

// DescribeHostWebShellWithCallback invokes the rds.DescribeHostWebShell API asynchronously
func (client *Client) DescribeHostWebShellWithCallback(request *DescribeHostWebShellRequest, callback func(response *DescribeHostWebShellResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHostWebShellResponse
		var err error
		defer close(result)
		response, err = client.DescribeHostWebShell(request)
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

// DescribeHostWebShellRequest is the request struct for api DescribeHostWebShell
type DescribeHostWebShellRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	HostName             string           `position:"Query" name:"HostName"`
	AccountName          string           `position:"Query" name:"AccountName"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AccountPassword      string           `position:"Query" name:"AccountPassword"`
}

// DescribeHostWebShellResponse is the response struct for api DescribeHostWebShell
type DescribeHostWebShellResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	LoginUrl  string `json:"LoginUrl" xml:"LoginUrl"`
}

// CreateDescribeHostWebShellRequest creates a request to invoke DescribeHostWebShell API
func CreateDescribeHostWebShellRequest() (request *DescribeHostWebShellRequest) {
	request = &DescribeHostWebShellRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeHostWebShell", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHostWebShellResponse creates a response to parse from DescribeHostWebShell response
func CreateDescribeHostWebShellResponse() (response *DescribeHostWebShellResponse) {
	response = &DescribeHostWebShellResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
