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

// ListExcessiveDeviceRegistrationApplications invokes the csas.ListExcessiveDeviceRegistrationApplications API synchronously
func (client *Client) ListExcessiveDeviceRegistrationApplications(request *ListExcessiveDeviceRegistrationApplicationsRequest) (response *ListExcessiveDeviceRegistrationApplicationsResponse, err error) {
	response = CreateListExcessiveDeviceRegistrationApplicationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListExcessiveDeviceRegistrationApplicationsWithChan invokes the csas.ListExcessiveDeviceRegistrationApplications API asynchronously
func (client *Client) ListExcessiveDeviceRegistrationApplicationsWithChan(request *ListExcessiveDeviceRegistrationApplicationsRequest) (<-chan *ListExcessiveDeviceRegistrationApplicationsResponse, <-chan error) {
	responseChan := make(chan *ListExcessiveDeviceRegistrationApplicationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListExcessiveDeviceRegistrationApplications(request)
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

// ListExcessiveDeviceRegistrationApplicationsWithCallback invokes the csas.ListExcessiveDeviceRegistrationApplications API asynchronously
func (client *Client) ListExcessiveDeviceRegistrationApplicationsWithCallback(request *ListExcessiveDeviceRegistrationApplicationsRequest, callback func(response *ListExcessiveDeviceRegistrationApplicationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListExcessiveDeviceRegistrationApplicationsResponse
		var err error
		defer close(result)
		response, err = client.ListExcessiveDeviceRegistrationApplications(request)
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

// ListExcessiveDeviceRegistrationApplicationsRequest is the request struct for api ListExcessiveDeviceRegistrationApplications
type ListExcessiveDeviceRegistrationApplicationsRequest struct {
	*requests.RpcRequest
	DeviceTag      string           `position:"Query" name:"DeviceTag"`
	Mac            string           `position:"Query" name:"Mac"`
	Hostname       string           `position:"Query" name:"Hostname"`
	SourceIp       string           `position:"Query" name:"SourceIp"`
	SaseUserId     string           `position:"Query" name:"SaseUserId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	Department     string           `position:"Query" name:"Department"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	ApplicationIds *[]string        `position:"Query" name:"ApplicationIds"  type:"Repeated"`
	Statuses       *[]string        `position:"Query" name:"Statuses"  type:"Repeated"`
	Username       string           `position:"Query" name:"Username"`
}

// ListExcessiveDeviceRegistrationApplicationsResponse is the response struct for api ListExcessiveDeviceRegistrationApplications
type ListExcessiveDeviceRegistrationApplicationsResponse struct {
	*responses.BaseResponse
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	TotalNum     int64      `json:"TotalNum" xml:"TotalNum"`
	Applications []DataList `json:"Applications" xml:"Applications"`
}

// CreateListExcessiveDeviceRegistrationApplicationsRequest creates a request to invoke ListExcessiveDeviceRegistrationApplications API
func CreateListExcessiveDeviceRegistrationApplicationsRequest() (request *ListExcessiveDeviceRegistrationApplicationsRequest) {
	request = &ListExcessiveDeviceRegistrationApplicationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "ListExcessiveDeviceRegistrationApplications", "", "")
	request.Method = requests.GET
	return
}

// CreateListExcessiveDeviceRegistrationApplicationsResponse creates a response to parse from ListExcessiveDeviceRegistrationApplications response
func CreateListExcessiveDeviceRegistrationApplicationsResponse() (response *ListExcessiveDeviceRegistrationApplicationsResponse) {
	response = &ListExcessiveDeviceRegistrationApplicationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
