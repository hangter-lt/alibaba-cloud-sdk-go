package gpdb

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

// DescribeDiagnosisDimensions invokes the gpdb.DescribeDiagnosisDimensions API synchronously
func (client *Client) DescribeDiagnosisDimensions(request *DescribeDiagnosisDimensionsRequest) (response *DescribeDiagnosisDimensionsResponse, err error) {
	response = CreateDescribeDiagnosisDimensionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDiagnosisDimensionsWithChan invokes the gpdb.DescribeDiagnosisDimensions API asynchronously
func (client *Client) DescribeDiagnosisDimensionsWithChan(request *DescribeDiagnosisDimensionsRequest) (<-chan *DescribeDiagnosisDimensionsResponse, <-chan error) {
	responseChan := make(chan *DescribeDiagnosisDimensionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDiagnosisDimensions(request)
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

// DescribeDiagnosisDimensionsWithCallback invokes the gpdb.DescribeDiagnosisDimensions API asynchronously
func (client *Client) DescribeDiagnosisDimensionsWithCallback(request *DescribeDiagnosisDimensionsRequest, callback func(response *DescribeDiagnosisDimensionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDiagnosisDimensionsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDiagnosisDimensions(request)
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

// DescribeDiagnosisDimensionsRequest is the request struct for api DescribeDiagnosisDimensions
type DescribeDiagnosisDimensionsRequest struct {
	*requests.RpcRequest
	DBInstanceId string `position:"Query" name:"DBInstanceId"`
}

// DescribeDiagnosisDimensionsResponse is the response struct for api DescribeDiagnosisDimensions
type DescribeDiagnosisDimensionsResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Databases []string `json:"Databases" xml:"Databases"`
	UserNames []string `json:"UserNames" xml:"UserNames"`
}

// CreateDescribeDiagnosisDimensionsRequest creates a request to invoke DescribeDiagnosisDimensions API
func CreateDescribeDiagnosisDimensionsRequest() (request *DescribeDiagnosisDimensionsRequest) {
	request = &DescribeDiagnosisDimensionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "DescribeDiagnosisDimensions", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDiagnosisDimensionsResponse creates a response to parse from DescribeDiagnosisDimensions response
func CreateDescribeDiagnosisDimensionsResponse() (response *DescribeDiagnosisDimensionsResponse) {
	response = &DescribeDiagnosisDimensionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
