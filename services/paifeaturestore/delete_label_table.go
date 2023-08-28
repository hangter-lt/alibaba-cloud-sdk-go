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

// DeleteLabelTable invokes the paifeaturestore.DeleteLabelTable API synchronously
func (client *Client) DeleteLabelTable(request *DeleteLabelTableRequest) (response *DeleteLabelTableResponse, err error) {
	response = CreateDeleteLabelTableResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLabelTableWithChan invokes the paifeaturestore.DeleteLabelTable API asynchronously
func (client *Client) DeleteLabelTableWithChan(request *DeleteLabelTableRequest) (<-chan *DeleteLabelTableResponse, <-chan error) {
	responseChan := make(chan *DeleteLabelTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLabelTable(request)
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

// DeleteLabelTableWithCallback invokes the paifeaturestore.DeleteLabelTable API asynchronously
func (client *Client) DeleteLabelTableWithCallback(request *DeleteLabelTableRequest, callback func(response *DeleteLabelTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLabelTableResponse
		var err error
		defer close(result)
		response, err = client.DeleteLabelTable(request)
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

// DeleteLabelTableRequest is the request struct for api DeleteLabelTable
type DeleteLabelTableRequest struct {
	*requests.RoaRequest
	InstanceId   string `position:"Path" name:"InstanceId"`
	LabelTableId string `position:"Path" name:"LabelTableId"`
}

// DeleteLabelTableResponse is the response struct for api DeleteLabelTable
type DeleteLabelTableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLabelTableRequest creates a request to invoke DeleteLabelTable API
func CreateDeleteLabelTableRequest() (request *DeleteLabelTableRequest) {
	request = &DeleteLabelTableRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "DeleteLabelTable", "/api/v1/instances/[InstanceId]/labeltables/[LabelTableId]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteLabelTableResponse creates a response to parse from DeleteLabelTable response
func CreateDeleteLabelTableResponse() (response *DeleteLabelTableResponse) {
	response = &DeleteLabelTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
