package linkwan

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

// ListNodeGroups invokes the linkwan.ListNodeGroups API synchronously
func (client *Client) ListNodeGroups(request *ListNodeGroupsRequest) (response *ListNodeGroupsResponse, err error) {
	response = CreateListNodeGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodeGroupsWithChan invokes the linkwan.ListNodeGroups API asynchronously
func (client *Client) ListNodeGroupsWithChan(request *ListNodeGroupsRequest) (<-chan *ListNodeGroupsResponse, <-chan error) {
	responseChan := make(chan *ListNodeGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodeGroups(request)
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

// ListNodeGroupsWithCallback invokes the linkwan.ListNodeGroups API asynchronously
func (client *Client) ListNodeGroupsWithCallback(request *ListNodeGroupsRequest, callback func(response *ListNodeGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodeGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListNodeGroups(request)
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

// ListNodeGroupsRequest is the request struct for api ListNodeGroups
type ListNodeGroupsRequest struct {
	*requests.RpcRequest
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	FuzzyJoinEui  string           `position:"Query" name:"FuzzyJoinEui"`
	FuzzyDevEui   string           `position:"Query" name:"FuzzyDevEui"`
	Limit         requests.Integer `position:"Query" name:"Limit"`
	FuzzyName     string           `position:"Query" name:"FuzzyName"`
	Offset        requests.Integer `position:"Query" name:"Offset"`
	Ascending     requests.Boolean `position:"Query" name:"Ascending"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	SortingField  string           `position:"Query" name:"SortingField"`
}

// ListNodeGroupsResponse is the response struct for api ListNodeGroups
type ListNodeGroupsResponse struct {
	*responses.BaseResponse
	RequestId string               `json:"RequestId" xml:"RequestId"`
	Success   bool                 `json:"Success" xml:"Success"`
	Data      DataInListNodeGroups `json:"Data" xml:"Data"`
}

// CreateListNodeGroupsRequest creates a request to invoke ListNodeGroups API
func CreateListNodeGroupsRequest() (request *ListNodeGroupsRequest) {
	request = &ListNodeGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ListNodeGroups", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListNodeGroupsResponse creates a response to parse from ListNodeGroups response
func CreateListNodeGroupsResponse() (response *ListNodeGroupsResponse) {
	response = &ListNodeGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}