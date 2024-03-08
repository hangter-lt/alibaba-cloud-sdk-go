package dds

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

// DescribeBackupStorage invokes the dds.DescribeBackupStorage API synchronously
func (client *Client) DescribeBackupStorage(request *DescribeBackupStorageRequest) (response *DescribeBackupStorageResponse, err error) {
	response = CreateDescribeBackupStorageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupStorageWithChan invokes the dds.DescribeBackupStorage API asynchronously
func (client *Client) DescribeBackupStorageWithChan(request *DescribeBackupStorageRequest) (<-chan *DescribeBackupStorageResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupStorageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupStorage(request)
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

// DescribeBackupStorageWithCallback invokes the dds.DescribeBackupStorage API asynchronously
func (client *Client) DescribeBackupStorageWithCallback(request *DescribeBackupStorageRequest, callback func(response *DescribeBackupStorageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupStorageResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupStorage(request)
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

// DescribeBackupStorageRequest is the request struct for api DescribeBackupStorage
type DescribeBackupStorageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeBackupStorageResponse is the response struct for api DescribeBackupStorage
type DescribeBackupStorageResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	FullStorageSize int64  `json:"FullStorageSize" xml:"FullStorageSize"`
	LogStorageSize  int64  `json:"LogStorageSize" xml:"LogStorageSize"`
	FreeSize        int64  `json:"FreeSize" xml:"FreeSize"`
}

// CreateDescribeBackupStorageRequest creates a request to invoke DescribeBackupStorage API
func CreateDescribeBackupStorageRequest() (request *DescribeBackupStorageRequest) {
	request = &DescribeBackupStorageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeBackupStorage", "dds", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeBackupStorageResponse creates a response to parse from DescribeBackupStorage response
func CreateDescribeBackupStorageResponse() (response *DescribeBackupStorageResponse) {
	response = &DescribeBackupStorageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
