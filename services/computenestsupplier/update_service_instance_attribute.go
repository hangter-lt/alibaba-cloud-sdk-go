package computenestsupplier

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

// UpdateServiceInstanceAttribute invokes the computenestsupplier.UpdateServiceInstanceAttribute API synchronously
func (client *Client) UpdateServiceInstanceAttribute(request *UpdateServiceInstanceAttributeRequest) (response *UpdateServiceInstanceAttributeResponse, err error) {
	response = CreateUpdateServiceInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateServiceInstanceAttributeWithChan invokes the computenestsupplier.UpdateServiceInstanceAttribute API asynchronously
func (client *Client) UpdateServiceInstanceAttributeWithChan(request *UpdateServiceInstanceAttributeRequest) (<-chan *UpdateServiceInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *UpdateServiceInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateServiceInstanceAttribute(request)
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

// UpdateServiceInstanceAttributeWithCallback invokes the computenestsupplier.UpdateServiceInstanceAttribute API asynchronously
func (client *Client) UpdateServiceInstanceAttributeWithCallback(request *UpdateServiceInstanceAttributeRequest, callback func(response *UpdateServiceInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateServiceInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.UpdateServiceInstanceAttribute(request)
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

// UpdateServiceInstanceAttributeRequest is the request struct for api UpdateServiceInstanceAttribute
type UpdateServiceInstanceAttributeRequest struct {
	*requests.RpcRequest
	ServiceInstanceId string `position:"Query" name:"ServiceInstanceId"`
	EndTime           string `position:"Query" name:"EndTime"`
}

// UpdateServiceInstanceAttributeResponse is the response struct for api UpdateServiceInstanceAttribute
type UpdateServiceInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateServiceInstanceAttributeRequest creates a request to invoke UpdateServiceInstanceAttribute API
func CreateUpdateServiceInstanceAttributeRequest() (request *UpdateServiceInstanceAttributeRequest) {
	request = &UpdateServiceInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "UpdateServiceInstanceAttribute", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateServiceInstanceAttributeResponse creates a response to parse from UpdateServiceInstanceAttribute response
func CreateUpdateServiceInstanceAttributeResponse() (response *UpdateServiceInstanceAttributeResponse) {
	response = &UpdateServiceInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
