package ddoscoo

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

// ConfigL7RsPolicy invokes the ddoscoo.ConfigL7RsPolicy API synchronously
func (client *Client) ConfigL7RsPolicy(request *ConfigL7RsPolicyRequest) (response *ConfigL7RsPolicyResponse, err error) {
	response = CreateConfigL7RsPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigL7RsPolicyWithChan invokes the ddoscoo.ConfigL7RsPolicy API asynchronously
func (client *Client) ConfigL7RsPolicyWithChan(request *ConfigL7RsPolicyRequest) (<-chan *ConfigL7RsPolicyResponse, <-chan error) {
	responseChan := make(chan *ConfigL7RsPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigL7RsPolicy(request)
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

// ConfigL7RsPolicyWithCallback invokes the ddoscoo.ConfigL7RsPolicy API asynchronously
func (client *Client) ConfigL7RsPolicyWithCallback(request *ConfigL7RsPolicyRequest, callback func(response *ConfigL7RsPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigL7RsPolicyResponse
		var err error
		defer close(result)
		response, err = client.ConfigL7RsPolicy(request)
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

// ConfigL7RsPolicyRequest is the request struct for api ConfigL7RsPolicy
type ConfigL7RsPolicyRequest struct {
	*requests.RpcRequest
	UpstreamRetry   requests.Integer `position:"Query" name:"UpstreamRetry"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	Domain          string           `position:"Query" name:"Domain"`
	Policy          string           `position:"Query" name:"Policy"`
}

// ConfigL7RsPolicyResponse is the response struct for api ConfigL7RsPolicy
type ConfigL7RsPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigL7RsPolicyRequest creates a request to invoke ConfigL7RsPolicy API
func CreateConfigL7RsPolicyRequest() (request *ConfigL7RsPolicyRequest) {
	request = &ConfigL7RsPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ConfigL7RsPolicy", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigL7RsPolicyResponse creates a response to parse from ConfigL7RsPolicy response
func CreateConfigL7RsPolicyResponse() (response *ConfigL7RsPolicyResponse) {
	response = &ConfigL7RsPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
