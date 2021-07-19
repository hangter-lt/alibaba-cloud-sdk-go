package das

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

// GetRunningSqlConcurrencyControlRules invokes the das.GetRunningSqlConcurrencyControlRules API synchronously
func (client *Client) GetRunningSqlConcurrencyControlRules(request *GetRunningSqlConcurrencyControlRulesRequest) (response *GetRunningSqlConcurrencyControlRulesResponse, err error) {
	response = CreateGetRunningSqlConcurrencyControlRulesResponse()
	err = client.DoAction(request, response)
	return
}

// GetRunningSqlConcurrencyControlRulesWithChan invokes the das.GetRunningSqlConcurrencyControlRules API asynchronously
func (client *Client) GetRunningSqlConcurrencyControlRulesWithChan(request *GetRunningSqlConcurrencyControlRulesRequest) (<-chan *GetRunningSqlConcurrencyControlRulesResponse, <-chan error) {
	responseChan := make(chan *GetRunningSqlConcurrencyControlRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRunningSqlConcurrencyControlRules(request)
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

// GetRunningSqlConcurrencyControlRulesWithCallback invokes the das.GetRunningSqlConcurrencyControlRules API asynchronously
func (client *Client) GetRunningSqlConcurrencyControlRulesWithCallback(request *GetRunningSqlConcurrencyControlRulesRequest, callback func(response *GetRunningSqlConcurrencyControlRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRunningSqlConcurrencyControlRulesResponse
		var err error
		defer close(result)
		response, err = client.GetRunningSqlConcurrencyControlRules(request)
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

// GetRunningSqlConcurrencyControlRulesRequest is the request struct for api GetRunningSqlConcurrencyControlRules
type GetRunningSqlConcurrencyControlRulesRequest struct {
	*requests.RpcRequest
	ConsoleContext string           `position:"Query" name:"ConsoleContext"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	PageNo         requests.Integer `position:"Query" name:"PageNo"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// GetRunningSqlConcurrencyControlRulesResponse is the response struct for api GetRunningSqlConcurrencyControlRules
type GetRunningSqlConcurrencyControlRulesResponse struct {
	*responses.BaseResponse
	Code      string                                     `json:"Code" xml:"Code"`
	Message   string                                     `json:"Message" xml:"Message"`
	RequestId string                                     `json:"RequestId" xml:"RequestId"`
	Success   string                                     `json:"Success" xml:"Success"`
	Data      DataInGetRunningSqlConcurrencyControlRules `json:"Data" xml:"Data"`
}

// CreateGetRunningSqlConcurrencyControlRulesRequest creates a request to invoke GetRunningSqlConcurrencyControlRules API
func CreateGetRunningSqlConcurrencyControlRulesRequest() (request *GetRunningSqlConcurrencyControlRulesRequest) {
	request = &GetRunningSqlConcurrencyControlRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetRunningSqlConcurrencyControlRules", "das", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetRunningSqlConcurrencyControlRulesResponse creates a response to parse from GetRunningSqlConcurrencyControlRules response
func CreateGetRunningSqlConcurrencyControlRulesResponse() (response *GetRunningSqlConcurrencyControlRulesResponse) {
	response = &GetRunningSqlConcurrencyControlRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}