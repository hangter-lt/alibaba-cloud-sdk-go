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

// UpdateEdgeInstanceMessageRouting invokes the iot.UpdateEdgeInstanceMessageRouting API synchronously
func (client *Client) UpdateEdgeInstanceMessageRouting(request *UpdateEdgeInstanceMessageRoutingRequest) (response *UpdateEdgeInstanceMessageRoutingResponse, err error) {
	response = CreateUpdateEdgeInstanceMessageRoutingResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateEdgeInstanceMessageRoutingWithChan invokes the iot.UpdateEdgeInstanceMessageRouting API asynchronously
func (client *Client) UpdateEdgeInstanceMessageRoutingWithChan(request *UpdateEdgeInstanceMessageRoutingRequest) (<-chan *UpdateEdgeInstanceMessageRoutingResponse, <-chan error) {
	responseChan := make(chan *UpdateEdgeInstanceMessageRoutingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateEdgeInstanceMessageRouting(request)
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

// UpdateEdgeInstanceMessageRoutingWithCallback invokes the iot.UpdateEdgeInstanceMessageRouting API asynchronously
func (client *Client) UpdateEdgeInstanceMessageRoutingWithCallback(request *UpdateEdgeInstanceMessageRoutingRequest, callback func(response *UpdateEdgeInstanceMessageRoutingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateEdgeInstanceMessageRoutingResponse
		var err error
		defer close(result)
		response, err = client.UpdateEdgeInstanceMessageRouting(request)
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

// UpdateEdgeInstanceMessageRoutingRequest is the request struct for api UpdateEdgeInstanceMessageRouting
type UpdateEdgeInstanceMessageRoutingRequest struct {
	*requests.RpcRequest
	SourceData      string           `position:"Query" name:"SourceData"`
	TargetType      string           `position:"Query" name:"TargetType"`
	IotInstanceId   string           `position:"Query" name:"IotInstanceId"`
	SourceType      string           `position:"Query" name:"SourceType"`
	TopicFilter     string           `position:"Query" name:"TopicFilter"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
	RouteId         requests.Integer `position:"Query" name:"RouteId"`
	TargetData      string           `position:"Query" name:"TargetData"`
	ApiProduct      string           `position:"Body" name:"ApiProduct"`
	Name            string           `position:"Query" name:"Name"`
	ApiRevision     string           `position:"Body" name:"ApiRevision"`
	TargetIotHubQos requests.Integer `position:"Query" name:"TargetIotHubQos"`
}

// UpdateEdgeInstanceMessageRoutingResponse is the response struct for api UpdateEdgeInstanceMessageRouting
type UpdateEdgeInstanceMessageRoutingResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateEdgeInstanceMessageRoutingRequest creates a request to invoke UpdateEdgeInstanceMessageRouting API
func CreateUpdateEdgeInstanceMessageRoutingRequest() (request *UpdateEdgeInstanceMessageRoutingRequest) {
	request = &UpdateEdgeInstanceMessageRoutingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateEdgeInstanceMessageRouting", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateEdgeInstanceMessageRoutingResponse creates a response to parse from UpdateEdgeInstanceMessageRouting response
func CreateUpdateEdgeInstanceMessageRoutingResponse() (response *UpdateEdgeInstanceMessageRoutingResponse) {
	response = &UpdateEdgeInstanceMessageRoutingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
