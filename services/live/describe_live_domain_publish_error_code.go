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

// DescribeLiveDomainPublishErrorCode invokes the live.DescribeLiveDomainPublishErrorCode API synchronously
func (client *Client) DescribeLiveDomainPublishErrorCode(request *DescribeLiveDomainPublishErrorCodeRequest) (response *DescribeLiveDomainPublishErrorCodeResponse, err error) {
	response = CreateDescribeLiveDomainPublishErrorCodeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveDomainPublishErrorCodeWithChan invokes the live.DescribeLiveDomainPublishErrorCode API asynchronously
func (client *Client) DescribeLiveDomainPublishErrorCodeWithChan(request *DescribeLiveDomainPublishErrorCodeRequest) (<-chan *DescribeLiveDomainPublishErrorCodeResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveDomainPublishErrorCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveDomainPublishErrorCode(request)
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

// DescribeLiveDomainPublishErrorCodeWithCallback invokes the live.DescribeLiveDomainPublishErrorCode API asynchronously
func (client *Client) DescribeLiveDomainPublishErrorCodeWithCallback(request *DescribeLiveDomainPublishErrorCodeRequest, callback func(response *DescribeLiveDomainPublishErrorCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveDomainPublishErrorCodeResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveDomainPublishErrorCode(request)
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

// DescribeLiveDomainPublishErrorCodeRequest is the request struct for api DescribeLiveDomainPublishErrorCode
type DescribeLiveDomainPublishErrorCodeRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	AppName    string           `position:"Query" name:"AppName"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveDomainPublishErrorCodeResponse is the response struct for api DescribeLiveDomainPublishErrorCode
type DescribeLiveDomainPublishErrorCodeResponse struct {
	*responses.BaseResponse
	DataInterval     string `json:"DataInterval" xml:"DataInterval"`
	DomainName       string `json:"DomainName" xml:"DomainName"`
	EndTime          string `json:"EndTime" xml:"EndTime"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	StartTime        string `json:"StartTime" xml:"StartTime"`
	RealTimeCodeData []Rtcd `json:"RealTimeCodeData" xml:"RealTimeCodeData"`
}

// CreateDescribeLiveDomainPublishErrorCodeRequest creates a request to invoke DescribeLiveDomainPublishErrorCode API
func CreateDescribeLiveDomainPublishErrorCodeRequest() (request *DescribeLiveDomainPublishErrorCodeRequest) {
	request = &DescribeLiveDomainPublishErrorCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainPublishErrorCode", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveDomainPublishErrorCodeResponse creates a response to parse from DescribeLiveDomainPublishErrorCode response
func CreateDescribeLiveDomainPublishErrorCodeResponse() (response *DescribeLiveDomainPublishErrorCodeResponse) {
	response = &DescribeLiveDomainPublishErrorCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
