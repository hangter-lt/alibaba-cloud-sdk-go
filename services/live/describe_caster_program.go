package live

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

// DescribeCasterProgram invokes the live.DescribeCasterProgram API synchronously
func (client *Client) DescribeCasterProgram(request *DescribeCasterProgramRequest) (response *DescribeCasterProgramResponse, err error) {
	response = CreateDescribeCasterProgramResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCasterProgramWithChan invokes the live.DescribeCasterProgram API asynchronously
func (client *Client) DescribeCasterProgramWithChan(request *DescribeCasterProgramRequest) (<-chan *DescribeCasterProgramResponse, <-chan error) {
	responseChan := make(chan *DescribeCasterProgramResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCasterProgram(request)
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

// DescribeCasterProgramWithCallback invokes the live.DescribeCasterProgram API asynchronously
func (client *Client) DescribeCasterProgramWithCallback(request *DescribeCasterProgramRequest, callback func(response *DescribeCasterProgramResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCasterProgramResponse
		var err error
		defer close(result)
		response, err = client.DescribeCasterProgram(request)
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

// DescribeCasterProgramRequest is the request struct for api DescribeCasterProgram
type DescribeCasterProgramRequest struct {
	*requests.RpcRequest
	StartTime   string           `position:"Query" name:"StartTime"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CasterId    string           `position:"Query" name:"CasterId"`
	EpisodeType string           `position:"Query" name:"EpisodeType"`
	EndTime     string           `position:"Query" name:"EndTime"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	EpisodeId   string           `position:"Query" name:"EpisodeId"`
	Status      requests.Integer `position:"Query" name:"Status"`
}

// DescribeCasterProgramResponse is the response struct for api DescribeCasterProgram
type DescribeCasterProgramResponse struct {
	*responses.BaseResponse
	CasterId      string   `json:"CasterId" xml:"CasterId"`
	ProgramEffect int      `json:"ProgramEffect" xml:"ProgramEffect"`
	ProgramName   string   `json:"ProgramName" xml:"ProgramName"`
	RequestId     string   `json:"RequestId" xml:"RequestId"`
	Total         int      `json:"Total" xml:"Total"`
	Episodes      Episodes `json:"Episodes" xml:"Episodes"`
}

// CreateDescribeCasterProgramRequest creates a request to invoke DescribeCasterProgram API
func CreateDescribeCasterProgramRequest() (request *DescribeCasterProgramRequest) {
	request = &DescribeCasterProgramRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeCasterProgram", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCasterProgramResponse creates a response to parse from DescribeCasterProgram response
func CreateDescribeCasterProgramResponse() (response *DescribeCasterProgramResponse) {
	response = &DescribeCasterProgramResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
