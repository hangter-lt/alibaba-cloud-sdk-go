package kms

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

// DescribeWhiteBoxKey invokes the kms.DescribeWhiteBoxKey API synchronously
func (client *Client) DescribeWhiteBoxKey(request *DescribeWhiteBoxKeyRequest) (response *DescribeWhiteBoxKeyResponse, err error) {
	response = CreateDescribeWhiteBoxKeyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWhiteBoxKeyWithChan invokes the kms.DescribeWhiteBoxKey API asynchronously
func (client *Client) DescribeWhiteBoxKeyWithChan(request *DescribeWhiteBoxKeyRequest) (<-chan *DescribeWhiteBoxKeyResponse, <-chan error) {
	responseChan := make(chan *DescribeWhiteBoxKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWhiteBoxKey(request)
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

// DescribeWhiteBoxKeyWithCallback invokes the kms.DescribeWhiteBoxKey API asynchronously
func (client *Client) DescribeWhiteBoxKeyWithCallback(request *DescribeWhiteBoxKeyRequest, callback func(response *DescribeWhiteBoxKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWhiteBoxKeyResponse
		var err error
		defer close(result)
		response, err = client.DescribeWhiteBoxKey(request)
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

// DescribeWhiteBoxKeyRequest is the request struct for api DescribeWhiteBoxKey
type DescribeWhiteBoxKeyRequest struct {
	*requests.RpcRequest
	KeyId string `position:"Query" name:"KeyId"`
}

// DescribeWhiteBoxKeyResponse is the response struct for api DescribeWhiteBoxKey
type DescribeWhiteBoxKeyResponse struct {
	*responses.BaseResponse
	RequestId             string `json:"RequestId" xml:"RequestId"`
	KeyId                 string `json:"KeyId" xml:"KeyId"`
	Arn                   string `json:"Arn" xml:"Arn"`
	Name                  string `json:"Name" xml:"Name"`
	Algorithm             string `json:"Algorithm" xml:"Algorithm"`
	Status                string `json:"Status" xml:"Status"`
	EnableObtainKeyStatus string `json:"EnableObtainKeyStatus" xml:"EnableObtainKeyStatus"`
	Description           string `json:"Description" xml:"Description"`
	CreateTime            string `json:"CreateTime" xml:"CreateTime"`
}

// CreateDescribeWhiteBoxKeyRequest creates a request to invoke DescribeWhiteBoxKey API
func CreateDescribeWhiteBoxKeyRequest() (request *DescribeWhiteBoxKeyRequest) {
	request = &DescribeWhiteBoxKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "DescribeWhiteBoxKey", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeWhiteBoxKeyResponse creates a response to parse from DescribeWhiteBoxKey response
func CreateDescribeWhiteBoxKeyResponse() (response *DescribeWhiteBoxKeyResponse) {
	response = &DescribeWhiteBoxKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}