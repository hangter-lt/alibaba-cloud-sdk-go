package cloudauth

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

// DescribeSmartStatisticsPageList invokes the cloudauth.DescribeSmartStatisticsPageList API synchronously
func (client *Client) DescribeSmartStatisticsPageList(request *DescribeSmartStatisticsPageListRequest) (response *DescribeSmartStatisticsPageListResponse, err error) {
	response = CreateDescribeSmartStatisticsPageListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSmartStatisticsPageListWithChan invokes the cloudauth.DescribeSmartStatisticsPageList API asynchronously
func (client *Client) DescribeSmartStatisticsPageListWithChan(request *DescribeSmartStatisticsPageListRequest) (<-chan *DescribeSmartStatisticsPageListResponse, <-chan error) {
	responseChan := make(chan *DescribeSmartStatisticsPageListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSmartStatisticsPageList(request)
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

// DescribeSmartStatisticsPageListWithCallback invokes the cloudauth.DescribeSmartStatisticsPageList API asynchronously
func (client *Client) DescribeSmartStatisticsPageListWithCallback(request *DescribeSmartStatisticsPageListRequest, callback func(response *DescribeSmartStatisticsPageListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSmartStatisticsPageListResponse
		var err error
		defer close(result)
		response, err = client.DescribeSmartStatisticsPageList(request)
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

// DescribeSmartStatisticsPageListRequest is the request struct for api DescribeSmartStatisticsPageList
type DescribeSmartStatisticsPageListRequest struct {
	*requests.RpcRequest
	StartDate   string `position:"Query" name:"StartDate"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	PageSize    string `position:"Query" name:"PageSize"`
	CurrentPage string `position:"Query" name:"CurrentPage"`
	EndDate     string `position:"Query" name:"EndDate"`
	ServiceCode string `position:"Query" name:"ServiceCode"`
	SceneId     string `position:"Query" name:"SceneId"`
}

// DescribeSmartStatisticsPageListResponse is the response struct for api DescribeSmartStatisticsPageList
type DescribeSmartStatisticsPageListResponse struct {
	*responses.BaseResponse
	CurrentPage int                                          `json:"CurrentPage" xml:"CurrentPage"`
	TotalPage   int                                          `json:"TotalPage" xml:"TotalPage"`
	PageSize    int                                          `json:"PageSize" xml:"PageSize"`
	RequestId   string                                       `json:"RequestId" xml:"RequestId"`
	TotalCount  int                                          `json:"TotalCount" xml:"TotalCount"`
	Items       []ItemsItemInDescribeSmartStatisticsPageList `json:"Items" xml:"Items"`
}

// CreateDescribeSmartStatisticsPageListRequest creates a request to invoke DescribeSmartStatisticsPageList API
func CreateDescribeSmartStatisticsPageListRequest() (request *DescribeSmartStatisticsPageListRequest) {
	request = &DescribeSmartStatisticsPageListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeSmartStatisticsPageList", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSmartStatisticsPageListResponse creates a response to parse from DescribeSmartStatisticsPageList response
func CreateDescribeSmartStatisticsPageListResponse() (response *DescribeSmartStatisticsPageListResponse) {
	response = &DescribeSmartStatisticsPageListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
