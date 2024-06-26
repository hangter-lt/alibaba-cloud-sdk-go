package rds

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

// ModifyDBProxyEndpoint invokes the rds.ModifyDBProxyEndpoint API synchronously
func (client *Client) ModifyDBProxyEndpoint(request *ModifyDBProxyEndpointRequest) (response *ModifyDBProxyEndpointResponse, err error) {
	response = CreateModifyDBProxyEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBProxyEndpointWithChan invokes the rds.ModifyDBProxyEndpoint API asynchronously
func (client *Client) ModifyDBProxyEndpointWithChan(request *ModifyDBProxyEndpointRequest) (<-chan *ModifyDBProxyEndpointResponse, <-chan error) {
	responseChan := make(chan *ModifyDBProxyEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBProxyEndpoint(request)
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

// ModifyDBProxyEndpointWithCallback invokes the rds.ModifyDBProxyEndpoint API asynchronously
func (client *Client) ModifyDBProxyEndpointWithCallback(request *ModifyDBProxyEndpointRequest, callback func(response *ModifyDBProxyEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBProxyEndpointResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBProxyEndpoint(request)
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

// ModifyDBProxyEndpointRequest is the request struct for api ModifyDBProxyEndpoint
type ModifyDBProxyEndpointRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConfigDBProxyFeatures            string           `position:"Query" name:"ConfigDBProxyFeatures"`
	DBInstanceId                     string           `position:"Query" name:"DBInstanceId"`
	ReadOnlyInstanceWeight           string           `position:"Query" name:"ReadOnlyInstanceWeight"`
	ReadOnlyInstanceMaxDelayTime     string           `position:"Query" name:"ReadOnlyInstanceMaxDelayTime"`
	ResourceOwnerAccount             string           `position:"Query" name:"ResourceOwnerAccount"`
	DbEndpointAliases                string           `position:"Query" name:"DbEndpointAliases"`
	DBProxyEngineType                string           `position:"Query" name:"DBProxyEngineType"`
	DbEndpointOperator               string           `position:"Query" name:"DbEndpointOperator"`
	DbEndpointType                   string           `position:"Query" name:"DbEndpointType"`
	OwnerId                          requests.Integer `position:"Query" name:"OwnerId"`
	DbEndpointReadWriteMode          string           `position:"Query" name:"DbEndpointReadWriteMode"`
	DBProxyEndpointId                string           `position:"Query" name:"DBProxyEndpointId"`
	ReadOnlyInstanceDistributionType string           `position:"Query" name:"ReadOnlyInstanceDistributionType"`
}

// ModifyDBProxyEndpointResponse is the response struct for api ModifyDBProxyEndpoint
type ModifyDBProxyEndpointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBProxyEndpointRequest creates a request to invoke ModifyDBProxyEndpoint API
func CreateModifyDBProxyEndpointRequest() (request *ModifyDBProxyEndpointRequest) {
	request = &ModifyDBProxyEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBProxyEndpoint", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBProxyEndpointResponse creates a response to parse from ModifyDBProxyEndpoint response
func CreateModifyDBProxyEndpointResponse() (response *ModifyDBProxyEndpointResponse) {
	response = &ModifyDBProxyEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
