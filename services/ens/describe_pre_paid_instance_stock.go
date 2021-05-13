package ens

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

// DescribePrePaidInstanceStock invokes the ens.DescribePrePaidInstanceStock API synchronously
func (client *Client) DescribePrePaidInstanceStock(request *DescribePrePaidInstanceStockRequest) (response *DescribePrePaidInstanceStockResponse, err error) {
	response = CreateDescribePrePaidInstanceStockResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePrePaidInstanceStockWithChan invokes the ens.DescribePrePaidInstanceStock API asynchronously
func (client *Client) DescribePrePaidInstanceStockWithChan(request *DescribePrePaidInstanceStockRequest) (<-chan *DescribePrePaidInstanceStockResponse, <-chan error) {
	responseChan := make(chan *DescribePrePaidInstanceStockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePrePaidInstanceStock(request)
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

// DescribePrePaidInstanceStockWithCallback invokes the ens.DescribePrePaidInstanceStock API asynchronously
func (client *Client) DescribePrePaidInstanceStockWithCallback(request *DescribePrePaidInstanceStockRequest, callback func(response *DescribePrePaidInstanceStockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePrePaidInstanceStockResponse
		var err error
		defer close(result)
		response, err = client.DescribePrePaidInstanceStock(request)
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

// DescribePrePaidInstanceStockRequest is the request struct for api DescribePrePaidInstanceStock
type DescribePrePaidInstanceStockRequest struct {
	*requests.RpcRequest
	InstanceSpec   string           `position:"Query"`
	EnsRegionId    string           `position:"Query"`
	SystemDiskSize requests.Integer `position:"Query"`
	DataDiskSize   requests.Integer `position:"Query"`
}

// DescribePrePaidInstanceStockResponse is the response struct for api DescribePrePaidInstanceStock
type DescribePrePaidInstanceStockResponse struct {
	*responses.BaseResponse
	AvaliableCount int    `json:"AvaliableCount" xml:"AvaliableCount"`
	Cores          int    `json:"Cores" xml:"Cores"`
	DataDiskSize   int    `json:"DataDiskSize" xml:"DataDiskSize"`
	EnsRegionId    string `json:"EnsRegionId" xml:"EnsRegionId"`
	InstanceSpec   string `json:"InstanceSpec" xml:"InstanceSpec"`
	Memory         int    `json:"Memory" xml:"Memory"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	SystemDiskSize int    `json:"SystemDiskSize" xml:"SystemDiskSize"`
}

// CreateDescribePrePaidInstanceStockRequest creates a request to invoke DescribePrePaidInstanceStock API
func CreateDescribePrePaidInstanceStockRequest() (request *DescribePrePaidInstanceStockRequest) {
	request = &DescribePrePaidInstanceStockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribePrePaidInstanceStock", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribePrePaidInstanceStockResponse creates a response to parse from DescribePrePaidInstanceStock response
func CreateDescribePrePaidInstanceStockResponse() (response *DescribePrePaidInstanceStockResponse) {
	response = &DescribePrePaidInstanceStockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
