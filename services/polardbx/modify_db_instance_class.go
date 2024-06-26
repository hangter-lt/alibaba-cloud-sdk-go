package polardbx

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

// ModifyDBInstanceClass invokes the polardbx.ModifyDBInstanceClass API synchronously
func (client *Client) ModifyDBInstanceClass(request *ModifyDBInstanceClassRequest) (response *ModifyDBInstanceClassResponse, err error) {
	response = CreateModifyDBInstanceClassResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceClassWithChan invokes the polardbx.ModifyDBInstanceClass API asynchronously
func (client *Client) ModifyDBInstanceClassWithChan(request *ModifyDBInstanceClassRequest) (<-chan *ModifyDBInstanceClassResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceClassResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceClass(request)
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

// ModifyDBInstanceClassWithCallback invokes the polardbx.ModifyDBInstanceClass API asynchronously
func (client *Client) ModifyDBInstanceClassWithCallback(request *ModifyDBInstanceClassRequest, callback func(response *ModifyDBInstanceClassResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceClassResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceClass(request)
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

// ModifyDBInstanceClassRequest is the request struct for api ModifyDBInstanceClass
type ModifyDBInstanceClassRequest struct {
	*requests.RpcRequest
	SpecifiedDNSpecMapJson string           `position:"Query" name:"SpecifiedDNSpecMapJson"`
	CnClass                string           `position:"Query" name:"CnClass"`
	TargetDBInstanceClass  string           `position:"Query" name:"TargetDBInstanceClass"`
	SpecifiedDNScale       requests.Boolean `position:"Query" name:"SpecifiedDNScale"`
	DBInstanceName         string           `position:"Query" name:"DBInstanceName"`
	ClientToken            string           `position:"Query" name:"ClientToken"`
	SwitchTimeMode         string           `position:"Query" name:"SwitchTimeMode"`
	SwitchTime             string           `position:"Query" name:"SwitchTime"`
	DnClass                string           `position:"Query" name:"DnClass"`
}

// ModifyDBInstanceClassResponse is the response struct for api ModifyDBInstanceClass
type ModifyDBInstanceClassResponse struct {
	*responses.BaseResponse
	OrderId   string `json:"OrderId" xml:"OrderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceClassRequest creates a request to invoke ModifyDBInstanceClass API
func CreateModifyDBInstanceClassRequest() (request *ModifyDBInstanceClassRequest) {
	request = &ModifyDBInstanceClassRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "ModifyDBInstanceClass", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceClassResponse creates a response to parse from ModifyDBInstanceClass response
func CreateModifyDBInstanceClassResponse() (response *ModifyDBInstanceClassResponse) {
	response = &ModifyDBInstanceClassResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
