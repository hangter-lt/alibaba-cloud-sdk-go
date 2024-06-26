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

// DeleteDB invokes the polardbx.DeleteDB API synchronously
func (client *Client) DeleteDB(request *DeleteDBRequest) (response *DeleteDBResponse, err error) {
	response = CreateDeleteDBResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDBWithChan invokes the polardbx.DeleteDB API asynchronously
func (client *Client) DeleteDBWithChan(request *DeleteDBRequest) (<-chan *DeleteDBResponse, <-chan error) {
	responseChan := make(chan *DeleteDBResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDB(request)
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

// DeleteDBWithCallback invokes the polardbx.DeleteDB API asynchronously
func (client *Client) DeleteDBWithCallback(request *DeleteDBRequest, callback func(response *DeleteDBResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDBResponse
		var err error
		defer close(result)
		response, err = client.DeleteDB(request)
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

// DeleteDBRequest is the request struct for api DeleteDB
type DeleteDBRequest struct {
	*requests.RpcRequest
	DBInstanceName string `position:"Query" name:"DBInstanceName"`
	DbName         string `position:"Query" name:"DbName"`
}

// DeleteDBResponse is the response struct for api DeleteDB
type DeleteDBResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteDBRequest creates a request to invoke DeleteDB API
func CreateDeleteDBRequest() (request *DeleteDBRequest) {
	request = &DeleteDBRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "DeleteDB", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDBResponse creates a response to parse from DeleteDB response
func CreateDeleteDBResponse() (response *DeleteDBResponse) {
	response = &DeleteDBResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
