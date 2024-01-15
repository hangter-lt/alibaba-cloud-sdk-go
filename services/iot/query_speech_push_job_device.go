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

// QuerySpeechPushJobDevice invokes the iot.QuerySpeechPushJobDevice API synchronously
func (client *Client) QuerySpeechPushJobDevice(request *QuerySpeechPushJobDeviceRequest) (response *QuerySpeechPushJobDeviceResponse, err error) {
	response = CreateQuerySpeechPushJobDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySpeechPushJobDeviceWithChan invokes the iot.QuerySpeechPushJobDevice API asynchronously
func (client *Client) QuerySpeechPushJobDeviceWithChan(request *QuerySpeechPushJobDeviceRequest) (<-chan *QuerySpeechPushJobDeviceResponse, <-chan error) {
	responseChan := make(chan *QuerySpeechPushJobDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySpeechPushJobDevice(request)
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

// QuerySpeechPushJobDeviceWithCallback invokes the iot.QuerySpeechPushJobDevice API asynchronously
func (client *Client) QuerySpeechPushJobDeviceWithCallback(request *QuerySpeechPushJobDeviceRequest, callback func(response *QuerySpeechPushJobDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySpeechPushJobDeviceResponse
		var err error
		defer close(result)
		response, err = client.QuerySpeechPushJobDevice(request)
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

// QuerySpeechPushJobDeviceRequest is the request struct for api QuerySpeechPushJobDevice
type QuerySpeechPushJobDeviceRequest struct {
	*requests.RpcRequest
	PageId        requests.Integer `position:"Body" name:"PageId"`
	IotInstanceId string           `position:"Body" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	JobCode       string           `position:"Body" name:"JobCode"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Body" name:"DeviceName"`
	Status        string           `position:"Body" name:"Status"`
}

// QuerySpeechPushJobDeviceResponse is the response struct for api QuerySpeechPushJobDevice
type QuerySpeechPushJobDeviceResponse struct {
	*responses.BaseResponse
	RequestId    string                         `json:"RequestId" xml:"RequestId"`
	Success      bool                           `json:"Success" xml:"Success"`
	Code         string                         `json:"Code" xml:"Code"`
	ErrorMessage string                         `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQuerySpeechPushJobDevice `json:"Data" xml:"Data"`
}

// CreateQuerySpeechPushJobDeviceRequest creates a request to invoke QuerySpeechPushJobDevice API
func CreateQuerySpeechPushJobDeviceRequest() (request *QuerySpeechPushJobDeviceRequest) {
	request = &QuerySpeechPushJobDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QuerySpeechPushJobDevice", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySpeechPushJobDeviceResponse creates a response to parse from QuerySpeechPushJobDevice response
func CreateQuerySpeechPushJobDeviceResponse() (response *QuerySpeechPushJobDeviceResponse) {
	response = &QuerySpeechPushJobDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
