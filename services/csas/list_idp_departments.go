package csas

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

// ListIdpDepartments invokes the csas.ListIdpDepartments API synchronously
func (client *Client) ListIdpDepartments(request *ListIdpDepartmentsRequest) (response *ListIdpDepartmentsResponse, err error) {
	response = CreateListIdpDepartmentsResponse()
	err = client.DoAction(request, response)
	return
}

// ListIdpDepartmentsWithChan invokes the csas.ListIdpDepartments API asynchronously
func (client *Client) ListIdpDepartmentsWithChan(request *ListIdpDepartmentsRequest) (<-chan *ListIdpDepartmentsResponse, <-chan error) {
	responseChan := make(chan *ListIdpDepartmentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIdpDepartments(request)
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

// ListIdpDepartmentsWithCallback invokes the csas.ListIdpDepartments API asynchronously
func (client *Client) ListIdpDepartmentsWithCallback(request *ListIdpDepartmentsRequest, callback func(response *ListIdpDepartmentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIdpDepartmentsResponse
		var err error
		defer close(result)
		response, err = client.ListIdpDepartments(request)
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

// ListIdpDepartmentsRequest is the request struct for api ListIdpDepartments
type ListIdpDepartmentsRequest struct {
	*requests.RpcRequest
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	IdpConfigId string           `position:"Query" name:"IdpConfigId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
}

// ListIdpDepartmentsResponse is the response struct for api ListIdpDepartments
type ListIdpDepartmentsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListIdpDepartmentsRequest creates a request to invoke ListIdpDepartments API
func CreateListIdpDepartmentsRequest() (request *ListIdpDepartmentsRequest) {
	request = &ListIdpDepartmentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "ListIdpDepartments", "", "")
	request.Method = requests.GET
	return
}

// CreateListIdpDepartmentsResponse creates a response to parse from ListIdpDepartments response
func CreateListIdpDepartmentsResponse() (response *ListIdpDepartmentsResponse) {
	response = &ListIdpDepartmentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
