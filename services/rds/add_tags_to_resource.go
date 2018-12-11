package rds

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

// AddTagsToResource invokes the rds.AddTagsToResource API synchronously
// api document: https://help.aliyun.com/api/rds/addtagstoresource.html
func (client *Client) AddTagsToResource(request *AddTagsToResourceRequest) (response *AddTagsToResourceResponse, err error) {
	response = CreateAddTagsToResourceResponse()
	err = client.DoAction(request, response)
	return
}

// AddTagsToResourceWithChan invokes the rds.AddTagsToResource API asynchronously
// api document: https://help.aliyun.com/api/rds/addtagstoresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTagsToResourceWithChan(request *AddTagsToResourceRequest) (<-chan *AddTagsToResourceResponse, <-chan error) {
	responseChan := make(chan *AddTagsToResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTagsToResource(request)
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

// AddTagsToResourceWithCallback invokes the rds.AddTagsToResource API asynchronously
// api document: https://help.aliyun.com/api/rds/addtagstoresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddTagsToResourceWithCallback(request *AddTagsToResourceRequest, callback func(response *AddTagsToResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTagsToResourceResponse
		var err error
		defer close(result)
		response, err = client.AddTagsToResource(request)
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

// AddTagsToResourceRequest is the request struct for api AddTagsToResource
type AddTagsToResourceRequest struct {
	*requests.RpcRequest
	Tag4Value            string           `position:"Query" name:"Tag.4.value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Tag2Key              string           `position:"Query" name:"Tag.2.key"`
	Tag5Key              string           `position:"Query" name:"Tag.5.key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Tag3Key              string           `position:"Query" name:"Tag.3.key"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Value            string           `position:"Query" name:"Tag.5.value"`
	Tags                 string           `position:"Query" name:"Tags"`
	Tag1Key              string           `position:"Query" name:"Tag.1.key"`
	Tag1Value            string           `position:"Query" name:"Tag.1.value"`
	Tag2Value            string           `position:"Query" name:"Tag.2.value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.key"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	Tag3Value            string           `position:"Query" name:"Tag.3.value"`
	ProxyId              string           `position:"Query" name:"proxyId"`
}

// AddTagsToResourceResponse is the response struct for api AddTagsToResource
type AddTagsToResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddTagsToResourceRequest creates a request to invoke AddTagsToResource API
func CreateAddTagsToResourceRequest() (request *AddTagsToResourceRequest) {
	request = &AddTagsToResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "AddTagsToResource", "Rds", "openAPI")
	return
}

// CreateAddTagsToResourceResponse creates a response to parse from AddTagsToResource response
func CreateAddTagsToResourceResponse() (response *AddTagsToResourceResponse) {
	response = &AddTagsToResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
