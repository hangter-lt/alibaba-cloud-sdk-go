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

// EnableAlertTemplate invokes the arms.EnableAlertTemplate API synchronously
func (client *Client) EnableAlertTemplate(request *EnableAlertTemplateRequest) (response *EnableAlertTemplateResponse, err error) {
	response = CreateEnableAlertTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// EnableAlertTemplateWithChan invokes the arms.EnableAlertTemplate API asynchronously
func (client *Client) EnableAlertTemplateWithChan(request *EnableAlertTemplateRequest) (<-chan *EnableAlertTemplateResponse, <-chan error) {
	responseChan := make(chan *EnableAlertTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableAlertTemplate(request)
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

// EnableAlertTemplateWithCallback invokes the arms.EnableAlertTemplate API asynchronously
func (client *Client) EnableAlertTemplateWithCallback(request *EnableAlertTemplateRequest, callback func(response *EnableAlertTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableAlertTemplateResponse
		var err error
		defer close(result)
		response, err = client.EnableAlertTemplate(request)
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

// EnableAlertTemplateRequest is the request struct for api EnableAlertTemplate
type EnableAlertTemplateRequest struct {
	*requests.RpcRequest
	Id          requests.Integer `position:"Query" name:"Id"`
	ProxyUserId string           `position:"Query" name:"ProxyUserId"`
}

// EnableAlertTemplateResponse is the response struct for api EnableAlertTemplate
type EnableAlertTemplateResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateEnableAlertTemplateRequest creates a request to invoke EnableAlertTemplate API
func CreateEnableAlertTemplateRequest() (request *EnableAlertTemplateRequest) {
	request = &EnableAlertTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2021-05-19", "EnableAlertTemplate", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableAlertTemplateResponse creates a response to parse from EnableAlertTemplate response
func CreateEnableAlertTemplateResponse() (response *EnableAlertTemplateResponse) {
	response = &EnableAlertTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
