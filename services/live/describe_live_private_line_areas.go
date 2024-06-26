package live

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

// DescribeLivePrivateLineAreas invokes the live.DescribeLivePrivateLineAreas API synchronously
func (client *Client) DescribeLivePrivateLineAreas(request *DescribeLivePrivateLineAreasRequest) (response *DescribeLivePrivateLineAreasResponse, err error) {
	response = CreateDescribeLivePrivateLineAreasResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLivePrivateLineAreasWithChan invokes the live.DescribeLivePrivateLineAreas API asynchronously
func (client *Client) DescribeLivePrivateLineAreasWithChan(request *DescribeLivePrivateLineAreasRequest) (<-chan *DescribeLivePrivateLineAreasResponse, <-chan error) {
	responseChan := make(chan *DescribeLivePrivateLineAreasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLivePrivateLineAreas(request)
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

// DescribeLivePrivateLineAreasWithCallback invokes the live.DescribeLivePrivateLineAreas API asynchronously
func (client *Client) DescribeLivePrivateLineAreasWithCallback(request *DescribeLivePrivateLineAreasRequest, callback func(response *DescribeLivePrivateLineAreasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLivePrivateLineAreasResponse
		var err error
		defer close(result)
		response, err = client.DescribeLivePrivateLineAreas(request)
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

// DescribeLivePrivateLineAreasRequest is the request struct for api DescribeLivePrivateLineAreas
type DescribeLivePrivateLineAreasRequest struct {
	*requests.RpcRequest
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLivePrivateLineAreasResponse is the response struct for api DescribeLivePrivateLineAreas
type DescribeLivePrivateLineAreasResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	LiveAreasList LiveAreasList `json:"LiveAreasList" xml:"LiveAreasList"`
}

// CreateDescribeLivePrivateLineAreasRequest creates a request to invoke DescribeLivePrivateLineAreas API
func CreateDescribeLivePrivateLineAreasRequest() (request *DescribeLivePrivateLineAreasRequest) {
	request = &DescribeLivePrivateLineAreasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLivePrivateLineAreas", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLivePrivateLineAreasResponse creates a response to parse from DescribeLivePrivateLineAreas response
func CreateDescribeLivePrivateLineAreasResponse() (response *DescribeLivePrivateLineAreasResponse) {
	response = &DescribeLivePrivateLineAreasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
