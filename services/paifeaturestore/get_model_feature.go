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

// GetModelFeature invokes the paifeaturestore.GetModelFeature API synchronously
func (client *Client) GetModelFeature(request *GetModelFeatureRequest) (response *GetModelFeatureResponse, err error) {
	response = CreateGetModelFeatureResponse()
	err = client.DoAction(request, response)
	return
}

// GetModelFeatureWithChan invokes the paifeaturestore.GetModelFeature API asynchronously
func (client *Client) GetModelFeatureWithChan(request *GetModelFeatureRequest) (<-chan *GetModelFeatureResponse, <-chan error) {
	responseChan := make(chan *GetModelFeatureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetModelFeature(request)
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

// GetModelFeatureWithCallback invokes the paifeaturestore.GetModelFeature API asynchronously
func (client *Client) GetModelFeatureWithCallback(request *GetModelFeatureRequest, callback func(response *GetModelFeatureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetModelFeatureResponse
		var err error
		defer close(result)
		response, err = client.GetModelFeature(request)
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

// GetModelFeatureRequest is the request struct for api GetModelFeature
type GetModelFeatureRequest struct {
	*requests.RoaRequest
	ModelFeatureId string `position:"Path" name:"ModelFeatureId"`
	InstanceId     string `position:"Path" name:"InstanceId"`
}

// GetModelFeatureResponse is the response struct for api GetModelFeature
type GetModelFeatureResponse struct {
	*responses.BaseResponse
	RequestId          string         `json:"RequestId" xml:"RequestId"`
	ProjectId          string         `json:"ProjectId" xml:"ProjectId"`
	ProjectName        string         `json:"ProjectName" xml:"ProjectName"`
	Name               string         `json:"Name" xml:"Name"`
	Owner              string         `json:"Owner" xml:"Owner"`
	GmtCreateTime      string         `json:"GmtCreateTime" xml:"GmtCreateTime"`
	GmtModifiedTime    string         `json:"GmtModifiedTime" xml:"GmtModifiedTime"`
	LabelTableId       string         `json:"LabelTableId" xml:"LabelTableId"`
	LabelTableName     string         `json:"LabelTableName" xml:"LabelTableName"`
	TrainingSetTable   string         `json:"TrainingSetTable" xml:"TrainingSetTable"`
	TrainingSetFGTable string         `json:"TrainingSetFGTable" xml:"TrainingSetFGTable"`
	Relations          Relations      `json:"Relations" xml:"Relations"`
	Features           []FeaturesItem `json:"Features" xml:"Features"`
}

// CreateGetModelFeatureRequest creates a request to invoke GetModelFeature API
func CreateGetModelFeatureRequest() (request *GetModelFeatureRequest) {
	request = &GetModelFeatureRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "GetModelFeature", "/api/v1/instances/[InstanceId]/modelfeatures/[ModelFeatureId]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetModelFeatureResponse creates a response to parse from GetModelFeature response
func CreateGetModelFeatureResponse() (response *GetModelFeatureResponse) {
	response = &GetModelFeatureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
