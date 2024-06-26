package sls

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

// InitProjectAlertResource invokes the sls.InitProjectAlertResource API synchronously
func (client *Client) InitProjectAlertResource(request *InitProjectAlertResourceRequest) (response *InitProjectAlertResourceResponse, err error) {
	response = CreateInitProjectAlertResourceResponse()
	err = client.DoAction(request, response)
	return
}

// InitProjectAlertResourceWithChan invokes the sls.InitProjectAlertResource API asynchronously
func (client *Client) InitProjectAlertResourceWithChan(request *InitProjectAlertResourceRequest) (<-chan *InitProjectAlertResourceResponse, <-chan error) {
	responseChan := make(chan *InitProjectAlertResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitProjectAlertResource(request)
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

// InitProjectAlertResourceWithCallback invokes the sls.InitProjectAlertResource API asynchronously
func (client *Client) InitProjectAlertResourceWithCallback(request *InitProjectAlertResourceRequest, callback func(response *InitProjectAlertResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitProjectAlertResourceResponse
		var err error
		defer close(result)
		response, err = client.InitProjectAlertResource(request)
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

// InitProjectAlertResourceRequest is the request struct for api InitProjectAlertResource
type InitProjectAlertResourceRequest struct {
	*requests.RpcRequest
	ProjectName    string `position:"Body" name:"ProjectName"`
	SlsAccessToken string `position:"Query" name:"SlsAccessToken"`
	Region         string `position:"Body" name:"Region"`
}

// InitProjectAlertResourceResponse is the response struct for api InitProjectAlertResource
type InitProjectAlertResourceResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateInitProjectAlertResourceRequest creates a request to invoke InitProjectAlertResource API
func CreateInitProjectAlertResourceRequest() (request *InitProjectAlertResourceRequest) {
	request = &InitProjectAlertResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2019-10-23", "InitProjectAlertResource", "sls", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInitProjectAlertResourceResponse creates a response to parse from InitProjectAlertResource response
func CreateInitProjectAlertResourceResponse() (response *InitProjectAlertResourceResponse) {
	response = &InitProjectAlertResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
