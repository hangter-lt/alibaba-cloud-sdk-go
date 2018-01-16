package mts

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

func (client *Client) DeleteMediaTag(request *DeleteMediaTagRequest) (response *DeleteMediaTagResponse, err error) {
	response = CreateDeleteMediaTagResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteMediaTagWithChan(request *DeleteMediaTagRequest) (<-chan *DeleteMediaTagResponse, <-chan error) {
	responseChan := make(chan *DeleteMediaTagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMediaTag(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DeleteMediaTagWithCallback(request *DeleteMediaTagRequest, callback func(response *DeleteMediaTagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMediaTagResponse
		var err error
		defer close(result)
		response, err = client.DeleteMediaTag(request)
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

type DeleteMediaTagRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Tag                  string           `position:"Query" name:"Tag"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
}

type DeleteMediaTagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteMediaTagRequest() (request *DeleteMediaTagRequest) {
	request = &DeleteMediaTagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "DeleteMediaTag", "mts", "openAPI")
	return
}

func CreateDeleteMediaTagResponse() (response *DeleteMediaTagResponse) {
	response = &DeleteMediaTagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
