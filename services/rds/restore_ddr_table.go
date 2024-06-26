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

// RestoreDdrTable invokes the rds.RestoreDdrTable API synchronously
func (client *Client) RestoreDdrTable(request *RestoreDdrTableRequest) (response *RestoreDdrTableResponse, err error) {
	response = CreateRestoreDdrTableResponse()
	err = client.DoAction(request, response)
	return
}

// RestoreDdrTableWithChan invokes the rds.RestoreDdrTable API asynchronously
func (client *Client) RestoreDdrTableWithChan(request *RestoreDdrTableRequest) (<-chan *RestoreDdrTableResponse, <-chan error) {
	responseChan := make(chan *RestoreDdrTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestoreDdrTable(request)
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

// RestoreDdrTableWithCallback invokes the rds.RestoreDdrTable API asynchronously
func (client *Client) RestoreDdrTableWithCallback(request *RestoreDdrTableRequest, callback func(response *RestoreDdrTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestoreDdrTableResponse
		var err error
		defer close(result)
		response, err = client.RestoreDdrTable(request)
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

// RestoreDdrTableRequest is the request struct for api RestoreDdrTable
type RestoreDdrTableRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceDBInstanceName string           `position:"Query" name:"SourceDBInstanceName"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	UserBakSetURL        string           `position:"Query" name:"UserBakSetURL"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	TableMeta            string           `position:"Query" name:"TableMeta"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	BackupSetRegion      string           `position:"Query" name:"BackupSetRegion"`
	BackupSetType        string           `position:"Query" name:"BackupSetType"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	BakSetName           string           `position:"Query" name:"BakSetName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	RestoreType          string           `position:"Query" name:"RestoreType"`
	SourceRegion         string           `position:"Query" name:"SourceRegion"`
}

// RestoreDdrTableResponse is the response struct for api RestoreDdrTable
type RestoreDdrTableResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
}

// CreateRestoreDdrTableRequest creates a request to invoke RestoreDdrTable API
func CreateRestoreDdrTableRequest() (request *RestoreDdrTableRequest) {
	request = &RestoreDdrTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "RestoreDdrTable", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRestoreDdrTableResponse creates a response to parse from RestoreDdrTable response
func CreateRestoreDdrTableResponse() (response *RestoreDdrTableResponse) {
	response = &RestoreDdrTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
