package computenestsupplier

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

// ListAcrImageTags invokes the computenestsupplier.ListAcrImageTags API synchronously
func (client *Client) ListAcrImageTags(request *ListAcrImageTagsRequest) (response *ListAcrImageTagsResponse, err error) {
	response = CreateListAcrImageTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAcrImageTagsWithChan invokes the computenestsupplier.ListAcrImageTags API asynchronously
func (client *Client) ListAcrImageTagsWithChan(request *ListAcrImageTagsRequest) (<-chan *ListAcrImageTagsResponse, <-chan error) {
	responseChan := make(chan *ListAcrImageTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAcrImageTags(request)
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

// ListAcrImageTagsWithCallback invokes the computenestsupplier.ListAcrImageTags API asynchronously
func (client *Client) ListAcrImageTagsWithCallback(request *ListAcrImageTagsRequest, callback func(response *ListAcrImageTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAcrImageTagsResponse
		var err error
		defer close(result)
		response, err = client.ListAcrImageTags(request)
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

// ListAcrImageTagsRequest is the request struct for api ListAcrImageTags
type ListAcrImageTagsRequest struct {
	*requests.RpcRequest
	RepoId       string           `position:"Query" name:"RepoId"`
	NextToken    string           `position:"Query" name:"NextToken"`
	ArtifactType string           `position:"Query" name:"ArtifactType"`
	MaxResults   requests.Integer `position:"Query" name:"MaxResults"`
}

// ListAcrImageTagsResponse is the response struct for api ListAcrImageTags
type ListAcrImageTagsResponse struct {
	*responses.BaseResponse
	RequestId  string       `json:"RequestId" xml:"RequestId"`
	NextToken  string       `json:"NextToken" xml:"NextToken"`
	MaxResults int          `json:"MaxResults" xml:"MaxResults"`
	TotalCount int          `json:"TotalCount" xml:"TotalCount"`
	Images     []ImagesItem `json:"Images" xml:"Images"`
}

// CreateListAcrImageTagsRequest creates a request to invoke ListAcrImageTags API
func CreateListAcrImageTagsRequest() (request *ListAcrImageTagsRequest) {
	request = &ListAcrImageTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "ListAcrImageTags", "", "")
	request.Method = requests.POST
	return
}

// CreateListAcrImageTagsResponse creates a response to parse from ListAcrImageTags response
func CreateListAcrImageTagsResponse() (response *ListAcrImageTagsResponse) {
	response = &ListAcrImageTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
