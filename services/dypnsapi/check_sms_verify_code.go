package dypnsapi

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

// CheckSmsVerifyCode invokes the dypnsapi.CheckSmsVerifyCode API synchronously
func (client *Client) CheckSmsVerifyCode(request *CheckSmsVerifyCodeRequest) (response *CheckSmsVerifyCodeResponse, err error) {
	response = CreateCheckSmsVerifyCodeResponse()
	err = client.DoAction(request, response)
	return
}

// CheckSmsVerifyCodeWithChan invokes the dypnsapi.CheckSmsVerifyCode API asynchronously
func (client *Client) CheckSmsVerifyCodeWithChan(request *CheckSmsVerifyCodeRequest) (<-chan *CheckSmsVerifyCodeResponse, <-chan error) {
	responseChan := make(chan *CheckSmsVerifyCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckSmsVerifyCode(request)
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

// CheckSmsVerifyCodeWithCallback invokes the dypnsapi.CheckSmsVerifyCode API asynchronously
func (client *Client) CheckSmsVerifyCodeWithCallback(request *CheckSmsVerifyCodeRequest, callback func(response *CheckSmsVerifyCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckSmsVerifyCodeResponse
		var err error
		defer close(result)
		response, err = client.CheckSmsVerifyCode(request)
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

// CheckSmsVerifyCodeRequest is the request struct for api CheckSmsVerifyCode
type CheckSmsVerifyCodeRequest struct {
	*requests.RpcRequest
	CaseAuthPolicy       requests.Integer `position:"Query" name:"CaseAuthPolicy"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CountryCode          string           `position:"Query" name:"CountryCode"`
	PhoneNumber          string           `position:"Query" name:"PhoneNumber"`
	ExtendFunction       string           `position:"Query" name:"ExtendFunction"`
	VerifyCode           string           `position:"Query" name:"VerifyCode"`
	RouteName            string           `position:"Query" name:"RouteName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SchemeName           string           `position:"Query" name:"SchemeName"`
	OutId                string           `position:"Query" name:"OutId"`
}

// CheckSmsVerifyCodeResponse is the response struct for api CheckSmsVerifyCode
type CheckSmsVerifyCodeResponse struct {
	*responses.BaseResponse
	AccessDeniedDetail string `json:"AccessDeniedDetail" xml:"AccessDeniedDetail"`
	Message            string `json:"Message" xml:"Message"`
	Code               string `json:"Code" xml:"Code"`
	Success            bool   `json:"Success" xml:"Success"`
	Model              Model  `json:"Model" xml:"Model"`
}

// CreateCheckSmsVerifyCodeRequest creates a request to invoke CheckSmsVerifyCode API
func CreateCheckSmsVerifyCodeRequest() (request *CheckSmsVerifyCodeRequest) {
	request = &CheckSmsVerifyCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "CheckSmsVerifyCode", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckSmsVerifyCodeResponse creates a response to parse from CheckSmsVerifyCode response
func CreateCheckSmsVerifyCodeResponse() (response *CheckSmsVerifyCodeResponse) {
	response = &CheckSmsVerifyCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
