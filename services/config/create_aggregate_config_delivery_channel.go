package config

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

// CreateAggregateConfigDeliveryChannel invokes the config.CreateAggregateConfigDeliveryChannel API synchronously
func (client *Client) CreateAggregateConfigDeliveryChannel(request *CreateAggregateConfigDeliveryChannelRequest) (response *CreateAggregateConfigDeliveryChannelResponse, err error) {
	response = CreateCreateAggregateConfigDeliveryChannelResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAggregateConfigDeliveryChannelWithChan invokes the config.CreateAggregateConfigDeliveryChannel API asynchronously
func (client *Client) CreateAggregateConfigDeliveryChannelWithChan(request *CreateAggregateConfigDeliveryChannelRequest) (<-chan *CreateAggregateConfigDeliveryChannelResponse, <-chan error) {
	responseChan := make(chan *CreateAggregateConfigDeliveryChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAggregateConfigDeliveryChannel(request)
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

// CreateAggregateConfigDeliveryChannelWithCallback invokes the config.CreateAggregateConfigDeliveryChannel API asynchronously
func (client *Client) CreateAggregateConfigDeliveryChannelWithCallback(request *CreateAggregateConfigDeliveryChannelRequest, callback func(response *CreateAggregateConfigDeliveryChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAggregateConfigDeliveryChannelResponse
		var err error
		defer close(result)
		response, err = client.CreateAggregateConfigDeliveryChannel(request)
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

// CreateAggregateConfigDeliveryChannelRequest is the request struct for api CreateAggregateConfigDeliveryChannel
type CreateAggregateConfigDeliveryChannelRequest struct {
	*requests.RpcRequest
	NonCompliantNotification            requests.Boolean `position:"Query" name:"NonCompliantNotification"`
	ClientToken                         string           `position:"Query" name:"ClientToken"`
	ConfigurationSnapshot               requests.Boolean `position:"Query" name:"ConfigurationSnapshot"`
	Description                         string           `position:"Query" name:"Description"`
	AggregatorId                        string           `position:"Query" name:"AggregatorId"`
	DeliveryChannelTargetArn            string           `position:"Query" name:"DeliveryChannelTargetArn"`
	DeliveryChannelCondition            string           `position:"Query" name:"DeliveryChannelCondition"`
	ConfigurationItemChangeNotification requests.Boolean `position:"Query" name:"ConfigurationItemChangeNotification"`
	DeliveryChannelName                 string           `position:"Query" name:"DeliveryChannelName"`
	OversizedDataOSSTargetArn           string           `position:"Query" name:"OversizedDataOSSTargetArn"`
	DeliveryChannelType                 string           `position:"Query" name:"DeliveryChannelType"`
}

// CreateAggregateConfigDeliveryChannelResponse is the response struct for api CreateAggregateConfigDeliveryChannel
type CreateAggregateConfigDeliveryChannelResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DeliveryChannelId string `json:"DeliveryChannelId" xml:"DeliveryChannelId"`
}

// CreateCreateAggregateConfigDeliveryChannelRequest creates a request to invoke CreateAggregateConfigDeliveryChannel API
func CreateCreateAggregateConfigDeliveryChannelRequest() (request *CreateAggregateConfigDeliveryChannelRequest) {
	request = &CreateAggregateConfigDeliveryChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "CreateAggregateConfigDeliveryChannel", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateAggregateConfigDeliveryChannelResponse creates a response to parse from CreateAggregateConfigDeliveryChannel response
func CreateCreateAggregateConfigDeliveryChannelResponse() (response *CreateAggregateConfigDeliveryChannelResponse) {
	response = &CreateAggregateConfigDeliveryChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
