package dms_enterprise

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

// GetSparkJobLog invokes the dms_enterprise.GetSparkJobLog API synchronously
func (client *Client) GetSparkJobLog(request *GetSparkJobLogRequest) (response *GetSparkJobLogResponse, err error) {
	response = CreateGetSparkJobLogResponse()
	err = client.DoAction(request, response)
	return
}

// GetSparkJobLogWithChan invokes the dms_enterprise.GetSparkJobLog API asynchronously
func (client *Client) GetSparkJobLogWithChan(request *GetSparkJobLogRequest) (<-chan *GetSparkJobLogResponse, <-chan error) {
	responseChan := make(chan *GetSparkJobLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSparkJobLog(request)
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

// GetSparkJobLogWithCallback invokes the dms_enterprise.GetSparkJobLog API asynchronously
func (client *Client) GetSparkJobLogWithCallback(request *GetSparkJobLogRequest, callback func(response *GetSparkJobLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSparkJobLogResponse
		var err error
		defer close(result)
		response, err = client.GetSparkJobLog(request)
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

// GetSparkJobLogRequest is the request struct for api GetSparkJobLog
type GetSparkJobLogRequest struct {
	*requests.RpcRequest
	JobId requests.Integer `position:"Query" name:"JobId"`
	Tid   requests.Integer `position:"Query" name:"Tid"`
}

// GetSparkJobLogResponse is the response struct for api GetSparkJobLog
type GetSparkJobLogResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	Log          string `json:"Log" xml:"Log"`
}

// CreateGetSparkJobLogRequest creates a request to invoke GetSparkJobLog API
func CreateGetSparkJobLogRequest() (request *GetSparkJobLogRequest) {
	request = &GetSparkJobLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetSparkJobLog", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetSparkJobLogResponse creates a response to parse from GetSparkJobLog response
func CreateGetSparkJobLogResponse() (response *GetSparkJobLogResponse) {
	response = &GetSparkJobLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}