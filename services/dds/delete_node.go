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

// DeleteNode invokes the dds.DeleteNode API synchronously
func (client *Client) DeleteNode(request *DeleteNodeRequest) (response *DeleteNodeResponse, err error) {
	response = CreateDeleteNodeResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNodeWithChan invokes the dds.DeleteNode API asynchronously
func (client *Client) DeleteNodeWithChan(request *DeleteNodeRequest) (<-chan *DeleteNodeResponse, <-chan error) {
	responseChan := make(chan *DeleteNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNode(request)
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

// DeleteNodeWithCallback invokes the dds.DeleteNode API asynchronously
func (client *Client) DeleteNodeWithCallback(request *DeleteNodeRequest, callback func(response *DeleteNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNodeResponse
		var err error
		defer close(result)
		response, err = client.DeleteNode(request)
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

// DeleteNodeRequest is the request struct for api DeleteNode
type DeleteNodeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteNodeResponse is the response struct for api DeleteNode
type DeleteNodeResponse struct {
	*responses.BaseResponse
	TaskId    int    `json:"TaskId" xml:"TaskId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateDeleteNodeRequest creates a request to invoke DeleteNode API
func CreateDeleteNodeRequest() (request *DeleteNodeRequest) {
	request = &DeleteNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DeleteNode", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteNodeResponse creates a response to parse from DeleteNode response
func CreateDeleteNodeResponse() (response *DeleteNodeResponse) {
	response = &DeleteNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
