package iot

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

// EnableSceneRule invokes the iot.EnableSceneRule API synchronously
// api document: https://help.aliyun.com/api/iot/enablescenerule.html
func (client *Client) EnableSceneRule(request *EnableSceneRuleRequest) (response *EnableSceneRuleResponse, err error) {
	response = CreateEnableSceneRuleResponse()
	err = client.DoAction(request, response)
	return
}

// EnableSceneRuleWithChan invokes the iot.EnableSceneRule API asynchronously
// api document: https://help.aliyun.com/api/iot/enablescenerule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableSceneRuleWithChan(request *EnableSceneRuleRequest) (<-chan *EnableSceneRuleResponse, <-chan error) {
	responseChan := make(chan *EnableSceneRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableSceneRule(request)
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

// EnableSceneRuleWithCallback invokes the iot.EnableSceneRule API asynchronously
// api document: https://help.aliyun.com/api/iot/enablescenerule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) EnableSceneRuleWithCallback(request *EnableSceneRuleRequest, callback func(response *EnableSceneRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableSceneRuleResponse
		var err error
		defer close(result)
		response, err = client.EnableSceneRule(request)
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

// EnableSceneRuleRequest is the request struct for api EnableSceneRule
type EnableSceneRuleRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	RuleId        string `position:"Query" name:"RuleId"`
}

// EnableSceneRuleResponse is the response struct for api EnableSceneRule
type EnableSceneRuleResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
}

// CreateEnableSceneRuleRequest creates a request to invoke EnableSceneRule API
func CreateEnableSceneRuleRequest() (request *EnableSceneRuleRequest) {
	request = &EnableSceneRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "EnableSceneRule", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableSceneRuleResponse creates a response to parse from EnableSceneRule response
func CreateEnableSceneRuleResponse() (response *EnableSceneRuleResponse) {
	response = &EnableSceneRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}