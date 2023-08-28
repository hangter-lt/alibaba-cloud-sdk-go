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

// ExportModelFeatureTrainingSetTable invokes the paifeaturestore.ExportModelFeatureTrainingSetTable API synchronously
func (client *Client) ExportModelFeatureTrainingSetTable(request *ExportModelFeatureTrainingSetTableRequest) (response *ExportModelFeatureTrainingSetTableResponse, err error) {
	response = CreateExportModelFeatureTrainingSetTableResponse()
	err = client.DoAction(request, response)
	return
}

// ExportModelFeatureTrainingSetTableWithChan invokes the paifeaturestore.ExportModelFeatureTrainingSetTable API asynchronously
func (client *Client) ExportModelFeatureTrainingSetTableWithChan(request *ExportModelFeatureTrainingSetTableRequest) (<-chan *ExportModelFeatureTrainingSetTableResponse, <-chan error) {
	responseChan := make(chan *ExportModelFeatureTrainingSetTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportModelFeatureTrainingSetTable(request)
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

// ExportModelFeatureTrainingSetTableWithCallback invokes the paifeaturestore.ExportModelFeatureTrainingSetTable API asynchronously
func (client *Client) ExportModelFeatureTrainingSetTableWithCallback(request *ExportModelFeatureTrainingSetTableRequest, callback func(response *ExportModelFeatureTrainingSetTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportModelFeatureTrainingSetTableResponse
		var err error
		defer close(result)
		response, err = client.ExportModelFeatureTrainingSetTable(request)
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

// ExportModelFeatureTrainingSetTableRequest is the request struct for api ExportModelFeatureTrainingSetTable
type ExportModelFeatureTrainingSetTableRequest struct {
	*requests.RoaRequest
	ModelFeatureId string `position:"Path" name:"ModelFeatureId"`
	Body           string `position:"Body" name:"body"`
	InstanceId     string `position:"Path" name:"InstanceId"`
}

// ExportModelFeatureTrainingSetTableResponse is the response struct for api ExportModelFeatureTrainingSetTable
type ExportModelFeatureTrainingSetTableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateExportModelFeatureTrainingSetTableRequest creates a request to invoke ExportModelFeatureTrainingSetTable API
func CreateExportModelFeatureTrainingSetTableRequest() (request *ExportModelFeatureTrainingSetTableRequest) {
	request = &ExportModelFeatureTrainingSetTableRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "ExportModelFeatureTrainingSetTable", "/api/v1/instances/[InstanceId]/modelfeatures/[ModelFeatureId]/action/exporttrainingsettable", "", "")
	request.Method = requests.POST
	return
}

// CreateExportModelFeatureTrainingSetTableResponse creates a response to parse from ExportModelFeatureTrainingSetTable response
func CreateExportModelFeatureTrainingSetTableResponse() (response *ExportModelFeatureTrainingSetTableResponse) {
	response = &ExportModelFeatureTrainingSetTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
