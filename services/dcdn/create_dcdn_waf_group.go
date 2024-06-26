package dcdn

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

// CreateDcdnWafGroup invokes the dcdn.CreateDcdnWafGroup API synchronously
func (client *Client) CreateDcdnWafGroup(request *CreateDcdnWafGroupRequest) (response *CreateDcdnWafGroupResponse, err error) {
	response = CreateCreateDcdnWafGroupResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDcdnWafGroupWithChan invokes the dcdn.CreateDcdnWafGroup API asynchronously
func (client *Client) CreateDcdnWafGroupWithChan(request *CreateDcdnWafGroupRequest) (<-chan *CreateDcdnWafGroupResponse, <-chan error) {
	responseChan := make(chan *CreateDcdnWafGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDcdnWafGroup(request)
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

// CreateDcdnWafGroupWithCallback invokes the dcdn.CreateDcdnWafGroup API asynchronously
func (client *Client) CreateDcdnWafGroupWithCallback(request *CreateDcdnWafGroupRequest, callback func(response *CreateDcdnWafGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDcdnWafGroupResponse
		var err error
		defer close(result)
		response, err = client.CreateDcdnWafGroup(request)
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

// CreateDcdnWafGroupRequest is the request struct for api CreateDcdnWafGroup
type CreateDcdnWafGroupRequest struct {
	*requests.RpcRequest
	Subscribe  string           `position:"Body" name:"Subscribe"`
	Name       string           `position:"Body" name:"Name"`
	TemplateId requests.Integer `position:"Body" name:"TemplateId"`
}

// CreateDcdnWafGroupResponse is the response struct for api CreateDcdnWafGroup
type CreateDcdnWafGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        int64  `json:"Id" xml:"Id"`
}

// CreateCreateDcdnWafGroupRequest creates a request to invoke CreateDcdnWafGroup API
func CreateCreateDcdnWafGroupRequest() (request *CreateDcdnWafGroupRequest) {
	request = &CreateDcdnWafGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "CreateDcdnWafGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDcdnWafGroupResponse creates a response to parse from CreateDcdnWafGroup response
func CreateCreateDcdnWafGroupResponse() (response *CreateDcdnWafGroupResponse) {
	response = &CreateDcdnWafGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
