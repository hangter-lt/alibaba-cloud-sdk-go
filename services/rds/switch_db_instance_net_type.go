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

// SwitchDBInstanceNetType invokes the rds.SwitchDBInstanceNetType API synchronously
func (client *Client) SwitchDBInstanceNetType(request *SwitchDBInstanceNetTypeRequest) (response *SwitchDBInstanceNetTypeResponse, err error) {
	response = CreateSwitchDBInstanceNetTypeResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchDBInstanceNetTypeWithChan invokes the rds.SwitchDBInstanceNetType API asynchronously
func (client *Client) SwitchDBInstanceNetTypeWithChan(request *SwitchDBInstanceNetTypeRequest) (<-chan *SwitchDBInstanceNetTypeResponse, <-chan error) {
	responseChan := make(chan *SwitchDBInstanceNetTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchDBInstanceNetType(request)
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

// SwitchDBInstanceNetTypeWithCallback invokes the rds.SwitchDBInstanceNetType API asynchronously
func (client *Client) SwitchDBInstanceNetTypeWithCallback(request *SwitchDBInstanceNetTypeRequest, callback func(response *SwitchDBInstanceNetTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchDBInstanceNetTypeResponse
		var err error
		defer close(result)
		response, err = client.SwitchDBInstanceNetType(request)
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

// SwitchDBInstanceNetTypeRequest is the request struct for api SwitchDBInstanceNetType
type SwitchDBInstanceNetTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string           `position:"Query" name:"ConnectionStringPrefix"`
	ClientToken            string           `position:"Query" name:"ClientToken"`
	DBInstanceId           string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	ConnectionStringType   string           `position:"Query" name:"ConnectionStringType"`
	Port                   string           `position:"Query" name:"Port"`
}

// SwitchDBInstanceNetTypeResponse is the response struct for api SwitchDBInstanceNetType
type SwitchDBInstanceNetTypeResponse struct {
	*responses.BaseResponse
	NewConnectionString string `json:"NewConnectionString" xml:"NewConnectionString"`
	RequestId           string `json:"RequestId" xml:"RequestId"`
	OldConnectionString string `json:"OldConnectionString" xml:"OldConnectionString"`
}

// CreateSwitchDBInstanceNetTypeRequest creates a request to invoke SwitchDBInstanceNetType API
func CreateSwitchDBInstanceNetTypeRequest() (request *SwitchDBInstanceNetTypeRequest) {
	request = &SwitchDBInstanceNetTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "SwitchDBInstanceNetType", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSwitchDBInstanceNetTypeResponse creates a response to parse from SwitchDBInstanceNetType response
func CreateSwitchDBInstanceNetTypeResponse() (response *SwitchDBInstanceNetTypeResponse) {
	response = &SwitchDBInstanceNetTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
