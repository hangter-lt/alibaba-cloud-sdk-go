package dcdn

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

// DescribeDcdnCertificateList invokes the dcdn.DescribeDcdnCertificateList API synchronously
func (client *Client) DescribeDcdnCertificateList(request *DescribeDcdnCertificateListRequest) (response *DescribeDcdnCertificateListResponse, err error) {
	response = CreateDescribeDcdnCertificateListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnCertificateListWithChan invokes the dcdn.DescribeDcdnCertificateList API asynchronously
func (client *Client) DescribeDcdnCertificateListWithChan(request *DescribeDcdnCertificateListRequest) (<-chan *DescribeDcdnCertificateListResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnCertificateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnCertificateList(request)
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

// DescribeDcdnCertificateListWithCallback invokes the dcdn.DescribeDcdnCertificateList API asynchronously
func (client *Client) DescribeDcdnCertificateListWithCallback(request *DescribeDcdnCertificateListRequest, callback func(response *DescribeDcdnCertificateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnCertificateListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnCertificateList(request)
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

// DescribeDcdnCertificateListRequest is the request struct for api DescribeDcdnCertificateList
type DescribeDcdnCertificateListRequest struct {
	*requests.RpcRequest
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnCertificateListResponse is the response struct for api DescribeDcdnCertificateList
type DescribeDcdnCertificateListResponse struct {
	*responses.BaseResponse
	RequestId            string               `json:"RequestId" xml:"RequestId"`
	CertificateListModel CertificateListModel `json:"CertificateListModel" xml:"CertificateListModel"`
}

// CreateDescribeDcdnCertificateListRequest creates a request to invoke DescribeDcdnCertificateList API
func CreateDescribeDcdnCertificateListRequest() (request *DescribeDcdnCertificateListRequest) {
	request = &DescribeDcdnCertificateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnCertificateList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDcdnCertificateListResponse creates a response to parse from DescribeDcdnCertificateList response
func CreateDescribeDcdnCertificateListResponse() (response *DescribeDcdnCertificateListResponse) {
	response = &DescribeDcdnCertificateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
