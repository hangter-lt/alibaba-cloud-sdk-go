package cloudcallcenter

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

// QueryVnNavigationScripts invokes the cloudcallcenter.QueryVnNavigationScripts API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryvnnavigationscripts.html
func (client *Client) QueryVnNavigationScripts(request *QueryVnNavigationScriptsRequest) (response *QueryVnNavigationScriptsResponse, err error) {
	response = CreateQueryVnNavigationScriptsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryVnNavigationScriptsWithChan invokes the cloudcallcenter.QueryVnNavigationScripts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryvnnavigationscripts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryVnNavigationScriptsWithChan(request *QueryVnNavigationScriptsRequest) (<-chan *QueryVnNavigationScriptsResponse, <-chan error) {
	responseChan := make(chan *QueryVnNavigationScriptsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryVnNavigationScripts(request)
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

// QueryVnNavigationScriptsWithCallback invokes the cloudcallcenter.QueryVnNavigationScripts API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/queryvnnavigationscripts.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryVnNavigationScriptsWithCallback(request *QueryVnNavigationScriptsRequest, callback func(response *QueryVnNavigationScriptsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryVnNavigationScriptsResponse
		var err error
		defer close(result)
		response, err = client.QueryVnNavigationScripts(request)
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

// QueryVnNavigationScriptsRequest is the request struct for api QueryVnNavigationScripts
type QueryVnNavigationScriptsRequest struct {
	*requests.RpcRequest
	Type       string           `position:"Query" name:"Type"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	KeyWord    string           `position:"Query" name:"KeyWord"`
	CategoryId string           `position:"Query" name:"CategoryId"`
}

// QueryVnNavigationScriptsResponse is the response struct for api QueryVnNavigationScripts
type QueryVnNavigationScriptsResponse struct {
	*responses.BaseResponse
	RequestId         string             `json:"RequestId" xml:"RequestId"`
	TotalCount        int64              `json:"TotalCount" xml:"TotalCount"`
	PageNumber        int                `json:"PageNumber" xml:"PageNumber"`
	PageSize          int                `json:"PageSize" xml:"PageSize"`
	NavigationScripts []NavigationScript `json:"NavigationScripts" xml:"NavigationScripts"`
}

// CreateQueryVnNavigationScriptsRequest creates a request to invoke QueryVnNavigationScripts API
func CreateQueryVnNavigationScriptsRequest() (request *QueryVnNavigationScriptsRequest) {
	request = &QueryVnNavigationScriptsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "QueryVnNavigationScripts", "", "")
	request.Method = requests.GET
	return
}

// CreateQueryVnNavigationScriptsResponse creates a response to parse from QueryVnNavigationScripts response
func CreateQueryVnNavigationScriptsResponse() (response *QueryVnNavigationScriptsResponse) {
	response = &QueryVnNavigationScriptsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}