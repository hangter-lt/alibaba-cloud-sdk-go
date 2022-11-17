package edas

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

// DeleteApplicationScalingRule invokes the edas.DeleteApplicationScalingRule API synchronously
func (client *Client) DeleteApplicationScalingRule(request *DeleteApplicationScalingRuleRequest) (response *DeleteApplicationScalingRuleResponse, err error) {
	response = CreateDeleteApplicationScalingRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteApplicationScalingRuleWithChan invokes the edas.DeleteApplicationScalingRule API asynchronously
func (client *Client) DeleteApplicationScalingRuleWithChan(request *DeleteApplicationScalingRuleRequest) (<-chan *DeleteApplicationScalingRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteApplicationScalingRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteApplicationScalingRule(request)
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

// DeleteApplicationScalingRuleWithCallback invokes the edas.DeleteApplicationScalingRule API asynchronously
func (client *Client) DeleteApplicationScalingRuleWithCallback(request *DeleteApplicationScalingRuleRequest, callback func(response *DeleteApplicationScalingRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteApplicationScalingRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteApplicationScalingRule(request)
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

// DeleteApplicationScalingRuleRequest is the request struct for api DeleteApplicationScalingRule
type DeleteApplicationScalingRuleRequest struct {
	*requests.RoaRequest
	ScalingRuleName string `position:"Query" name:"ScalingRuleName"`
	AppId           string `position:"Query" name:"AppId"`
}

// DeleteApplicationScalingRuleResponse is the response struct for api DeleteApplicationScalingRule
type DeleteApplicationScalingRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteApplicationScalingRuleRequest creates a request to invoke DeleteApplicationScalingRule API
func CreateDeleteApplicationScalingRuleRequest() (request *DeleteApplicationScalingRuleRequest) {
	request = &DeleteApplicationScalingRuleRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "DeleteApplicationScalingRule", "/pop/v1/eam/scale/application_scaling_rule", "Edas", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateDeleteApplicationScalingRuleResponse creates a response to parse from DeleteApplicationScalingRule response
func CreateDeleteApplicationScalingRuleResponse() (response *DeleteApplicationScalingRuleResponse) {
	response = &DeleteApplicationScalingRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
