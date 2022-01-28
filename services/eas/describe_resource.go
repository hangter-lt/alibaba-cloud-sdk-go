package eas

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

// DescribeResource invokes the eas.DescribeResource API synchronously
func (client *Client) DescribeResource(request *DescribeResourceRequest) (response *DescribeResourceResponse, err error) {
	response = CreateDescribeResourceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeResourceWithChan invokes the eas.DescribeResource API asynchronously
func (client *Client) DescribeResourceWithChan(request *DescribeResourceRequest) (<-chan *DescribeResourceResponse, <-chan error) {
	responseChan := make(chan *DescribeResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResource(request)
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

// DescribeResourceWithCallback invokes the eas.DescribeResource API asynchronously
func (client *Client) DescribeResourceWithCallback(request *DescribeResourceRequest, callback func(response *DescribeResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourceResponse
		var err error
		defer close(result)
		response, err = client.DescribeResource(request)
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

// DescribeResourceRequest is the request struct for api DescribeResource
type DescribeResourceRequest struct {
	*requests.RoaRequest
	ResourceId string `position:"Path" name:"ResourceId"`
	ClusterId  string `position:"Path" name:"ClusterId"`
}

// DescribeResourceResponse is the response struct for api DescribeResource
type DescribeResourceResponse struct {
	*responses.BaseResponse
}

// CreateDescribeResourceRequest creates a request to invoke DescribeResource API
func CreateDescribeResourceRequest() (request *DescribeResourceRequest) {
	request = &DescribeResourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "DescribeResource", "/api/v2/resources/[ClusterId]/[ResourceId]", "eas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeResourceResponse creates a response to parse from DescribeResource response
func CreateDescribeResourceResponse() (response *DescribeResourceResponse) {
	response = &DescribeResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}