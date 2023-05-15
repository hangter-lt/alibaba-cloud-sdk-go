package dms_enterprise

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

// SkipDataCorrectRowCheck invokes the dms_enterprise.SkipDataCorrectRowCheck API synchronously
func (client *Client) SkipDataCorrectRowCheck(request *SkipDataCorrectRowCheckRequest) (response *SkipDataCorrectRowCheckResponse, err error) {
	response = CreateSkipDataCorrectRowCheckResponse()
	err = client.DoAction(request, response)
	return
}

// SkipDataCorrectRowCheckWithChan invokes the dms_enterprise.SkipDataCorrectRowCheck API asynchronously
func (client *Client) SkipDataCorrectRowCheckWithChan(request *SkipDataCorrectRowCheckRequest) (<-chan *SkipDataCorrectRowCheckResponse, <-chan error) {
	responseChan := make(chan *SkipDataCorrectRowCheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SkipDataCorrectRowCheck(request)
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

// SkipDataCorrectRowCheckWithCallback invokes the dms_enterprise.SkipDataCorrectRowCheck API asynchronously
func (client *Client) SkipDataCorrectRowCheckWithCallback(request *SkipDataCorrectRowCheckRequest, callback func(response *SkipDataCorrectRowCheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SkipDataCorrectRowCheckResponse
		var err error
		defer close(result)
		response, err = client.SkipDataCorrectRowCheck(request)
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

// SkipDataCorrectRowCheckRequest is the request struct for api SkipDataCorrectRowCheck
type SkipDataCorrectRowCheckRequest struct {
	*requests.RpcRequest
	Reason  string           `position:"Query" name:"Reason"`
	Tid     requests.Integer `position:"Query" name:"Tid"`
	OrderId requests.Integer `position:"Query" name:"OrderId"`
}

// SkipDataCorrectRowCheckResponse is the response struct for api SkipDataCorrectRowCheck
type SkipDataCorrectRowCheckResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateSkipDataCorrectRowCheckRequest creates a request to invoke SkipDataCorrectRowCheck API
func CreateSkipDataCorrectRowCheckRequest() (request *SkipDataCorrectRowCheckRequest) {
	request = &SkipDataCorrectRowCheckRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "SkipDataCorrectRowCheck", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSkipDataCorrectRowCheckResponse creates a response to parse from SkipDataCorrectRowCheck response
func CreateSkipDataCorrectRowCheckResponse() (response *SkipDataCorrectRowCheckResponse) {
	response = &SkipDataCorrectRowCheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
