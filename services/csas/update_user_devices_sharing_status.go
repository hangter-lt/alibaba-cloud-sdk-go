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

// UpdateUserDevicesSharingStatus invokes the csas.UpdateUserDevicesSharingStatus API synchronously
func (client *Client) UpdateUserDevicesSharingStatus(request *UpdateUserDevicesSharingStatusRequest) (response *UpdateUserDevicesSharingStatusResponse, err error) {
	response = CreateUpdateUserDevicesSharingStatusResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUserDevicesSharingStatusWithChan invokes the csas.UpdateUserDevicesSharingStatus API asynchronously
func (client *Client) UpdateUserDevicesSharingStatusWithChan(request *UpdateUserDevicesSharingStatusRequest) (<-chan *UpdateUserDevicesSharingStatusResponse, <-chan error) {
	responseChan := make(chan *UpdateUserDevicesSharingStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUserDevicesSharingStatus(request)
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

// UpdateUserDevicesSharingStatusWithCallback invokes the csas.UpdateUserDevicesSharingStatus API asynchronously
func (client *Client) UpdateUserDevicesSharingStatusWithCallback(request *UpdateUserDevicesSharingStatusRequest, callback func(response *UpdateUserDevicesSharingStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUserDevicesSharingStatusResponse
		var err error
		defer close(result)
		response, err = client.UpdateUserDevicesSharingStatus(request)
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

// UpdateUserDevicesSharingStatusRequest is the request struct for api UpdateUserDevicesSharingStatus
type UpdateUserDevicesSharingStatusRequest struct {
	*requests.RpcRequest
	SharingStatus requests.Boolean `position:"Body" name:"SharingStatus"`
	DeviceTags    *[]string        `position:"Body" name:"DeviceTags"  type:"Repeated"`
	SourceIp      string           `position:"Query" name:"SourceIp"`
}

// UpdateUserDevicesSharingStatusResponse is the response struct for api UpdateUserDevicesSharingStatus
type UpdateUserDevicesSharingStatusResponse struct {
	*responses.BaseResponse
	RequestId string                                 `json:"RequestId" xml:"RequestId"`
	Devices   []DataInUpdateUserDevicesSharingStatus `json:"Devices" xml:"Devices"`
}

// CreateUpdateUserDevicesSharingStatusRequest creates a request to invoke UpdateUserDevicesSharingStatus API
func CreateUpdateUserDevicesSharingStatusRequest() (request *UpdateUserDevicesSharingStatusRequest) {
	request = &UpdateUserDevicesSharingStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "UpdateUserDevicesSharingStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateUserDevicesSharingStatusResponse creates a response to parse from UpdateUserDevicesSharingStatus response
func CreateUpdateUserDevicesSharingStatusResponse() (response *UpdateUserDevicesSharingStatusResponse) {
	response = &UpdateUserDevicesSharingStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
