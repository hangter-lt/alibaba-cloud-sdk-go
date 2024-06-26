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

// CreatePostgresExtensions invokes the rds.CreatePostgresExtensions API synchronously
func (client *Client) CreatePostgresExtensions(request *CreatePostgresExtensionsRequest) (response *CreatePostgresExtensionsResponse, err error) {
	response = CreateCreatePostgresExtensionsResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePostgresExtensionsWithChan invokes the rds.CreatePostgresExtensions API asynchronously
func (client *Client) CreatePostgresExtensionsWithChan(request *CreatePostgresExtensionsRequest) (<-chan *CreatePostgresExtensionsResponse, <-chan error) {
	responseChan := make(chan *CreatePostgresExtensionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePostgresExtensions(request)
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

// CreatePostgresExtensionsWithCallback invokes the rds.CreatePostgresExtensions API asynchronously
func (client *Client) CreatePostgresExtensionsWithCallback(request *CreatePostgresExtensionsRequest, callback func(response *CreatePostgresExtensionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePostgresExtensionsResponse
		var err error
		defer close(result)
		response, err = client.CreatePostgresExtensions(request)
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

// CreatePostgresExtensionsRequest is the request struct for api CreatePostgresExtensions
type CreatePostgresExtensionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	RiskConfirmed        requests.Boolean `position:"Query" name:"RiskConfirmed"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	AccountName          string           `position:"Query" name:"AccountName"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	SourceDatabase       string           `position:"Query" name:"SourceDatabase"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBNames              string           `position:"Query" name:"DBNames"`
	Extensions           string           `position:"Query" name:"Extensions"`
}

// CreatePostgresExtensionsResponse is the response struct for api CreatePostgresExtensions
type CreatePostgresExtensionsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreatePostgresExtensionsRequest creates a request to invoke CreatePostgresExtensions API
func CreateCreatePostgresExtensionsRequest() (request *CreatePostgresExtensionsRequest) {
	request = &CreatePostgresExtensionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreatePostgresExtensions", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePostgresExtensionsResponse creates a response to parse from CreatePostgresExtensions response
func CreateCreatePostgresExtensionsResponse() (response *CreatePostgresExtensionsResponse) {
	response = &CreatePostgresExtensionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
