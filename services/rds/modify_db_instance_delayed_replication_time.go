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

// ModifyDBInstanceDelayedReplicationTime invokes the rds.ModifyDBInstanceDelayedReplicationTime API synchronously
func (client *Client) ModifyDBInstanceDelayedReplicationTime(request *ModifyDBInstanceDelayedReplicationTimeRequest) (response *ModifyDBInstanceDelayedReplicationTimeResponse, err error) {
	response = CreateModifyDBInstanceDelayedReplicationTimeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceDelayedReplicationTimeWithChan invokes the rds.ModifyDBInstanceDelayedReplicationTime API asynchronously
func (client *Client) ModifyDBInstanceDelayedReplicationTimeWithChan(request *ModifyDBInstanceDelayedReplicationTimeRequest) (<-chan *ModifyDBInstanceDelayedReplicationTimeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceDelayedReplicationTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceDelayedReplicationTime(request)
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

// ModifyDBInstanceDelayedReplicationTimeWithCallback invokes the rds.ModifyDBInstanceDelayedReplicationTime API asynchronously
func (client *Client) ModifyDBInstanceDelayedReplicationTimeWithCallback(request *ModifyDBInstanceDelayedReplicationTimeRequest, callback func(response *ModifyDBInstanceDelayedReplicationTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceDelayedReplicationTimeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceDelayedReplicationTime(request)
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

// ModifyDBInstanceDelayedReplicationTimeRequest is the request struct for api ModifyDBInstanceDelayedReplicationTime
type ModifyDBInstanceDelayedReplicationTimeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ReadSQLReplicationTime string           `position:"Query" name:"ReadSQLReplicationTime"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
}

// ModifyDBInstanceDelayedReplicationTimeResponse is the response struct for api ModifyDBInstanceDelayedReplicationTime
type ModifyDBInstanceDelayedReplicationTimeResponse struct {
	*responses.BaseResponse
	DBInstanceId           string `json:"DBInstanceId" xml:"DBInstanceId"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
	TaskId                 string `json:"TaskId" xml:"TaskId"`
	ReadSQLReplicationTime string `json:"ReadSQLReplicationTime" xml:"ReadSQLReplicationTime"`
}

// CreateModifyDBInstanceDelayedReplicationTimeRequest creates a request to invoke ModifyDBInstanceDelayedReplicationTime API
func CreateModifyDBInstanceDelayedReplicationTimeRequest() (request *ModifyDBInstanceDelayedReplicationTimeRequest) {
	request = &ModifyDBInstanceDelayedReplicationTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceDelayedReplicationTime", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceDelayedReplicationTimeResponse creates a response to parse from ModifyDBInstanceDelayedReplicationTime response
func CreateModifyDBInstanceDelayedReplicationTimeResponse() (response *ModifyDBInstanceDelayedReplicationTimeResponse) {
	response = &ModifyDBInstanceDelayedReplicationTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
