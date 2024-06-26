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

// ListServiceUsages invokes the computenestsupplier.ListServiceUsages API synchronously
func (client *Client) ListServiceUsages(request *ListServiceUsagesRequest) (response *ListServiceUsagesResponse, err error) {
	response = CreateListServiceUsagesResponse()
	err = client.DoAction(request, response)
	return
}

// ListServiceUsagesWithChan invokes the computenestsupplier.ListServiceUsages API asynchronously
func (client *Client) ListServiceUsagesWithChan(request *ListServiceUsagesRequest) (<-chan *ListServiceUsagesResponse, <-chan error) {
	responseChan := make(chan *ListServiceUsagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListServiceUsages(request)
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

// ListServiceUsagesWithCallback invokes the computenestsupplier.ListServiceUsages API asynchronously
func (client *Client) ListServiceUsagesWithCallback(request *ListServiceUsagesRequest, callback func(response *ListServiceUsagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListServiceUsagesResponse
		var err error
		defer close(result)
		response, err = client.ListServiceUsages(request)
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

// ListServiceUsagesRequest is the request struct for api ListServiceUsages
type ListServiceUsagesRequest struct {
	*requests.RpcRequest
	NextToken    string                     `position:"Query" name:"NextToken"`
	SupplierRole string                     `position:"Query" name:"SupplierRole"`
	Filter       *[]ListServiceUsagesFilter `position:"Query" name:"Filter"  type:"Repeated"`
	MaxResults   requests.Integer           `position:"Query" name:"MaxResults"`
}

// ListServiceUsagesFilter is a repeated param struct in ListServiceUsagesRequest
type ListServiceUsagesFilter struct {
	Name  string    `name:"Name"`
	Value *[]string `name:"Value" type:"Repeated"`
}

// ListServiceUsagesResponse is the response struct for api ListServiceUsages
type ListServiceUsagesResponse struct {
	*responses.BaseResponse
	RequestId     string   `json:"RequestId" xml:"RequestId"`
	NextToken     string   `json:"NextToken" xml:"NextToken"`
	MaxResults    int      `json:"MaxResults" xml:"MaxResults"`
	TotalCount    int      `json:"TotalCount" xml:"TotalCount"`
	ServiceUsages []Policy `json:"ServiceUsages" xml:"ServiceUsages"`
}

// CreateListServiceUsagesRequest creates a request to invoke ListServiceUsages API
func CreateListServiceUsagesRequest() (request *ListServiceUsagesRequest) {
	request = &ListServiceUsagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "ListServiceUsages", "", "")
	request.Method = requests.POST
	return
}

// CreateListServiceUsagesResponse creates a response to parse from ListServiceUsages response
func CreateListServiceUsagesResponse() (response *ListServiceUsagesResponse) {
	response = &ListServiceUsagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
