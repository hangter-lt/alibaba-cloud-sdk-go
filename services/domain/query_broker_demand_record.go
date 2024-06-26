package domain

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

// QueryBrokerDemandRecord invokes the domain.QueryBrokerDemandRecord API synchronously
func (client *Client) QueryBrokerDemandRecord(request *QueryBrokerDemandRecordRequest) (response *QueryBrokerDemandRecordResponse, err error) {
	response = CreateQueryBrokerDemandRecordResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBrokerDemandRecordWithChan invokes the domain.QueryBrokerDemandRecord API asynchronously
func (client *Client) QueryBrokerDemandRecordWithChan(request *QueryBrokerDemandRecordRequest) (<-chan *QueryBrokerDemandRecordResponse, <-chan error) {
	responseChan := make(chan *QueryBrokerDemandRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryBrokerDemandRecord(request)
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

// QueryBrokerDemandRecordWithCallback invokes the domain.QueryBrokerDemandRecord API asynchronously
func (client *Client) QueryBrokerDemandRecordWithCallback(request *QueryBrokerDemandRecordRequest, callback func(response *QueryBrokerDemandRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryBrokerDemandRecordResponse
		var err error
		defer close(result)
		response, err = client.QueryBrokerDemandRecord(request)
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

// QueryBrokerDemandRecordRequest is the request struct for api QueryBrokerDemandRecord
type QueryBrokerDemandRecordRequest struct {
	*requests.RpcRequest
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	BizId       string           `position:"Query" name:"BizId"`
}

// QueryBrokerDemandRecordResponse is the response struct for api QueryBrokerDemandRecord
type QueryBrokerDemandRecordResponse struct {
	*responses.BaseResponse
	CurrentPageNum int                           `json:"CurrentPageNum" xml:"CurrentPageNum"`
	PageSize       int                           `json:"PageSize" xml:"PageSize"`
	RequestId      string                        `json:"RequestId" xml:"RequestId"`
	TotalPageNum   int                           `json:"TotalPageNum" xml:"TotalPageNum"`
	TotalItemNum   int                           `json:"TotalItemNum" xml:"TotalItemNum"`
	Data           DataInQueryBrokerDemandRecord `json:"Data" xml:"Data"`
}

// CreateQueryBrokerDemandRecordRequest creates a request to invoke QueryBrokerDemandRecord API
func CreateQueryBrokerDemandRecordRequest() (request *QueryBrokerDemandRecordRequest) {
	request = &QueryBrokerDemandRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "QueryBrokerDemandRecord", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryBrokerDemandRecordResponse creates a response to parse from QueryBrokerDemandRecord response
func CreateQueryBrokerDemandRecordResponse() (response *QueryBrokerDemandRecordResponse) {
	response = &QueryBrokerDemandRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
