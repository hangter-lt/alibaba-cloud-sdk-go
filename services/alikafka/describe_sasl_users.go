package alikafka

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

// DescribeSaslUsers invokes the alikafka.DescribeSaslUsers API synchronously
func (client *Client) DescribeSaslUsers(request *DescribeSaslUsersRequest) (response *DescribeSaslUsersResponse, err error) {
	response = CreateDescribeSaslUsersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSaslUsersWithChan invokes the alikafka.DescribeSaslUsers API asynchronously
func (client *Client) DescribeSaslUsersWithChan(request *DescribeSaslUsersRequest) (<-chan *DescribeSaslUsersResponse, <-chan error) {
	responseChan := make(chan *DescribeSaslUsersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSaslUsers(request)
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

// DescribeSaslUsersWithCallback invokes the alikafka.DescribeSaslUsers API asynchronously
func (client *Client) DescribeSaslUsersWithCallback(request *DescribeSaslUsersRequest, callback func(response *DescribeSaslUsersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSaslUsersResponse
		var err error
		defer close(result)
		response, err = client.DescribeSaslUsers(request)
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

// DescribeSaslUsersRequest is the request struct for api DescribeSaslUsers
type DescribeSaslUsersRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
}

// DescribeSaslUsersResponse is the response struct for api DescribeSaslUsers
type DescribeSaslUsersResponse struct {
	*responses.BaseResponse
	Code         int          `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Success      bool         `json:"Success" xml:"Success"`
	SaslUserList SaslUserList `json:"SaslUserList" xml:"SaslUserList"`
}

// CreateDescribeSaslUsersRequest creates a request to invoke DescribeSaslUsers API
func CreateDescribeSaslUsersRequest() (request *DescribeSaslUsersRequest) {
	request = &DescribeSaslUsersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "DescribeSaslUsers", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSaslUsersResponse creates a response to parse from DescribeSaslUsers response
func CreateDescribeSaslUsersResponse() (response *DescribeSaslUsersResponse) {
	response = &DescribeSaslUsersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
