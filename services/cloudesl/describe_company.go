package cloudesl

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

// DescribeCompany invokes the cloudesl.DescribeCompany API synchronously
func (client *Client) DescribeCompany(request *DescribeCompanyRequest) (response *DescribeCompanyResponse, err error) {
	response = CreateDescribeCompanyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCompanyWithChan invokes the cloudesl.DescribeCompany API asynchronously
func (client *Client) DescribeCompanyWithChan(request *DescribeCompanyRequest) (<-chan *DescribeCompanyResponse, <-chan error) {
	responseChan := make(chan *DescribeCompanyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCompany(request)
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

// DescribeCompanyWithCallback invokes the cloudesl.DescribeCompany API asynchronously
func (client *Client) DescribeCompanyWithCallback(request *DescribeCompanyRequest, callback func(response *DescribeCompanyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCompanyResponse
		var err error
		defer close(result)
		response, err = client.DescribeCompany(request)
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

// DescribeCompanyRequest is the request struct for api DescribeCompany
type DescribeCompanyRequest struct {
	*requests.RpcRequest
}

// DescribeCompanyResponse is the response struct for api DescribeCompany
type DescribeCompanyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	CompanyId string `json:"CompanyId" xml:"CompanyId"`
	Platform  string `json:"Platform" xml:"Platform"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateDescribeCompanyRequest creates a request to invoke DescribeCompany API
func CreateDescribeCompanyRequest() (request *DescribeCompanyRequest) {
	request = &DescribeCompanyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2018-08-01", "DescribeCompany", "cloudesl", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCompanyResponse creates a response to parse from DescribeCompany response
func CreateDescribeCompanyResponse() (response *DescribeCompanyResponse) {
	response = &DescribeCompanyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}