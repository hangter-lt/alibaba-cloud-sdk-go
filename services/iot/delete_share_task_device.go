package iot

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

// DeleteShareTaskDevice invokes the iot.DeleteShareTaskDevice API synchronously
func (client *Client) DeleteShareTaskDevice(request *DeleteShareTaskDeviceRequest) (response *DeleteShareTaskDeviceResponse, err error) {
	response = CreateDeleteShareTaskDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteShareTaskDeviceWithChan invokes the iot.DeleteShareTaskDevice API asynchronously
func (client *Client) DeleteShareTaskDeviceWithChan(request *DeleteShareTaskDeviceRequest) (<-chan *DeleteShareTaskDeviceResponse, <-chan error) {
	responseChan := make(chan *DeleteShareTaskDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteShareTaskDevice(request)
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

// DeleteShareTaskDeviceWithCallback invokes the iot.DeleteShareTaskDevice API asynchronously
func (client *Client) DeleteShareTaskDeviceWithCallback(request *DeleteShareTaskDeviceRequest, callback func(response *DeleteShareTaskDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteShareTaskDeviceResponse
		var err error
		defer close(result)
		response, err = client.DeleteShareTaskDevice(request)
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

// DeleteShareTaskDeviceRequest is the request struct for api DeleteShareTaskDevice
type DeleteShareTaskDeviceRequest struct {
	*requests.RpcRequest
	IotInstanceId string    `position:"Body" name:"IotInstanceId"`
	IotIdList     *[]string `position:"Body" name:"IotIdList"  type:"Repeated"`
	ShareTaskId   string    `position:"Body" name:"ShareTaskId"`
	ApiProduct    string    `position:"Body" name:"ApiProduct"`
	ApiRevision   string    `position:"Body" name:"ApiRevision"`
}

// DeleteShareTaskDeviceResponse is the response struct for api DeleteShareTaskDevice
type DeleteShareTaskDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateDeleteShareTaskDeviceRequest creates a request to invoke DeleteShareTaskDevice API
func CreateDeleteShareTaskDeviceRequest() (request *DeleteShareTaskDeviceRequest) {
	request = &DeleteShareTaskDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteShareTaskDevice", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteShareTaskDeviceResponse creates a response to parse from DeleteShareTaskDevice response
func CreateDeleteShareTaskDeviceResponse() (response *DeleteShareTaskDeviceResponse) {
	response = &DeleteShareTaskDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
