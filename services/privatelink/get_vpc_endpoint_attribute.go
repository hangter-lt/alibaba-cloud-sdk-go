package privatelink

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

// GetVpcEndpointAttribute invokes the privatelink.GetVpcEndpointAttribute API synchronously
func (client *Client) GetVpcEndpointAttribute(request *GetVpcEndpointAttributeRequest) (response *GetVpcEndpointAttributeResponse, err error) {
	response = CreateGetVpcEndpointAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// GetVpcEndpointAttributeWithChan invokes the privatelink.GetVpcEndpointAttribute API asynchronously
func (client *Client) GetVpcEndpointAttributeWithChan(request *GetVpcEndpointAttributeRequest) (<-chan *GetVpcEndpointAttributeResponse, <-chan error) {
	responseChan := make(chan *GetVpcEndpointAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVpcEndpointAttribute(request)
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

// GetVpcEndpointAttributeWithCallback invokes the privatelink.GetVpcEndpointAttribute API asynchronously
func (client *Client) GetVpcEndpointAttributeWithCallback(request *GetVpcEndpointAttributeRequest, callback func(response *GetVpcEndpointAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVpcEndpointAttributeResponse
		var err error
		defer close(result)
		response, err = client.GetVpcEndpointAttribute(request)
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

// GetVpcEndpointAttributeRequest is the request struct for api GetVpcEndpointAttribute
type GetVpcEndpointAttributeRequest struct {
	*requests.RpcRequest
	EndpointId string `position:"Query" name:"EndpointId"`
}

// GetVpcEndpointAttributeResponse is the response struct for api GetVpcEndpointAttribute
type GetVpcEndpointAttributeResponse struct {
	*responses.BaseResponse
	Bandwidth              string `json:"Bandwidth" xml:"Bandwidth"`
	ConnectionStatus       string `json:"ConnectionStatus" xml:"ConnectionStatus"`
	CreateTime             string `json:"CreateTime" xml:"CreateTime"`
	EndpointBusinessStatus string `json:"EndpointBusinessStatus" xml:"EndpointBusinessStatus"`
	EndpointDescription    string `json:"EndpointDescription" xml:"EndpointDescription"`
	EndpointDomain         string `json:"EndpointDomain" xml:"EndpointDomain"`
	EndpointId             string `json:"EndpointId" xml:"EndpointId"`
	EndpointName           string `json:"EndpointName" xml:"EndpointName"`
	EndpointStatus         string `json:"EndpointStatus" xml:"EndpointStatus"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
	ServiceId              string `json:"ServiceId" xml:"ServiceId"`
	ServiceName            string `json:"ServiceName" xml:"ServiceName"`
	VpcId                  string `json:"VpcId" xml:"VpcId"`
	RegionId               string `json:"RegionId" xml:"RegionId"`
}

// CreateGetVpcEndpointAttributeRequest creates a request to invoke GetVpcEndpointAttribute API
func CreateGetVpcEndpointAttributeRequest() (request *GetVpcEndpointAttributeRequest) {
	request = &GetVpcEndpointAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Privatelink", "2020-04-15", "GetVpcEndpointAttribute", "privatelink", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVpcEndpointAttributeResponse creates a response to parse from GetVpcEndpointAttribute response
func CreateGetVpcEndpointAttributeResponse() (response *GetVpcEndpointAttributeResponse) {
	response = &GetVpcEndpointAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
