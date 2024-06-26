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

// UpdatePostgresExtensions invokes the rds.UpdatePostgresExtensions API synchronously
func (client *Client) UpdatePostgresExtensions(request *UpdatePostgresExtensionsRequest) (response *UpdatePostgresExtensionsResponse, err error) {
	response = CreateUpdatePostgresExtensionsResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePostgresExtensionsWithChan invokes the rds.UpdatePostgresExtensions API asynchronously
func (client *Client) UpdatePostgresExtensionsWithChan(request *UpdatePostgresExtensionsRequest) (<-chan *UpdatePostgresExtensionsResponse, <-chan error) {
	responseChan := make(chan *UpdatePostgresExtensionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePostgresExtensions(request)
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

// UpdatePostgresExtensionsWithCallback invokes the rds.UpdatePostgresExtensions API asynchronously
func (client *Client) UpdatePostgresExtensionsWithCallback(request *UpdatePostgresExtensionsRequest, callback func(response *UpdatePostgresExtensionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePostgresExtensionsResponse
		var err error
		defer close(result)
		response, err = client.UpdatePostgresExtensions(request)
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

// UpdatePostgresExtensionsRequest is the request struct for api UpdatePostgresExtensions
type UpdatePostgresExtensionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBNames              string           `position:"Query" name:"DBNames"`
	Extensions           string           `position:"Query" name:"Extensions"`
}

// UpdatePostgresExtensionsResponse is the response struct for api UpdatePostgresExtensions
type UpdatePostgresExtensionsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdatePostgresExtensionsRequest creates a request to invoke UpdatePostgresExtensions API
func CreateUpdatePostgresExtensionsRequest() (request *UpdatePostgresExtensionsRequest) {
	request = &UpdatePostgresExtensionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "UpdatePostgresExtensions", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePostgresExtensionsResponse creates a response to parse from UpdatePostgresExtensions response
func CreateUpdatePostgresExtensionsResponse() (response *UpdatePostgresExtensionsResponse) {
	response = &UpdatePostgresExtensionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
