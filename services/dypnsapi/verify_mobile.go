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

// VerifyMobile invokes the dypnsapi.VerifyMobile API synchronously
func (client *Client) VerifyMobile(request *VerifyMobileRequest) (response *VerifyMobileResponse, err error) {
	response = CreateVerifyMobileResponse()
	err = client.DoAction(request, response)
	return
}

// VerifyMobileWithChan invokes the dypnsapi.VerifyMobile API asynchronously
func (client *Client) VerifyMobileWithChan(request *VerifyMobileRequest) (<-chan *VerifyMobileResponse, <-chan error) {
	responseChan := make(chan *VerifyMobileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VerifyMobile(request)
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

// VerifyMobileWithCallback invokes the dypnsapi.VerifyMobile API asynchronously
func (client *Client) VerifyMobileWithCallback(request *VerifyMobileRequest, callback func(response *VerifyMobileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VerifyMobileResponse
		var err error
		defer close(result)
		response, err = client.VerifyMobile(request)
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

// VerifyMobileRequest is the request struct for api VerifyMobile
type VerifyMobileRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PhoneNumber          string           `position:"Query" name:"PhoneNumber"`
	AccessCode           string           `position:"Query" name:"AccessCode"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OutId                string           `position:"Query" name:"OutId"`
}

// VerifyMobileResponse is the response struct for api VerifyMobile
type VerifyMobileResponse struct {
	*responses.BaseResponse
	Code                string              `json:"Code" xml:"Code"`
	Message             string              `json:"Message" xml:"Message"`
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	GateVerifyResultDTO GateVerifyResultDTO `json:"GateVerifyResultDTO" xml:"GateVerifyResultDTO"`
}

// CreateVerifyMobileRequest creates a request to invoke VerifyMobile API
func CreateVerifyMobileRequest() (request *VerifyMobileRequest) {
	request = &VerifyMobileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "VerifyMobile", "", "")
	request.Method = requests.POST
	return
}

// CreateVerifyMobileResponse creates a response to parse from VerifyMobile response
func CreateVerifyMobileResponse() (response *VerifyMobileResponse) {
	response = &VerifyMobileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
