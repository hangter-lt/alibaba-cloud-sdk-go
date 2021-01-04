package alidns

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

// AddDnsGtmMonitor invokes the alidns.AddDnsGtmMonitor API synchronously
func (client *Client) AddDnsGtmMonitor(request *AddDnsGtmMonitorRequest) (response *AddDnsGtmMonitorResponse, err error) {
	response = CreateAddDnsGtmMonitorResponse()
	err = client.DoAction(request, response)
	return
}

// AddDnsGtmMonitorWithChan invokes the alidns.AddDnsGtmMonitor API asynchronously
func (client *Client) AddDnsGtmMonitorWithChan(request *AddDnsGtmMonitorRequest) (<-chan *AddDnsGtmMonitorResponse, <-chan error) {
	responseChan := make(chan *AddDnsGtmMonitorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDnsGtmMonitor(request)
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

// AddDnsGtmMonitorWithCallback invokes the alidns.AddDnsGtmMonitor API asynchronously
func (client *Client) AddDnsGtmMonitorWithCallback(request *AddDnsGtmMonitorRequest, callback func(response *AddDnsGtmMonitorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDnsGtmMonitorResponse
		var err error
		defer close(result)
		response, err = client.AddDnsGtmMonitor(request)
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

// AddDnsGtmMonitorRequest is the request struct for api AddDnsGtmMonitor
type AddDnsGtmMonitorRequest struct {
	*requests.RpcRequest
	MonitorExtendInfo string                         `position:"Query" name:"MonitorExtendInfo"`
	Timeout           requests.Integer               `position:"Query" name:"Timeout"`
	AddrPoolId        string                         `position:"Query" name:"AddrPoolId"`
	UserClientIp      string                         `position:"Query" name:"UserClientIp"`
	EvaluationCount   requests.Integer               `position:"Query" name:"EvaluationCount"`
	ProtocolType      string                         `position:"Query" name:"ProtocolType"`
	Interval          requests.Integer               `position:"Query" name:"Interval"`
	Lang              string                         `position:"Query" name:"Lang"`
	IspCityNode       *[]AddDnsGtmMonitorIspCityNode `position:"Query" name:"IspCityNode"  type:"Repeated"`
}

// AddDnsGtmMonitorIspCityNode is a repeated param struct in AddDnsGtmMonitorRequest
type AddDnsGtmMonitorIspCityNode struct {
	CityCode string `name:"CityCode"`
	IspCode  string `name:"IspCode"`
}

// AddDnsGtmMonitorResponse is the response struct for api AddDnsGtmMonitor
type AddDnsGtmMonitorResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	MonitorConfigId string `json:"MonitorConfigId" xml:"MonitorConfigId"`
}

// CreateAddDnsGtmMonitorRequest creates a request to invoke AddDnsGtmMonitor API
func CreateAddDnsGtmMonitorRequest() (request *AddDnsGtmMonitorRequest) {
	request = &AddDnsGtmMonitorRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "AddDnsGtmMonitor", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddDnsGtmMonitorResponse creates a response to parse from AddDnsGtmMonitor response
func CreateAddDnsGtmMonitorResponse() (response *AddDnsGtmMonitorResponse) {
	response = &AddDnsGtmMonitorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
