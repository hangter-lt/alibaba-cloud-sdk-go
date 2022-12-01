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

// ListPrometheusAlertTemplates invokes the arms.ListPrometheusAlertTemplates API synchronously
func (client *Client) ListPrometheusAlertTemplates(request *ListPrometheusAlertTemplatesRequest) (response *ListPrometheusAlertTemplatesResponse, err error) {
	response = CreateListPrometheusAlertTemplatesResponse()
	err = client.DoAction(request, response)
	return
}

// ListPrometheusAlertTemplatesWithChan invokes the arms.ListPrometheusAlertTemplates API asynchronously
func (client *Client) ListPrometheusAlertTemplatesWithChan(request *ListPrometheusAlertTemplatesRequest) (<-chan *ListPrometheusAlertTemplatesResponse, <-chan error) {
	responseChan := make(chan *ListPrometheusAlertTemplatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPrometheusAlertTemplates(request)
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

// ListPrometheusAlertTemplatesWithCallback invokes the arms.ListPrometheusAlertTemplates API asynchronously
func (client *Client) ListPrometheusAlertTemplatesWithCallback(request *ListPrometheusAlertTemplatesRequest, callback func(response *ListPrometheusAlertTemplatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPrometheusAlertTemplatesResponse
		var err error
		defer close(result)
		response, err = client.ListPrometheusAlertTemplates(request)
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

// ListPrometheusAlertTemplatesRequest is the request struct for api ListPrometheusAlertTemplates
type ListPrometheusAlertTemplatesRequest struct {
	*requests.RpcRequest
	ProductCode string `position:"Query" name:"ProductCode"`
	ClusterId   string `position:"Query" name:"ClusterId"`
	ProxyUserId string `position:"Query" name:"ProxyUserId"`
}

// ListPrometheusAlertTemplatesResponse is the response struct for api ListPrometheusAlertTemplates
type ListPrometheusAlertTemplatesResponse struct {
	*responses.BaseResponse
	RequestId                string                    `json:"RequestId" xml:"RequestId"`
	PrometheusAlertTemplates []PrometheusAlertTemplate `json:"PrometheusAlertTemplates" xml:"PrometheusAlertTemplates"`
}

// CreateListPrometheusAlertTemplatesRequest creates a request to invoke ListPrometheusAlertTemplates API
func CreateListPrometheusAlertTemplatesRequest() (request *ListPrometheusAlertTemplatesRequest) {
	request = &ListPrometheusAlertTemplatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2021-05-19", "ListPrometheusAlertTemplates", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListPrometheusAlertTemplatesResponse creates a response to parse from ListPrometheusAlertTemplates response
func CreateListPrometheusAlertTemplatesResponse() (response *ListPrometheusAlertTemplatesResponse) {
	response = &ListPrometheusAlertTemplatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
