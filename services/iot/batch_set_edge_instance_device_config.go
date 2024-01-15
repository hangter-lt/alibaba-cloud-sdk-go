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

// BatchSetEdgeInstanceDeviceConfig invokes the iot.BatchSetEdgeInstanceDeviceConfig API synchronously
func (client *Client) BatchSetEdgeInstanceDeviceConfig(request *BatchSetEdgeInstanceDeviceConfigRequest) (response *BatchSetEdgeInstanceDeviceConfigResponse, err error) {
	response = CreateBatchSetEdgeInstanceDeviceConfigResponse()
	err = client.DoAction(request, response)
	return
}

// BatchSetEdgeInstanceDeviceConfigWithChan invokes the iot.BatchSetEdgeInstanceDeviceConfig API asynchronously
func (client *Client) BatchSetEdgeInstanceDeviceConfigWithChan(request *BatchSetEdgeInstanceDeviceConfigRequest) (<-chan *BatchSetEdgeInstanceDeviceConfigResponse, <-chan error) {
	responseChan := make(chan *BatchSetEdgeInstanceDeviceConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchSetEdgeInstanceDeviceConfig(request)
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

// BatchSetEdgeInstanceDeviceConfigWithCallback invokes the iot.BatchSetEdgeInstanceDeviceConfig API asynchronously
func (client *Client) BatchSetEdgeInstanceDeviceConfigWithCallback(request *BatchSetEdgeInstanceDeviceConfigRequest, callback func(response *BatchSetEdgeInstanceDeviceConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchSetEdgeInstanceDeviceConfigResponse
		var err error
		defer close(result)
		response, err = client.BatchSetEdgeInstanceDeviceConfig(request)
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

// BatchSetEdgeInstanceDeviceConfigRequest is the request struct for api BatchSetEdgeInstanceDeviceConfig
type BatchSetEdgeInstanceDeviceConfigRequest struct {
	*requests.RpcRequest
	DeviceConfigs *[]BatchSetEdgeInstanceDeviceConfigDeviceConfigs `position:"Query" name:"DeviceConfigs"  type:"Repeated"`
	IotInstanceId string                                           `position:"Query" name:"IotInstanceId"`
	InstanceId    string                                           `position:"Query" name:"InstanceId"`
	ApiProduct    string                                           `position:"Body" name:"ApiProduct"`
	ApiRevision   string                                           `position:"Body" name:"ApiRevision"`
}

// BatchSetEdgeInstanceDeviceConfigDeviceConfigs is a repeated param struct in BatchSetEdgeInstanceDeviceConfigRequest
type BatchSetEdgeInstanceDeviceConfigDeviceConfigs struct {
	IotId   string `name:"IotId"`
	Content string `name:"Content"`
}

// BatchSetEdgeInstanceDeviceConfigResponse is the response struct for api BatchSetEdgeInstanceDeviceConfig
type BatchSetEdgeInstanceDeviceConfigResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateBatchSetEdgeInstanceDeviceConfigRequest creates a request to invoke BatchSetEdgeInstanceDeviceConfig API
func CreateBatchSetEdgeInstanceDeviceConfigRequest() (request *BatchSetEdgeInstanceDeviceConfigRequest) {
	request = &BatchSetEdgeInstanceDeviceConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchSetEdgeInstanceDeviceConfig", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchSetEdgeInstanceDeviceConfigResponse creates a response to parse from BatchSetEdgeInstanceDeviceConfig response
func CreateBatchSetEdgeInstanceDeviceConfigResponse() (response *BatchSetEdgeInstanceDeviceConfigResponse) {
	response = &BatchSetEdgeInstanceDeviceConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
