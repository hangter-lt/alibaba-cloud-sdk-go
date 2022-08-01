package polardb

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

// ModifyDBClusterMigration invokes the polardb.ModifyDBClusterMigration API synchronously
func (client *Client) ModifyDBClusterMigration(request *ModifyDBClusterMigrationRequest) (response *ModifyDBClusterMigrationResponse, err error) {
	response = CreateModifyDBClusterMigrationResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterMigrationWithChan invokes the polardb.ModifyDBClusterMigration API asynchronously
func (client *Client) ModifyDBClusterMigrationWithChan(request *ModifyDBClusterMigrationRequest) (<-chan *ModifyDBClusterMigrationResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterMigrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterMigration(request)
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

// ModifyDBClusterMigrationWithCallback invokes the polardb.ModifyDBClusterMigration API asynchronously
func (client *Client) ModifyDBClusterMigrationWithCallback(request *ModifyDBClusterMigrationRequest, callback func(response *ModifyDBClusterMigrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterMigrationResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterMigration(request)
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

// ModifyDBClusterMigrationRequest is the request struct for api ModifyDBClusterMigration
type ModifyDBClusterMigrationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConnectionStrings     string           `position:"Query" name:"ConnectionStrings"`
	SecurityToken         string           `position:"Query" name:"SecurityToken"`
	NewMasterInstanceId   string           `position:"Query" name:"NewMasterInstanceId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId           string           `position:"Query" name:"DBClusterId"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	SourceRDSDBInstanceId string           `position:"Query" name:"SourceRDSDBInstanceId"`
	SwapConnectionString  string           `position:"Query" name:"SwapConnectionString"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyDBClusterMigrationResponse is the response struct for api ModifyDBClusterMigration
type ModifyDBClusterMigrationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterMigrationRequest creates a request to invoke ModifyDBClusterMigration API
func CreateModifyDBClusterMigrationRequest() (request *ModifyDBClusterMigrationRequest) {
	request = &ModifyDBClusterMigrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterMigration", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterMigrationResponse creates a response to parse from ModifyDBClusterMigration response
func CreateModifyDBClusterMigrationResponse() (response *ModifyDBClusterMigrationResponse) {
	response = &ModifyDBClusterMigrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
