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

// DescribeAnalyticdbByPrimaryDBInstance invokes the rds.DescribeAnalyticdbByPrimaryDBInstance API synchronously
func (client *Client) DescribeAnalyticdbByPrimaryDBInstance(request *DescribeAnalyticdbByPrimaryDBInstanceRequest) (response *DescribeAnalyticdbByPrimaryDBInstanceResponse, err error) {
	response = CreateDescribeAnalyticdbByPrimaryDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAnalyticdbByPrimaryDBInstanceWithChan invokes the rds.DescribeAnalyticdbByPrimaryDBInstance API asynchronously
func (client *Client) DescribeAnalyticdbByPrimaryDBInstanceWithChan(request *DescribeAnalyticdbByPrimaryDBInstanceRequest) (<-chan *DescribeAnalyticdbByPrimaryDBInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeAnalyticdbByPrimaryDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAnalyticdbByPrimaryDBInstance(request)
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

// DescribeAnalyticdbByPrimaryDBInstanceWithCallback invokes the rds.DescribeAnalyticdbByPrimaryDBInstance API asynchronously
func (client *Client) DescribeAnalyticdbByPrimaryDBInstanceWithCallback(request *DescribeAnalyticdbByPrimaryDBInstanceRequest, callback func(response *DescribeAnalyticdbByPrimaryDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAnalyticdbByPrimaryDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeAnalyticdbByPrimaryDBInstance(request)
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

// DescribeAnalyticdbByPrimaryDBInstanceRequest is the request struct for api DescribeAnalyticdbByPrimaryDBInstance
type DescribeAnalyticdbByPrimaryDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeAnalyticdbByPrimaryDBInstanceResponse is the response struct for api DescribeAnalyticdbByPrimaryDBInstance
type DescribeAnalyticdbByPrimaryDBInstanceResponse struct {
	*responses.BaseResponse
	AnalyticDBCount int    `json:"AnalyticDBCount" xml:"AnalyticDBCount"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeAnalyticdbByPrimaryDBInstanceRequest creates a request to invoke DescribeAnalyticdbByPrimaryDBInstance API
func CreateDescribeAnalyticdbByPrimaryDBInstanceRequest() (request *DescribeAnalyticdbByPrimaryDBInstanceRequest) {
	request = &DescribeAnalyticdbByPrimaryDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeAnalyticdbByPrimaryDBInstance", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAnalyticdbByPrimaryDBInstanceResponse creates a response to parse from DescribeAnalyticdbByPrimaryDBInstance response
func CreateDescribeAnalyticdbByPrimaryDBInstanceResponse() (response *DescribeAnalyticdbByPrimaryDBInstanceResponse) {
	response = &DescribeAnalyticdbByPrimaryDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
