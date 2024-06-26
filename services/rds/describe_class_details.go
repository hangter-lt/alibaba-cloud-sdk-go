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

// DescribeClassDetails invokes the rds.DescribeClassDetails API synchronously
func (client *Client) DescribeClassDetails(request *DescribeClassDetailsRequest) (response *DescribeClassDetailsResponse, err error) {
	response = CreateDescribeClassDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClassDetailsWithChan invokes the rds.DescribeClassDetails API asynchronously
func (client *Client) DescribeClassDetailsWithChan(request *DescribeClassDetailsRequest) (<-chan *DescribeClassDetailsResponse, <-chan error) {
	responseChan := make(chan *DescribeClassDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClassDetails(request)
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

// DescribeClassDetailsWithCallback invokes the rds.DescribeClassDetails API asynchronously
func (client *Client) DescribeClassDetailsWithCallback(request *DescribeClassDetailsRequest, callback func(response *DescribeClassDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClassDetailsResponse
		var err error
		defer close(result)
		response, err = client.DescribeClassDetails(request)
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

// DescribeClassDetailsRequest is the request struct for api DescribeClassDetails
type DescribeClassDetailsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	Engine               string           `position:"Query" name:"Engine"`
	ClassCode            string           `position:"Query" name:"ClassCode"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	CommodityCode        string           `position:"Query" name:"CommodityCode"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeClassDetailsResponse is the response struct for api DescribeClassDetails
type DescribeClassDetailsResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	ClassCode             string `json:"ClassCode" xml:"ClassCode"`
	MaxIOMBPS             string `json:"MaxIOMBPS" xml:"MaxIOMBPS"`
	MaxConnections        string `json:"MaxConnections" xml:"MaxConnections"`
	ClassGroup            string `json:"ClassGroup" xml:"ClassGroup"`
	Cpu                   string `json:"Cpu" xml:"Cpu"`
	InstructionSetArch    string `json:"InstructionSetArch" xml:"InstructionSetArch"`
	MemoryClass           string `json:"MemoryClass" xml:"MemoryClass"`
	MaxIOPS               string `json:"MaxIOPS" xml:"MaxIOPS"`
	ReferencePrice        string `json:"ReferencePrice" xml:"ReferencePrice"`
	Category              string `json:"Category" xml:"Category"`
	DBInstanceStorageType string `json:"DBInstanceStorageType" xml:"DBInstanceStorageType"`
}

// CreateDescribeClassDetailsRequest creates a request to invoke DescribeClassDetails API
func CreateDescribeClassDetailsRequest() (request *DescribeClassDetailsRequest) {
	request = &DescribeClassDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeClassDetails", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClassDetailsResponse creates a response to parse from DescribeClassDetails response
func CreateDescribeClassDetailsResponse() (response *DescribeClassDetailsResponse) {
	response = &DescribeClassDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
