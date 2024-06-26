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

// DescribeQuickSaleConfig invokes the rds.DescribeQuickSaleConfig API synchronously
func (client *Client) DescribeQuickSaleConfig(request *DescribeQuickSaleConfigRequest) (response *DescribeQuickSaleConfigResponse, err error) {
	response = CreateDescribeQuickSaleConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeQuickSaleConfigWithChan invokes the rds.DescribeQuickSaleConfig API asynchronously
func (client *Client) DescribeQuickSaleConfigWithChan(request *DescribeQuickSaleConfigRequest) (<-chan *DescribeQuickSaleConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeQuickSaleConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeQuickSaleConfig(request)
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

// DescribeQuickSaleConfigWithCallback invokes the rds.DescribeQuickSaleConfig API asynchronously
func (client *Client) DescribeQuickSaleConfigWithCallback(request *DescribeQuickSaleConfigRequest, callback func(response *DescribeQuickSaleConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeQuickSaleConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeQuickSaleConfig(request)
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

// DescribeQuickSaleConfigRequest is the request struct for api DescribeQuickSaleConfig
type DescribeQuickSaleConfigRequest struct {
	*requests.RpcRequest
	Commodity string `position:"Query" name:"Commodity"`
	Engine    string `position:"Query" name:"Engine"`
}

// DescribeQuickSaleConfigResponse is the response struct for api DescribeQuickSaleConfig
type DescribeQuickSaleConfigResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Commodity string                 `json:"Commodity" xml:"Commodity"`
	Items     map[string]interface{} `json:"Items" xml:"Items"`
}

// CreateDescribeQuickSaleConfigRequest creates a request to invoke DescribeQuickSaleConfig API
func CreateDescribeQuickSaleConfigRequest() (request *DescribeQuickSaleConfigRequest) {
	request = &DescribeQuickSaleConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeQuickSaleConfig", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeQuickSaleConfigResponse creates a response to parse from DescribeQuickSaleConfig response
func CreateDescribeQuickSaleConfigResponse() (response *DescribeQuickSaleConfigResponse) {
	response = &DescribeQuickSaleConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
