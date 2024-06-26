package live

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

// DeleteLiveStreamBlock invokes the live.DeleteLiveStreamBlock API synchronously
func (client *Client) DeleteLiveStreamBlock(request *DeleteLiveStreamBlockRequest) (response *DeleteLiveStreamBlockResponse, err error) {
	response = CreateDeleteLiveStreamBlockResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLiveStreamBlockWithChan invokes the live.DeleteLiveStreamBlock API asynchronously
func (client *Client) DeleteLiveStreamBlockWithChan(request *DeleteLiveStreamBlockRequest) (<-chan *DeleteLiveStreamBlockResponse, <-chan error) {
	responseChan := make(chan *DeleteLiveStreamBlockResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLiveStreamBlock(request)
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

// DeleteLiveStreamBlockWithCallback invokes the live.DeleteLiveStreamBlock API asynchronously
func (client *Client) DeleteLiveStreamBlockWithCallback(request *DeleteLiveStreamBlockRequest, callback func(response *DeleteLiveStreamBlockResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLiveStreamBlockResponse
		var err error
		defer close(result)
		response, err = client.DeleteLiveStreamBlock(request)
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

// DeleteLiveStreamBlockRequest is the request struct for api DeleteLiveStreamBlock
type DeleteLiveStreamBlockRequest struct {
	*requests.RpcRequest
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteLiveStreamBlockResponse is the response struct for api DeleteLiveStreamBlock
type DeleteLiveStreamBlockResponse struct {
	*responses.BaseResponse
	Description string `json:"Description" xml:"Description"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Status      string `json:"Status" xml:"Status"`
}

// CreateDeleteLiveStreamBlockRequest creates a request to invoke DeleteLiveStreamBlock API
func CreateDeleteLiveStreamBlockRequest() (request *DeleteLiveStreamBlockRequest) {
	request = &DeleteLiveStreamBlockRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DeleteLiveStreamBlock", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteLiveStreamBlockResponse creates a response to parse from DeleteLiveStreamBlock response
func CreateDeleteLiveStreamBlockResponse() (response *DeleteLiveStreamBlockResponse) {
	response = &DeleteLiveStreamBlockResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
