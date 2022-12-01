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

// SearchTraces invokes the arms.SearchTraces API synchronously
func (client *Client) SearchTraces(request *SearchTracesRequest) (response *SearchTracesResponse, err error) {
	response = CreateSearchTracesResponse()
	err = client.DoAction(request, response)
	return
}

// SearchTracesWithChan invokes the arms.SearchTraces API asynchronously
func (client *Client) SearchTracesWithChan(request *SearchTracesRequest) (<-chan *SearchTracesResponse, <-chan error) {
	responseChan := make(chan *SearchTracesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchTraces(request)
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

// SearchTracesWithCallback invokes the arms.SearchTraces API asynchronously
func (client *Client) SearchTracesWithCallback(request *SearchTracesRequest, callback func(response *SearchTracesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchTracesResponse
		var err error
		defer close(result)
		response, err = client.SearchTraces(request)
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

// SearchTracesRequest is the request struct for api SearchTraces
type SearchTracesRequest struct {
	*requests.RpcRequest
	EndTime          requests.Integer                `position:"Query" name:"EndTime"`
	Pid              string                          `position:"Query" name:"Pid"`
	StartTime        requests.Integer                `position:"Query" name:"StartTime"`
	Reverse          requests.Boolean                `position:"Query" name:"Reverse"`
	MinDuration      requests.Integer                `position:"Query" name:"MinDuration"`
	ServiceIp        string                          `position:"Query" name:"ServiceIp"`
	ExclusionFilters *[]SearchTracesExclusionFilters `position:"Query" name:"ExclusionFilters"  type:"Repeated"`
	OperationName    string                          `position:"Query" name:"OperationName"`
	ServiceName      string                          `position:"Query" name:"ServiceName"`
	Tag              *[]SearchTracesTag              `position:"Query" name:"Tag"  type:"Repeated"`
}

// SearchTracesExclusionFilters is a repeated param struct in SearchTracesRequest
type SearchTracesExclusionFilters struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// SearchTracesTag is a repeated param struct in SearchTracesRequest
type SearchTracesTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// SearchTracesResponse is the response struct for api SearchTraces
type SearchTracesResponse struct {
	*responses.BaseResponse
	RequestId  string      `json:"RequestId" xml:"RequestId"`
	TraceInfos []TraceInfo `json:"TraceInfos" xml:"TraceInfos"`
}

// CreateSearchTracesRequest creates a request to invoke SearchTraces API
func CreateSearchTracesRequest() (request *SearchTracesRequest) {
	request = &SearchTracesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "SearchTraces", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSearchTracesResponse creates a response to parse from SearchTraces response
func CreateSearchTracesResponse() (response *SearchTracesResponse) {
	response = &SearchTracesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
