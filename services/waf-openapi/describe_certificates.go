package waf_openapi

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

// DescribeCertificates invokes the waf_openapi.DescribeCertificates API synchronously
func (client *Client) DescribeCertificates(request *DescribeCertificatesRequest) (response *DescribeCertificatesResponse, err error) {
	response = CreateDescribeCertificatesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCertificatesWithChan invokes the waf_openapi.DescribeCertificates API asynchronously
func (client *Client) DescribeCertificatesWithChan(request *DescribeCertificatesRequest) (<-chan *DescribeCertificatesResponse, <-chan error) {
	responseChan := make(chan *DescribeCertificatesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCertificates(request)
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

// DescribeCertificatesWithCallback invokes the waf_openapi.DescribeCertificates API asynchronously
func (client *Client) DescribeCertificatesWithCallback(request *DescribeCertificatesRequest, callback func(response *DescribeCertificatesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCertificatesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCertificates(request)
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

// DescribeCertificatesRequest is the request struct for api DescribeCertificates
type DescribeCertificatesRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Lang            string `position:"Query" name:"Lang"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	Domain          string `position:"Query" name:"Domain"`
}

// DescribeCertificatesResponse is the response struct for api DescribeCertificates
type DescribeCertificatesResponse struct {
	*responses.BaseResponse
	RequestId    string        `json:"RequestId" xml:"RequestId"`
	Certificates []Certificate `json:"Certificates" xml:"Certificates"`
}

// CreateDescribeCertificatesRequest creates a request to invoke DescribeCertificates API
func CreateDescribeCertificatesRequest() (request *DescribeCertificatesRequest) {
	request = &DescribeCertificatesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DescribeCertificates", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCertificatesResponse creates a response to parse from DescribeCertificates response
func CreateDescribeCertificatesResponse() (response *DescribeCertificatesResponse) {
	response = &DescribeCertificatesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
