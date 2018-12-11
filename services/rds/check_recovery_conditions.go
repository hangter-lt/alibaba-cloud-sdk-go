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

// CheckRecoveryConditions invokes the rds.CheckRecoveryConditions API synchronously
// api document: https://help.aliyun.com/api/rds/checkrecoveryconditions.html
func (client *Client) CheckRecoveryConditions(request *CheckRecoveryConditionsRequest) (response *CheckRecoveryConditionsResponse, err error) {
	response = CreateCheckRecoveryConditionsResponse()
	err = client.DoAction(request, response)
	return
}

// CheckRecoveryConditionsWithChan invokes the rds.CheckRecoveryConditions API asynchronously
// api document: https://help.aliyun.com/api/rds/checkrecoveryconditions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckRecoveryConditionsWithChan(request *CheckRecoveryConditionsRequest) (<-chan *CheckRecoveryConditionsResponse, <-chan error) {
	responseChan := make(chan *CheckRecoveryConditionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckRecoveryConditions(request)
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

// CheckRecoveryConditionsWithCallback invokes the rds.CheckRecoveryConditions API asynchronously
// api document: https://help.aliyun.com/api/rds/checkrecoveryconditions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckRecoveryConditionsWithCallback(request *CheckRecoveryConditionsRequest, callback func(response *CheckRecoveryConditionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckRecoveryConditionsResponse
		var err error
		defer close(result)
		response, err = client.CheckRecoveryConditions(request)
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

// CheckRecoveryConditionsRequest is the request struct for api CheckRecoveryConditions
type CheckRecoveryConditionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupFile           string           `position:"Query" name:"BackupFile"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CheckRecoveryConditionsResponse is the response struct for api CheckRecoveryConditions
type CheckRecoveryConditionsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	DBInstanceId   string `json:"DBInstanceId" xml:"DBInstanceId"`
	RecoveryStatus string `json:"RecoveryStatus" xml:"RecoveryStatus"`
}

// CreateCheckRecoveryConditionsRequest creates a request to invoke CheckRecoveryConditions API
func CreateCheckRecoveryConditionsRequest() (request *CheckRecoveryConditionsRequest) {
	request = &CheckRecoveryConditionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CheckRecoveryConditions", "Rds", "openAPI")
	return
}

// CreateCheckRecoveryConditionsResponse creates a response to parse from CheckRecoveryConditions response
func CreateCheckRecoveryConditionsResponse() (response *CheckRecoveryConditionsResponse) {
	response = &CheckRecoveryConditionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
