package alikafka

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

// EnableAutoTopicCreation invokes the alikafka.EnableAutoTopicCreation API synchronously
func (client *Client) EnableAutoTopicCreation(request *EnableAutoTopicCreationRequest) (response *EnableAutoTopicCreationResponse, err error) {
	response = CreateEnableAutoTopicCreationResponse()
	err = client.DoAction(request, response)
	return
}

// EnableAutoTopicCreationWithChan invokes the alikafka.EnableAutoTopicCreation API asynchronously
func (client *Client) EnableAutoTopicCreationWithChan(request *EnableAutoTopicCreationRequest) (<-chan *EnableAutoTopicCreationResponse, <-chan error) {
	responseChan := make(chan *EnableAutoTopicCreationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EnableAutoTopicCreation(request)
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

// EnableAutoTopicCreationWithCallback invokes the alikafka.EnableAutoTopicCreation API asynchronously
func (client *Client) EnableAutoTopicCreationWithCallback(request *EnableAutoTopicCreationRequest, callback func(response *EnableAutoTopicCreationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EnableAutoTopicCreationResponse
		var err error
		defer close(result)
		response, err = client.EnableAutoTopicCreation(request)
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

// EnableAutoTopicCreationRequest is the request struct for api EnableAutoTopicCreation
type EnableAutoTopicCreationRequest struct {
	*requests.RpcRequest
	InstanceId   string           `position:"Query" name:"InstanceId"`
	Operate      string           `position:"Query" name:"Operate"`
	PartitionNum requests.Integer `position:"Query" name:"PartitionNum"`
}

// EnableAutoTopicCreationResponse is the response struct for api EnableAutoTopicCreation
type EnableAutoTopicCreationResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateEnableAutoTopicCreationRequest creates a request to invoke EnableAutoTopicCreation API
func CreateEnableAutoTopicCreationRequest() (request *EnableAutoTopicCreationRequest) {
	request = &EnableAutoTopicCreationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "EnableAutoTopicCreation", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEnableAutoTopicCreationResponse creates a response to parse from EnableAutoTopicCreation response
func CreateEnableAutoTopicCreationResponse() (response *EnableAutoTopicCreationResponse) {
	response = &EnableAutoTopicCreationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
