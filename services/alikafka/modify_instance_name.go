package alikafka

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

// ModifyInstanceName invokes the alikafka.ModifyInstanceName API synchronously
func (client *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
	response = CreateModifyInstanceNameResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceNameWithChan invokes the alikafka.ModifyInstanceName API asynchronously
func (client *Client) ModifyInstanceNameWithChan(request *ModifyInstanceNameRequest) (<-chan *ModifyInstanceNameResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceName(request)
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

// ModifyInstanceNameWithCallback invokes the alikafka.ModifyInstanceName API asynchronously
func (client *Client) ModifyInstanceNameWithCallback(request *ModifyInstanceNameRequest, callback func(response *ModifyInstanceNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceNameResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceName(request)
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

// ModifyInstanceNameRequest is the request struct for api ModifyInstanceName
type ModifyInstanceNameRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	InstanceName string `position:"Query" name:"InstanceName"`
}

// ModifyInstanceNameResponse is the response struct for api ModifyInstanceName
type ModifyInstanceNameResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateModifyInstanceNameRequest creates a request to invoke ModifyInstanceName API
func CreateModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
	request = &ModifyInstanceNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "ModifyInstanceName", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceNameResponse creates a response to parse from ModifyInstanceName response
func CreateModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
	response = &ModifyInstanceNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
