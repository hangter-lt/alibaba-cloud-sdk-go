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

// CreateCloudMigrationTask invokes the rds.CreateCloudMigrationTask API synchronously
func (client *Client) CreateCloudMigrationTask(request *CreateCloudMigrationTaskRequest) (response *CreateCloudMigrationTaskResponse, err error) {
	response = CreateCreateCloudMigrationTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCloudMigrationTaskWithChan invokes the rds.CreateCloudMigrationTask API asynchronously
func (client *Client) CreateCloudMigrationTaskWithChan(request *CreateCloudMigrationTaskRequest) (<-chan *CreateCloudMigrationTaskResponse, <-chan error) {
	responseChan := make(chan *CreateCloudMigrationTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCloudMigrationTask(request)
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

// CreateCloudMigrationTaskWithCallback invokes the rds.CreateCloudMigrationTask API asynchronously
func (client *Client) CreateCloudMigrationTaskWithCallback(request *CreateCloudMigrationTaskRequest, callback func(response *CreateCloudMigrationTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCloudMigrationTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateCloudMigrationTask(request)
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

// CreateCloudMigrationTaskRequest is the request struct for api CreateCloudMigrationTask
type CreateCloudMigrationTaskRequest struct {
	*requests.RpcRequest
	DBInstanceName       string           `position:"Query" name:"DBInstanceName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TaskName             string           `position:"Query" name:"TaskName"`
	SourceAccount        string           `position:"Query" name:"SourceAccount"`
	SourcePort           requests.Integer `position:"Query" name:"SourcePort"`
	SourcePassword       string           `position:"Query" name:"SourcePassword"`
	SourceIpAddress      string           `position:"Query" name:"SourceIpAddress"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SourceCategory       string           `position:"Query" name:"SourceCategory"`
}

// CreateCloudMigrationTaskResponse is the response struct for api CreateCloudMigrationTask
type CreateCloudMigrationTaskResponse struct {
	*responses.BaseResponse
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	TaskId         int64  `json:"TaskId" xml:"TaskId"`
	TaskName       string `json:"TaskName" xml:"TaskName"`
}

// CreateCreateCloudMigrationTaskRequest creates a request to invoke CreateCloudMigrationTask API
func CreateCreateCloudMigrationTaskRequest() (request *CreateCloudMigrationTaskRequest) {
	request = &CreateCloudMigrationTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateCloudMigrationTask", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCloudMigrationTaskResponse creates a response to parse from CreateCloudMigrationTask response
func CreateCreateCloudMigrationTaskResponse() (response *CreateCloudMigrationTaskResponse) {
	response = &CreateCloudMigrationTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
