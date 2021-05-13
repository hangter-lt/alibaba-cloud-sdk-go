package ens

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

// GetVmList invokes the ens.GetVmList API synchronously
func (client *Client) GetVmList(request *GetVmListRequest) (response *GetVmListResponse, err error) {
	response = CreateGetVmListResponse()
	err = client.DoAction(request, response)
	return
}

// GetVmListWithChan invokes the ens.GetVmList API asynchronously
func (client *Client) GetVmListWithChan(request *GetVmListRequest) (<-chan *GetVmListResponse, <-chan error) {
	responseChan := make(chan *GetVmListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVmList(request)
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

// GetVmListWithCallback invokes the ens.GetVmList API asynchronously
func (client *Client) GetVmListWithCallback(request *GetVmListRequest, callback func(response *GetVmListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVmListResponse
		var err error
		defer close(result)
		response, err = client.GetVmList(request)
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

// GetVmListRequest is the request struct for api GetVmList
type GetVmListRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query"`
	GroupUuid  string           `position:"Query"`
	PageSize   requests.Integer `position:"Query"`
	AliUid     requests.Integer `position:"Query"`
}

// GetVmListResponse is the response struct for api GetVmList
type GetVmListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Msg       string `json:"Msg" xml:"Msg"`
	Data      string `json:"Data" xml:"Data"`
	Desc      string `json:"Desc" xml:"Desc"`
}

// CreateGetVmListRequest creates a request to invoke GetVmList API
func CreateGetVmListRequest() (request *GetVmListRequest) {
	request = &GetVmListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "GetVmList", "ens", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetVmListResponse creates a response to parse from GetVmList response
func CreateGetVmListResponse() (response *GetVmListResponse) {
	response = &GetVmListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
