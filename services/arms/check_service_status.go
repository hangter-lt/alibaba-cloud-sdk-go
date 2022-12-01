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

// CheckServiceStatus invokes the arms.CheckServiceStatus API synchronously
func (client *Client) CheckServiceStatus(request *CheckServiceStatusRequest) (response *CheckServiceStatusResponse, err error) {
	response = CreateCheckServiceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckServiceStatusWithChan invokes the arms.CheckServiceStatus API asynchronously
func (client *Client) CheckServiceStatusWithChan(request *CheckServiceStatusRequest) (<-chan *CheckServiceStatusResponse, <-chan error) {
	responseChan := make(chan *CheckServiceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckServiceStatus(request)
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

// CheckServiceStatusWithCallback invokes the arms.CheckServiceStatus API asynchronously
func (client *Client) CheckServiceStatusWithCallback(request *CheckServiceStatusRequest, callback func(response *CheckServiceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckServiceStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckServiceStatus(request)
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

// CheckServiceStatusRequest is the request struct for api CheckServiceStatus
type CheckServiceStatusRequest struct {
	*requests.RpcRequest
	SvcCode string `position:"Query" name:"SvcCode"`
}

// CheckServiceStatusResponse is the response struct for api CheckServiceStatus
type CheckServiceStatusResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCheckServiceStatusRequest creates a request to invoke CheckServiceStatus API
func CreateCheckServiceStatusRequest() (request *CheckServiceStatusRequest) {
	request = &CheckServiceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2021-05-19", "CheckServiceStatus", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckServiceStatusResponse creates a response to parse from CheckServiceStatus response
func CreateCheckServiceStatusResponse() (response *CheckServiceStatusResponse) {
	response = &CheckServiceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
