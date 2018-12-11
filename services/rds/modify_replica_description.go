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

// ModifyReplicaDescription invokes the rds.ModifyReplicaDescription API synchronously
// api document: https://help.aliyun.com/api/rds/modifyreplicadescription.html
func (client *Client) ModifyReplicaDescription(request *ModifyReplicaDescriptionRequest) (response *ModifyReplicaDescriptionResponse, err error) {
	response = CreateModifyReplicaDescriptionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyReplicaDescriptionWithChan invokes the rds.ModifyReplicaDescription API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyreplicadescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaDescriptionWithChan(request *ModifyReplicaDescriptionRequest) (<-chan *ModifyReplicaDescriptionResponse, <-chan error) {
	responseChan := make(chan *ModifyReplicaDescriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyReplicaDescription(request)
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

// ModifyReplicaDescriptionWithCallback invokes the rds.ModifyReplicaDescription API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyreplicadescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyReplicaDescriptionWithCallback(request *ModifyReplicaDescriptionRequest, callback func(response *ModifyReplicaDescriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyReplicaDescriptionResponse
		var err error
		defer close(result)
		response, err = client.ModifyReplicaDescription(request)
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

// ModifyReplicaDescriptionRequest is the request struct for api ModifyReplicaDescription
type ModifyReplicaDescriptionRequest struct {
	*requests.RpcRequest
	ReplicaDescription   string           `position:"Query" name:"ReplicaDescription"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyReplicaDescriptionResponse is the response struct for api ModifyReplicaDescription
type ModifyReplicaDescriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyReplicaDescriptionRequest creates a request to invoke ModifyReplicaDescription API
func CreateModifyReplicaDescriptionRequest() (request *ModifyReplicaDescriptionRequest) {
	request = &ModifyReplicaDescriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyReplicaDescription", "Rds", "openAPI")
	return
}

// CreateModifyReplicaDescriptionResponse creates a response to parse from ModifyReplicaDescription response
func CreateModifyReplicaDescriptionResponse() (response *ModifyReplicaDescriptionResponse) {
	response = &ModifyReplicaDescriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
