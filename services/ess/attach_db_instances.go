package ess

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

// AttachDBInstances invokes the ess.AttachDBInstances API synchronously
func (client *Client) AttachDBInstances(request *AttachDBInstancesRequest) (response *AttachDBInstancesResponse, err error) {
	response = CreateAttachDBInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// AttachDBInstancesWithChan invokes the ess.AttachDBInstances API asynchronously
func (client *Client) AttachDBInstancesWithChan(request *AttachDBInstancesRequest) (<-chan *AttachDBInstancesResponse, <-chan error) {
	responseChan := make(chan *AttachDBInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachDBInstances(request)
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

// AttachDBInstancesWithCallback invokes the ess.AttachDBInstances API asynchronously
func (client *Client) AttachDBInstancesWithCallback(request *AttachDBInstancesRequest, callback func(response *AttachDBInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachDBInstancesResponse
		var err error
		defer close(result)
		response, err = client.AttachDBInstances(request)
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

// AttachDBInstancesRequest is the request struct for api AttachDBInstances
type AttachDBInstancesRequest struct {
	*requests.RpcRequest
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	ForceAttach          requests.Boolean `position:"Query" name:"ForceAttach"`
	Type                 string           `position:"Query" name:"Type"`
	AttachMode           string           `position:"Query" name:"AttachMode"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBInstance           *[]string        `position:"Query" name:"DBInstance"  type:"Repeated"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// AttachDBInstancesResponse is the response struct for api AttachDBInstances
type AttachDBInstancesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachDBInstancesRequest creates a request to invoke AttachDBInstances API
func CreateAttachDBInstancesRequest() (request *AttachDBInstancesRequest) {
	request = &AttachDBInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "AttachDBInstances", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachDBInstancesResponse creates a response to parse from AttachDBInstances response
func CreateAttachDBInstancesResponse() (response *AttachDBInstancesResponse) {
	response = &AttachDBInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
