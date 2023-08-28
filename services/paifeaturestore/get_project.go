package paifeaturestore

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

// GetProject invokes the paifeaturestore.GetProject API synchronously
func (client *Client) GetProject(request *GetProjectRequest) (response *GetProjectResponse, err error) {
	response = CreateGetProjectResponse()
	err = client.DoAction(request, response)
	return
}

// GetProjectWithChan invokes the paifeaturestore.GetProject API asynchronously
func (client *Client) GetProjectWithChan(request *GetProjectRequest) (<-chan *GetProjectResponse, <-chan error) {
	responseChan := make(chan *GetProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetProject(request)
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

// GetProjectWithCallback invokes the paifeaturestore.GetProject API asynchronously
func (client *Client) GetProjectWithCallback(request *GetProjectRequest, callback func(response *GetProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetProjectResponse
		var err error
		defer close(result)
		response, err = client.GetProject(request)
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

// GetProjectRequest is the request struct for api GetProject
type GetProjectRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	ProjectId  string `position:"Path" name:"ProjectId"`
}

// GetProjectResponse is the response struct for api GetProject
type GetProjectResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	Name                  string `json:"Name" xml:"Name"`
	Description           string `json:"Description" xml:"Description"`
	OfflineDatasourceType string `json:"OfflineDatasourceType" xml:"OfflineDatasourceType"`
	OfflineDatasourceId   string `json:"OfflineDatasourceId" xml:"OfflineDatasourceId"`
	OfflineDatasourceName string `json:"OfflineDatasourceName" xml:"OfflineDatasourceName"`
	OnlineDatasourceType  string `json:"OnlineDatasourceType" xml:"OnlineDatasourceType"`
	OnlineDatasourceId    string `json:"OnlineDatasourceId" xml:"OnlineDatasourceId"`
	OnlineDatasourceName  string `json:"OnlineDatasourceName" xml:"OnlineDatasourceName"`
	OfflineLifecycle      int    `json:"OfflineLifecycle" xml:"OfflineLifecycle"`
	FeatureEntityCount    int    `json:"FeatureEntityCount" xml:"FeatureEntityCount"`
	FeatureViewCount      int    `json:"FeatureViewCount" xml:"FeatureViewCount"`
	ModelCount            int    `json:"ModelCount" xml:"ModelCount"`
	Owner                 string `json:"Owner" xml:"Owner"`
	GmtCreateTime         string `json:"GmtCreateTime" xml:"GmtCreateTime"`
	GmtModifiedTime       string `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
}

// CreateGetProjectRequest creates a request to invoke GetProject API
func CreateGetProjectRequest() (request *GetProjectRequest) {
	request = &GetProjectRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "GetProject", "/api/v1/instances/[InstanceId]/projects/[ProjectId]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetProjectResponse creates a response to parse from GetProject response
func CreateGetProjectResponse() (response *GetProjectResponse) {
	response = &GetProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
