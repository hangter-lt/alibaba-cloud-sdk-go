package dataworks_public

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

// CreateQualityRelativeNode invokes the dataworks_public.CreateQualityRelativeNode API synchronously
func (client *Client) CreateQualityRelativeNode(request *CreateQualityRelativeNodeRequest) (response *CreateQualityRelativeNodeResponse, err error) {
	response = CreateCreateQualityRelativeNodeResponse()
	err = client.DoAction(request, response)
	return
}

// CreateQualityRelativeNodeWithChan invokes the dataworks_public.CreateQualityRelativeNode API asynchronously
func (client *Client) CreateQualityRelativeNodeWithChan(request *CreateQualityRelativeNodeRequest) (<-chan *CreateQualityRelativeNodeResponse, <-chan error) {
	responseChan := make(chan *CreateQualityRelativeNodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateQualityRelativeNode(request)
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

// CreateQualityRelativeNodeWithCallback invokes the dataworks_public.CreateQualityRelativeNode API asynchronously
func (client *Client) CreateQualityRelativeNodeWithCallback(request *CreateQualityRelativeNodeRequest, callback func(response *CreateQualityRelativeNodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateQualityRelativeNodeResponse
		var err error
		defer close(result)
		response, err = client.CreateQualityRelativeNode(request)
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

// CreateQualityRelativeNodeRequest is the request struct for api CreateQualityRelativeNode
type CreateQualityRelativeNodeRequest struct {
	*requests.RpcRequest
	ProjectName           string           `position:"Body" name:"ProjectName"`
	TargetNodeProjectId   requests.Integer `position:"Body" name:"TargetNodeProjectId"`
	MatchExpression       string           `position:"Body" name:"MatchExpression"`
	EnvType               string           `position:"Body" name:"EnvType"`
	TargetNodeProjectName string           `position:"Body" name:"TargetNodeProjectName"`
	TableName             string           `position:"Body" name:"TableName"`
	NodeId                requests.Integer `position:"Body" name:"NodeId"`
	ProjectId             requests.Integer `position:"Body" name:"ProjectId"`
}

// CreateQualityRelativeNodeResponse is the response struct for api CreateQualityRelativeNode
type CreateQualityRelativeNodeResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           bool   `json:"Data" xml:"Data"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateQualityRelativeNodeRequest creates a request to invoke CreateQualityRelativeNode API
func CreateCreateQualityRelativeNodeRequest() (request *CreateQualityRelativeNodeRequest) {
	request = &CreateQualityRelativeNodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateQualityRelativeNode", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateQualityRelativeNodeResponse creates a response to parse from CreateQualityRelativeNode response
func CreateCreateQualityRelativeNodeResponse() (response *CreateQualityRelativeNodeResponse) {
	response = &CreateQualityRelativeNodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}