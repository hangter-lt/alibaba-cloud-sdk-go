package idaas_doraemon

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

// ListAuthenticatorOpsLogs invokes the idaas_doraemon.ListAuthenticatorOpsLogs API synchronously
func (client *Client) ListAuthenticatorOpsLogs(request *ListAuthenticatorOpsLogsRequest) (response *ListAuthenticatorOpsLogsResponse, err error) {
	response = CreateListAuthenticatorOpsLogsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAuthenticatorOpsLogsWithChan invokes the idaas_doraemon.ListAuthenticatorOpsLogs API asynchronously
func (client *Client) ListAuthenticatorOpsLogsWithChan(request *ListAuthenticatorOpsLogsRequest) (<-chan *ListAuthenticatorOpsLogsResponse, <-chan error) {
	responseChan := make(chan *ListAuthenticatorOpsLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAuthenticatorOpsLogs(request)
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

// ListAuthenticatorOpsLogsWithCallback invokes the idaas_doraemon.ListAuthenticatorOpsLogs API asynchronously
func (client *Client) ListAuthenticatorOpsLogsWithCallback(request *ListAuthenticatorOpsLogsRequest, callback func(response *ListAuthenticatorOpsLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAuthenticatorOpsLogsResponse
		var err error
		defer close(result)
		response, err = client.ListAuthenticatorOpsLogs(request)
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

// ListAuthenticatorOpsLogsRequest is the request struct for api ListAuthenticatorOpsLogs
type ListAuthenticatorOpsLogsRequest struct {
	*requests.RpcRequest
	ToTime                requests.Integer `position:"Query" name:"ToTime"`
	UserId                string           `position:"Query" name:"UserId"`
	PageNumber            requests.Integer `position:"Query" name:"PageNumber"`
	PageSize              requests.Integer `position:"Query" name:"PageSize"`
	FromTime              requests.Integer `position:"Query" name:"FromTime"`
	AuthenticatorUuid     string           `position:"Query" name:"AuthenticatorUuid"`
	AuthenticatorType     string           `position:"Query" name:"AuthenticatorType"`
	ApplicationExternalId string           `position:"Query" name:"ApplicationExternalId"`
}

// ListAuthenticatorOpsLogsResponse is the response struct for api ListAuthenticatorOpsLogs
type ListAuthenticatorOpsLogsResponse struct {
	*responses.BaseResponse
	RequestId                string                         `json:"RequestId" xml:"RequestId"`
	TotalCount               int64                          `json:"TotalCount" xml:"TotalCount"`
	PageNumber               int64                          `json:"PageNumber" xml:"PageNumber"`
	PageSize                 int64                          `json:"PageSize" xml:"PageSize"`
	AuthenticationLogContent []AuthenticationLogContentItem `json:"AuthenticationLogContent" xml:"AuthenticationLogContent"`
}

// CreateListAuthenticatorOpsLogsRequest creates a request to invoke ListAuthenticatorOpsLogs API
func CreateListAuthenticatorOpsLogsRequest() (request *ListAuthenticatorOpsLogsRequest) {
	request = &ListAuthenticatorOpsLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idaas-doraemon", "2021-05-20", "ListAuthenticatorOpsLogs", "idaasauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAuthenticatorOpsLogsResponse creates a response to parse from ListAuthenticatorOpsLogs response
func CreateListAuthenticatorOpsLogsResponse() (response *ListAuthenticatorOpsLogsResponse) {
	response = &ListAuthenticatorOpsLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
