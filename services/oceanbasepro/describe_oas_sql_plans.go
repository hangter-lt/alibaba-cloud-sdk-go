package oceanbasepro

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

// DescribeOasSQLPlans invokes the oceanbasepro.DescribeOasSQLPlans API synchronously
func (client *Client) DescribeOasSQLPlans(request *DescribeOasSQLPlansRequest) (response *DescribeOasSQLPlansResponse, err error) {
	response = CreateDescribeOasSQLPlansResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOasSQLPlansWithChan invokes the oceanbasepro.DescribeOasSQLPlans API asynchronously
func (client *Client) DescribeOasSQLPlansWithChan(request *DescribeOasSQLPlansRequest) (<-chan *DescribeOasSQLPlansResponse, <-chan error) {
	responseChan := make(chan *DescribeOasSQLPlansResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOasSQLPlans(request)
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

// DescribeOasSQLPlansWithCallback invokes the oceanbasepro.DescribeOasSQLPlans API asynchronously
func (client *Client) DescribeOasSQLPlansWithCallback(request *DescribeOasSQLPlansRequest, callback func(response *DescribeOasSQLPlansResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOasSQLPlansResponse
		var err error
		defer close(result)
		response, err = client.DescribeOasSQLPlans(request)
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

// DescribeOasSQLPlansRequest is the request struct for api DescribeOasSQLPlans
type DescribeOasSQLPlansRequest struct {
	*requests.RpcRequest
	ReturnBriefInfo requests.Boolean `position:"Body" name:"ReturnBriefInfo"`
	StartTime       string           `position:"Body" name:"StartTime"`
	PlanUnionHash   string           `position:"Body" name:"PlanUnionHash"`
	DynamicSql      requests.Boolean `position:"Body" name:"DynamicSql"`
	TenantId        string           `position:"Body" name:"TenantId"`
	SqlId           string           `position:"Body" name:"SqlId"`
	EndTime         string           `position:"Body" name:"EndTime"`
	InstanceId      string           `position:"Body" name:"InstanceId"`
	DbName          string           `position:"Body" name:"DbName"`
	AcceptLanguage  string           `position:"Body" name:"AcceptLanguage"`
}

// DescribeOasSQLPlansResponse is the response struct for api DescribeOasSQLPlans
type DescribeOasSQLPlansResponse struct {
	*responses.BaseResponse
	RequestId string                          `json:"RequestId" xml:"RequestId"`
	Data      []DataItemInDescribeOasSQLPlans `json:"Data" xml:"Data"`
}

// CreateDescribeOasSQLPlansRequest creates a request to invoke DescribeOasSQLPlans API
func CreateDescribeOasSQLPlansRequest() (request *DescribeOasSQLPlansRequest) {
	request = &DescribeOasSQLPlansRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeOasSQLPlans", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeOasSQLPlansResponse creates a response to parse from DescribeOasSQLPlans response
func CreateDescribeOasSQLPlansResponse() (response *DescribeOasSQLPlansResponse) {
	response = &DescribeOasSQLPlansResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
