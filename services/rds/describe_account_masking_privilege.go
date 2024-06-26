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

// DescribeAccountMaskingPrivilege invokes the rds.DescribeAccountMaskingPrivilege API synchronously
func (client *Client) DescribeAccountMaskingPrivilege(request *DescribeAccountMaskingPrivilegeRequest) (response *DescribeAccountMaskingPrivilegeResponse, err error) {
	response = CreateDescribeAccountMaskingPrivilegeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAccountMaskingPrivilegeWithChan invokes the rds.DescribeAccountMaskingPrivilege API asynchronously
func (client *Client) DescribeAccountMaskingPrivilegeWithChan(request *DescribeAccountMaskingPrivilegeRequest) (<-chan *DescribeAccountMaskingPrivilegeResponse, <-chan error) {
	responseChan := make(chan *DescribeAccountMaskingPrivilegeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccountMaskingPrivilege(request)
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

// DescribeAccountMaskingPrivilegeWithCallback invokes the rds.DescribeAccountMaskingPrivilege API asynchronously
func (client *Client) DescribeAccountMaskingPrivilegeWithCallback(request *DescribeAccountMaskingPrivilegeRequest, callback func(response *DescribeAccountMaskingPrivilegeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccountMaskingPrivilegeResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccountMaskingPrivilege(request)
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

// DescribeAccountMaskingPrivilegeRequest is the request struct for api DescribeAccountMaskingPrivilege
type DescribeAccountMaskingPrivilegeRequest struct {
	*requests.RpcRequest
	DBInstanceName       string           `position:"Query" name:"DBInstanceName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	UserName             string           `position:"Query" name:"UserName"`
}

// DescribeAccountMaskingPrivilegeResponse is the response struct for api DescribeAccountMaskingPrivilege
type DescribeAccountMaskingPrivilegeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeAccountMaskingPrivilegeRequest creates a request to invoke DescribeAccountMaskingPrivilege API
func CreateDescribeAccountMaskingPrivilegeRequest() (request *DescribeAccountMaskingPrivilegeRequest) {
	request = &DescribeAccountMaskingPrivilegeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeAccountMaskingPrivilege", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAccountMaskingPrivilegeResponse creates a response to parse from DescribeAccountMaskingPrivilege response
func CreateDescribeAccountMaskingPrivilegeResponse() (response *DescribeAccountMaskingPrivilegeResponse) {
	response = &DescribeAccountMaskingPrivilegeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
