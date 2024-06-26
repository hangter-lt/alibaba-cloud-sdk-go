package live

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

// DeleteCasterComponent invokes the live.DeleteCasterComponent API synchronously
func (client *Client) DeleteCasterComponent(request *DeleteCasterComponentRequest) (response *DeleteCasterComponentResponse, err error) {
	response = CreateDeleteCasterComponentResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCasterComponentWithChan invokes the live.DeleteCasterComponent API asynchronously
func (client *Client) DeleteCasterComponentWithChan(request *DeleteCasterComponentRequest) (<-chan *DeleteCasterComponentResponse, <-chan error) {
	responseChan := make(chan *DeleteCasterComponentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCasterComponent(request)
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

// DeleteCasterComponentWithCallback invokes the live.DeleteCasterComponent API asynchronously
func (client *Client) DeleteCasterComponentWithCallback(request *DeleteCasterComponentRequest, callback func(response *DeleteCasterComponentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCasterComponentResponse
		var err error
		defer close(result)
		response, err = client.DeleteCasterComponent(request)
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

// DeleteCasterComponentRequest is the request struct for api DeleteCasterComponent
type DeleteCasterComponentRequest struct {
	*requests.RpcRequest
	ComponentId string           `position:"Query" name:"ComponentId"`
	CasterId    string           `position:"Query" name:"CasterId"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteCasterComponentResponse is the response struct for api DeleteCasterComponent
type DeleteCasterComponentResponse struct {
	*responses.BaseResponse
	CasterId    string `json:"CasterId" xml:"CasterId"`
	ComponentId string `json:"ComponentId" xml:"ComponentId"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCasterComponentRequest creates a request to invoke DeleteCasterComponent API
func CreateDeleteCasterComponentRequest() (request *DeleteCasterComponentRequest) {
	request = &DeleteCasterComponentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteCasterComponent", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteCasterComponentResponse creates a response to parse from DeleteCasterComponent response
func CreateDeleteCasterComponentResponse() (response *DeleteCasterComponentResponse) {
	response = &DeleteCasterComponentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
