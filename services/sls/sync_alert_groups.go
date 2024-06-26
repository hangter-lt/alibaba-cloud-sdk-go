package sls

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

// SyncAlertGroups invokes the sls.SyncAlertGroups API synchronously
func (client *Client) SyncAlertGroups(request *SyncAlertGroupsRequest) (response *SyncAlertGroupsResponse, err error) {
	response = CreateSyncAlertGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// SyncAlertGroupsWithChan invokes the sls.SyncAlertGroups API asynchronously
func (client *Client) SyncAlertGroupsWithChan(request *SyncAlertGroupsRequest) (<-chan *SyncAlertGroupsResponse, <-chan error) {
	responseChan := make(chan *SyncAlertGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SyncAlertGroups(request)
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

// SyncAlertGroupsWithCallback invokes the sls.SyncAlertGroups API asynchronously
func (client *Client) SyncAlertGroupsWithCallback(request *SyncAlertGroupsRequest, callback func(response *SyncAlertGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SyncAlertGroupsResponse
		var err error
		defer close(result)
		response, err = client.SyncAlertGroups(request)
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

// SyncAlertGroupsRequest is the request struct for api SyncAlertGroups
type SyncAlertGroupsRequest struct {
	*requests.RpcRequest
	App            string `position:"Body" name:"App"`
	Groups         string `position:"Body" name:"Groups"`
	SlsAccessToken string `position:"Query" name:"SlsAccessToken"`
}

// SyncAlertGroupsResponse is the response struct for api SyncAlertGroups
type SyncAlertGroupsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateSyncAlertGroupsRequest creates a request to invoke SyncAlertGroups API
func CreateSyncAlertGroupsRequest() (request *SyncAlertGroupsRequest) {
	request = &SyncAlertGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2019-10-23", "SyncAlertGroups", "sls", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSyncAlertGroupsResponse creates a response to parse from SyncAlertGroups response
func CreateSyncAlertGroupsResponse() (response *SyncAlertGroupsResponse) {
	response = &SyncAlertGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
