package arms

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

// OpenArmsDefaultSLR invokes the arms.OpenArmsDefaultSLR API synchronously
func (client *Client) OpenArmsDefaultSLR(request *OpenArmsDefaultSLRRequest) (response *OpenArmsDefaultSLRResponse, err error) {
	response = CreateOpenArmsDefaultSLRResponse()
	err = client.DoAction(request, response)
	return
}

// OpenArmsDefaultSLRWithChan invokes the arms.OpenArmsDefaultSLR API asynchronously
func (client *Client) OpenArmsDefaultSLRWithChan(request *OpenArmsDefaultSLRRequest) (<-chan *OpenArmsDefaultSLRResponse, <-chan error) {
	responseChan := make(chan *OpenArmsDefaultSLRResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenArmsDefaultSLR(request)
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

// OpenArmsDefaultSLRWithCallback invokes the arms.OpenArmsDefaultSLR API asynchronously
func (client *Client) OpenArmsDefaultSLRWithCallback(request *OpenArmsDefaultSLRRequest, callback func(response *OpenArmsDefaultSLRResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenArmsDefaultSLRResponse
		var err error
		defer close(result)
		response, err = client.OpenArmsDefaultSLR(request)
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

// OpenArmsDefaultSLRRequest is the request struct for api OpenArmsDefaultSLR
type OpenArmsDefaultSLRRequest struct {
	*requests.RpcRequest
}

// OpenArmsDefaultSLRResponse is the response struct for api OpenArmsDefaultSLR
type OpenArmsDefaultSLRResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOpenArmsDefaultSLRRequest creates a request to invoke OpenArmsDefaultSLR API
func CreateOpenArmsDefaultSLRRequest() (request *OpenArmsDefaultSLRRequest) {
	request = &OpenArmsDefaultSLRRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2021-05-19", "OpenArmsDefaultSLR", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOpenArmsDefaultSLRResponse creates a response to parse from OpenArmsDefaultSLR response
func CreateOpenArmsDefaultSLRResponse() (response *OpenArmsDefaultSLRResponse) {
	response = &OpenArmsDefaultSLRResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
