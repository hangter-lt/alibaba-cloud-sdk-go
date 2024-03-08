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

// ModifyResourceGroup invokes the dds.ModifyResourceGroup API synchronously
func (client *Client) ModifyResourceGroup(request *ModifyResourceGroupRequest) (response *ModifyResourceGroupResponse, err error) {
	response = CreateModifyResourceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyResourceGroupWithChan invokes the dds.ModifyResourceGroup API asynchronously
func (client *Client) ModifyResourceGroupWithChan(request *ModifyResourceGroupRequest) (<-chan *ModifyResourceGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyResourceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyResourceGroup(request)
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

// ModifyResourceGroupWithCallback invokes the dds.ModifyResourceGroup API asynchronously
func (client *Client) ModifyResourceGroupWithCallback(request *ModifyResourceGroupRequest, callback func(response *ModifyResourceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyResourceGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyResourceGroup(request)
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

// ModifyResourceGroupRequest is the request struct for api ModifyResourceGroup
type ModifyResourceGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyResourceGroupResponse is the response struct for api ModifyResourceGroup
type ModifyResourceGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyResourceGroupRequest creates a request to invoke ModifyResourceGroup API
func CreateModifyResourceGroupRequest() (request *ModifyResourceGroupRequest) {
	request = &ModifyResourceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyResourceGroup", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyResourceGroupResponse creates a response to parse from ModifyResourceGroup response
func CreateModifyResourceGroupResponse() (response *ModifyResourceGroupResponse) {
	response = &ModifyResourceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
