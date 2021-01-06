package cloudgameapi

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

// GetSession invokes the cloudgameapi.GetSession API synchronously
func (client *Client) GetSession(request *GetSessionRequest) (response *GetSessionResponse, err error) {
	response = CreateGetSessionResponse()
	err = client.DoAction(request, response)
	return
}

// GetSessionWithChan invokes the cloudgameapi.GetSession API asynchronously
func (client *Client) GetSessionWithChan(request *GetSessionRequest) (<-chan *GetSessionResponse, <-chan error) {
	responseChan := make(chan *GetSessionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSession(request)
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

// GetSessionWithCallback invokes the cloudgameapi.GetSession API asynchronously
func (client *Client) GetSessionWithCallback(request *GetSessionRequest, callback func(response *GetSessionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSessionResponse
		var err error
		defer close(result)
		response, err = client.GetSession(request)
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

// GetSessionRequest is the request struct for api GetSession
type GetSessionRequest struct {
	*requests.RpcRequest
	Token string `position:"Query" name:"Token"`
}

// GetSessionResponse is the response struct for api GetSession
type GetSessionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetSessionRequest creates a request to invoke GetSession API
func CreateGetSessionRequest() (request *GetSessionRequest) {
	request = &GetSessionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudGameAPI", "2020-07-28", "GetSession", "", "")
	request.Method = requests.POST
	return
}

// CreateGetSessionResponse creates a response to parse from GetSession response
func CreateGetSessionResponse() (response *GetSessionResponse) {
	response = &GetSessionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}