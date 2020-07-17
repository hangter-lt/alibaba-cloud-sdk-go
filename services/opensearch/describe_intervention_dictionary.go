package opensearch

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

// DescribeInterventionDictionary invokes the opensearch.DescribeInterventionDictionary API synchronously
// api document: https://help.aliyun.com/api/opensearch/describeinterventiondictionary.html
func (client *Client) DescribeInterventionDictionary(request *DescribeInterventionDictionaryRequest) (response *DescribeInterventionDictionaryResponse, err error) {
	response = CreateDescribeInterventionDictionaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInterventionDictionaryWithChan invokes the opensearch.DescribeInterventionDictionary API asynchronously
// api document: https://help.aliyun.com/api/opensearch/describeinterventiondictionary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInterventionDictionaryWithChan(request *DescribeInterventionDictionaryRequest) (<-chan *DescribeInterventionDictionaryResponse, <-chan error) {
	responseChan := make(chan *DescribeInterventionDictionaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInterventionDictionary(request)
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

// DescribeInterventionDictionaryWithCallback invokes the opensearch.DescribeInterventionDictionary API asynchronously
// api document: https://help.aliyun.com/api/opensearch/describeinterventiondictionary.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInterventionDictionaryWithCallback(request *DescribeInterventionDictionaryRequest, callback func(response *DescribeInterventionDictionaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInterventionDictionaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeInterventionDictionary(request)
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

// DescribeInterventionDictionaryRequest is the request struct for api DescribeInterventionDictionary
type DescribeInterventionDictionaryRequest struct {
	*requests.RoaRequest
	Name string `position:"Path" name:"name"`
}

// DescribeInterventionDictionaryResponse is the response struct for api DescribeInterventionDictionary
type DescribeInterventionDictionaryResponse struct {
	*responses.BaseResponse
	RequestId string                                 `json:"requestId" xml:"requestId"`
	Result    ResultInDescribeInterventionDictionary `json:"result" xml:"result"`
}

// CreateDescribeInterventionDictionaryRequest creates a request to invoke DescribeInterventionDictionary API
func CreateDescribeInterventionDictionaryRequest() (request *DescribeInterventionDictionaryRequest) {
	request = &DescribeInterventionDictionaryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeInterventionDictionary", "/v4/openapi/intervention-dictionaries/[name]", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeInterventionDictionaryResponse creates a response to parse from DescribeInterventionDictionary response
func CreateDescribeInterventionDictionaryResponse() (response *DescribeInterventionDictionaryResponse) {
	response = &DescribeInterventionDictionaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}