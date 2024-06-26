package sae

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

// GetWarningEventMetric invokes the sae.GetWarningEventMetric API synchronously
func (client *Client) GetWarningEventMetric(request *GetWarningEventMetricRequest) (response *GetWarningEventMetricResponse, err error) {
	response = CreateGetWarningEventMetricResponse()
	err = client.DoAction(request, response)
	return
}

// GetWarningEventMetricWithChan invokes the sae.GetWarningEventMetric API asynchronously
func (client *Client) GetWarningEventMetricWithChan(request *GetWarningEventMetricRequest) (<-chan *GetWarningEventMetricResponse, <-chan error) {
	responseChan := make(chan *GetWarningEventMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWarningEventMetric(request)
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

// GetWarningEventMetricWithCallback invokes the sae.GetWarningEventMetric API asynchronously
func (client *Client) GetWarningEventMetricWithCallback(request *GetWarningEventMetricRequest, callback func(response *GetWarningEventMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWarningEventMetricResponse
		var err error
		defer close(result)
		response, err = client.GetWarningEventMetric(request)
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

// GetWarningEventMetricRequest is the request struct for api GetWarningEventMetric
type GetWarningEventMetricRequest struct {
	*requests.RoaRequest
	AppSource   string           `position:"Query" name:"AppSource"`
	CpuStrategy string           `position:"Query" name:"CpuStrategy"`
	Limit       requests.Integer `position:"Query" name:"Limit"`
	EndTime     requests.Integer `position:"Query" name:"EndTime"`
	StartTime   requests.Integer `position:"Query" name:"StartTime"`
}

// GetWarningEventMetricResponse is the response struct for api GetWarningEventMetric
type GetWarningEventMetricResponse struct {
	*responses.BaseResponse
	Message   string           `json:"Message" xml:"Message"`
	RequestId string           `json:"RequestId" xml:"RequestId"`
	Code      string           `json:"Code" xml:"Code"`
	Success   bool             `json:"Success" xml:"Success"`
	Data      []EventMetricDto `json:"Data" xml:"Data"`
}

// CreateGetWarningEventMetricRequest creates a request to invoke GetWarningEventMetric API
func CreateGetWarningEventMetricRequest() (request *GetWarningEventMetricRequest) {
	request = &GetWarningEventMetricRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "GetWarningEventMetric", "/pop/v1/sam/getWarningEventMetric", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetWarningEventMetricResponse creates a response to parse from GetWarningEventMetric response
func CreateGetWarningEventMetricResponse() (response *GetWarningEventMetricResponse) {
	response = &GetWarningEventMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
