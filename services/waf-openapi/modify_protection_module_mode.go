package waf_openapi

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

// ModifyProtectionModuleMode invokes the waf_openapi.ModifyProtectionModuleMode API synchronously
func (client *Client) ModifyProtectionModuleMode(request *ModifyProtectionModuleModeRequest) (response *ModifyProtectionModuleModeResponse, err error) {
	response = CreateModifyProtectionModuleModeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyProtectionModuleModeWithChan invokes the waf_openapi.ModifyProtectionModuleMode API asynchronously
func (client *Client) ModifyProtectionModuleModeWithChan(request *ModifyProtectionModuleModeRequest) (<-chan *ModifyProtectionModuleModeResponse, <-chan error) {
	responseChan := make(chan *ModifyProtectionModuleModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyProtectionModuleMode(request)
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

// ModifyProtectionModuleModeWithCallback invokes the waf_openapi.ModifyProtectionModuleMode API asynchronously
func (client *Client) ModifyProtectionModuleModeWithCallback(request *ModifyProtectionModuleModeRequest, callback func(response *ModifyProtectionModuleModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyProtectionModuleModeResponse
		var err error
		defer close(result)
		response, err = client.ModifyProtectionModuleMode(request)
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

// ModifyProtectionModuleModeRequest is the request struct for api ModifyProtectionModuleMode
type ModifyProtectionModuleModeRequest struct {
	*requests.RpcRequest
	Mode            requests.Integer `position:"Query" name:"Mode"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Lang            string           `position:"Query" name:"Lang"`
	DefenseType     string           `position:"Query" name:"DefenseType"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	Domain          string           `position:"Query" name:"Domain"`
}

// ModifyProtectionModuleModeResponse is the response struct for api ModifyProtectionModuleMode
type ModifyProtectionModuleModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyProtectionModuleModeRequest creates a request to invoke ModifyProtectionModuleMode API
func CreateModifyProtectionModuleModeRequest() (request *ModifyProtectionModuleModeRequest) {
	request = &ModifyProtectionModuleModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "ModifyProtectionModuleMode", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyProtectionModuleModeResponse creates a response to parse from ModifyProtectionModuleMode response
func CreateModifyProtectionModuleModeResponse() (response *ModifyProtectionModuleModeResponse) {
	response = &ModifyProtectionModuleModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
