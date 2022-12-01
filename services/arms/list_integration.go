package arms

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

// ListIntegration invokes the arms.ListIntegration API synchronously
func (client *Client) ListIntegration(request *ListIntegrationRequest) (response *ListIntegrationResponse, err error) {
	response = CreateListIntegrationResponse()
	err = client.DoAction(request, response)
	return
}

// ListIntegrationWithChan invokes the arms.ListIntegration API asynchronously
func (client *Client) ListIntegrationWithChan(request *ListIntegrationRequest) (<-chan *ListIntegrationResponse, <-chan error) {
	responseChan := make(chan *ListIntegrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListIntegration(request)
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

// ListIntegrationWithCallback invokes the arms.ListIntegration API asynchronously
func (client *Client) ListIntegrationWithCallback(request *ListIntegrationRequest, callback func(response *ListIntegrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListIntegrationResponse
		var err error
		defer close(result)
		response, err = client.ListIntegration(request)
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

// ListIntegrationRequest is the request struct for api ListIntegration
type ListIntegrationRequest struct {
	*requests.RpcRequest
	Size                   requests.Integer `position:"Query" name:"Size"`
	IntegrationName        string           `position:"Query" name:"IntegrationName"`
	IsDetail               requests.Boolean `position:"Query" name:"IsDetail"`
	Page                   requests.Integer `position:"Query" name:"Page"`
	IntegrationProductType string           `position:"Query" name:"IntegrationProductType"`
}

// ListIntegrationResponse is the response struct for api ListIntegration
type ListIntegrationResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	PageInfo  PageInfo `json:"PageInfo" xml:"PageInfo"`
}

// CreateListIntegrationRequest creates a request to invoke ListIntegration API
func CreateListIntegrationRequest() (request *ListIntegrationRequest) {
	request = &ListIntegrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListIntegration", "arms", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListIntegrationResponse creates a response to parse from ListIntegration response
func CreateListIntegrationResponse() (response *ListIntegrationResponse) {
	response = &ListIntegrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
