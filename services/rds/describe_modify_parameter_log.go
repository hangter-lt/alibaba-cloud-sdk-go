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

// DescribeModifyParameterLog invokes the rds.DescribeModifyParameterLog API synchronously
func (client *Client) DescribeModifyParameterLog(request *DescribeModifyParameterLogRequest) (response *DescribeModifyParameterLogResponse, err error) {
	response = CreateDescribeModifyParameterLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeModifyParameterLogWithChan invokes the rds.DescribeModifyParameterLog API asynchronously
func (client *Client) DescribeModifyParameterLogWithChan(request *DescribeModifyParameterLogRequest) (<-chan *DescribeModifyParameterLogResponse, <-chan error) {
	responseChan := make(chan *DescribeModifyParameterLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeModifyParameterLog(request)
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

// DescribeModifyParameterLogWithCallback invokes the rds.DescribeModifyParameterLog API asynchronously
func (client *Client) DescribeModifyParameterLogWithCallback(request *DescribeModifyParameterLogRequest, callback func(response *DescribeModifyParameterLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeModifyParameterLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeModifyParameterLog(request)
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

// DescribeModifyParameterLogRequest is the request struct for api DescribeModifyParameterLog
type DescribeModifyParameterLogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeModifyParameterLogResponse is the response struct for api DescribeModifyParameterLog
type DescribeModifyParameterLogResponse struct {
	*responses.BaseResponse
	RequestId        string                            `json:"RequestId" xml:"RequestId"`
	PageRecordCount  int                               `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int                               `json:"TotalRecordCount" xml:"TotalRecordCount"`
	DBInstanceId     string                            `json:"DBInstanceId" xml:"DBInstanceId"`
	Engine           string                            `json:"Engine" xml:"Engine"`
	PageNumber       int                               `json:"PageNumber" xml:"PageNumber"`
	EngineVersion    string                            `json:"EngineVersion" xml:"EngineVersion"`
	Items            ItemsInDescribeModifyParameterLog `json:"Items" xml:"Items"`
}

// CreateDescribeModifyParameterLogRequest creates a request to invoke DescribeModifyParameterLog API
func CreateDescribeModifyParameterLogRequest() (request *DescribeModifyParameterLogRequest) {
	request = &DescribeModifyParameterLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeModifyParameterLog", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeModifyParameterLogResponse creates a response to parse from DescribeModifyParameterLog response
func CreateDescribeModifyParameterLogResponse() (response *DescribeModifyParameterLogResponse) {
	response = &DescribeModifyParameterLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
