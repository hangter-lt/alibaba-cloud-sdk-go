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

// ImportCustomAlertRules invokes the arms.ImportCustomAlertRules API synchronously
func (client *Client) ImportCustomAlertRules(request *ImportCustomAlertRulesRequest) (response *ImportCustomAlertRulesResponse, err error) {
	response = CreateImportCustomAlertRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ImportCustomAlertRulesWithChan invokes the arms.ImportCustomAlertRules API asynchronously
func (client *Client) ImportCustomAlertRulesWithChan(request *ImportCustomAlertRulesRequest) (<-chan *ImportCustomAlertRulesResponse, <-chan error) {
	responseChan := make(chan *ImportCustomAlertRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ImportCustomAlertRules(request)
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

// ImportCustomAlertRulesWithCallback invokes the arms.ImportCustomAlertRules API asynchronously
func (client *Client) ImportCustomAlertRulesWithCallback(request *ImportCustomAlertRulesRequest, callback func(response *ImportCustomAlertRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ImportCustomAlertRulesResponse
		var err error
		defer close(result)
		response, err = client.ImportCustomAlertRules(request)
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

// ImportCustomAlertRulesRequest is the request struct for api ImportCustomAlertRules
type ImportCustomAlertRulesRequest struct {
	*requests.RpcRequest
	IsAutoStart         requests.Boolean `position:"Query" name:"IsAutoStart"`
	ProxyUserId         string           `position:"Query" name:"ProxyUserId"`
	ContactGroupIds     string           `position:"Query" name:"ContactGroupIds"`
	TemplateAlertConfig string           `position:"Query" name:"TemplateAlertConfig"`
	TemplageAlertConfig string           `position:"Query" name:"TemplageAlertConfig"`
}

// ImportCustomAlertRulesResponse is the response struct for api ImportCustomAlertRules
type ImportCustomAlertRulesResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateImportCustomAlertRulesRequest creates a request to invoke ImportCustomAlertRules API
func CreateImportCustomAlertRulesRequest() (request *ImportCustomAlertRulesRequest) {
	request = &ImportCustomAlertRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2021-05-19", "ImportCustomAlertRules", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateImportCustomAlertRulesResponse creates a response to parse from ImportCustomAlertRules response
func CreateImportCustomAlertRulesResponse() (response *ImportCustomAlertRulesResponse) {
	response = &ImportCustomAlertRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
