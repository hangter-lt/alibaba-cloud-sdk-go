package vpcpeer

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

// UnTagResources invokes the vpcpeer.UnTagResources API synchronously
func (client *Client) UnTagResources(request *UnTagResourcesRequest) (response *UnTagResourcesResponse, err error) {
	response = CreateUnTagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// UnTagResourcesWithChan invokes the vpcpeer.UnTagResources API asynchronously
func (client *Client) UnTagResourcesWithChan(request *UnTagResourcesRequest) (<-chan *UnTagResourcesResponse, <-chan error) {
	responseChan := make(chan *UnTagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnTagResources(request)
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

// UnTagResourcesWithCallback invokes the vpcpeer.UnTagResources API asynchronously
func (client *Client) UnTagResourcesWithCallback(request *UnTagResourcesRequest, callback func(response *UnTagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnTagResourcesResponse
		var err error
		defer close(result)
		response, err = client.UnTagResources(request)
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

// UnTagResourcesRequest is the request struct for api UnTagResources
type UnTagResourcesRequest struct {
	*requests.RpcRequest
	ClientToken  string           `position:"Query" name:"ClientToken"`
	All          requests.Boolean `position:"Query" name:"All"`
	ResourceId   *[]string        `position:"Query" name:"ResourceId"  type:"Repeated"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	TagKey       *[]string        `position:"Query" name:"TagKey"  type:"Repeated"`
}

// UnTagResourcesResponse is the response struct for api UnTagResources
type UnTagResourcesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUnTagResourcesRequest creates a request to invoke UnTagResources API
func CreateUnTagResourcesRequest() (request *UnTagResourcesRequest) {
	request = &UnTagResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VpcPeer", "2022-01-01", "UnTagResources", "vpcpeer", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnTagResourcesResponse creates a response to parse from UnTagResources response
func CreateUnTagResourcesResponse() (response *UnTagResourcesResponse) {
	response = &UnTagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
