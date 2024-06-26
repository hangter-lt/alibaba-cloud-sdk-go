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

// CreateLivePrivateLine invokes the live.CreateLivePrivateLine API synchronously
func (client *Client) CreateLivePrivateLine(request *CreateLivePrivateLineRequest) (response *CreateLivePrivateLineResponse, err error) {
	response = CreateCreateLivePrivateLineResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLivePrivateLineWithChan invokes the live.CreateLivePrivateLine API asynchronously
func (client *Client) CreateLivePrivateLineWithChan(request *CreateLivePrivateLineRequest) (<-chan *CreateLivePrivateLineResponse, <-chan error) {
	responseChan := make(chan *CreateLivePrivateLineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLivePrivateLine(request)
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

// CreateLivePrivateLineWithCallback invokes the live.CreateLivePrivateLine API asynchronously
func (client *Client) CreateLivePrivateLineWithCallback(request *CreateLivePrivateLineRequest, callback func(response *CreateLivePrivateLineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLivePrivateLineResponse
		var err error
		defer close(result)
		response, err = client.CreateLivePrivateLine(request)
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

// CreateLivePrivateLineRequest is the request struct for api CreateLivePrivateLine
type CreateLivePrivateLineRequest struct {
	*requests.RpcRequest
	MaxBandwidth     string           `position:"Query" name:"MaxBandwidth"`
	Reuse            string           `position:"Query" name:"Reuse"`
	AccelerationArea string           `position:"Query" name:"AccelerationArea"`
	AppName          string           `position:"Query" name:"AppName"`
	StreamName       string           `position:"Query" name:"StreamName"`
	DomainName       string           `position:"Query" name:"DomainName"`
	OwnerId          requests.Integer `position:"Query" name:"OwnerId"`
	VideoCenter      string           `position:"Query" name:"VideoCenter"`
	AccelerationType string           `position:"Query" name:"AccelerationType"`
	InstanceId       string           `position:"Query" name:"InstanceId"`
}

// CreateLivePrivateLineResponse is the response struct for api CreateLivePrivateLine
type CreateLivePrivateLineResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateLivePrivateLineRequest creates a request to invoke CreateLivePrivateLine API
func CreateCreateLivePrivateLineRequest() (request *CreateLivePrivateLineRequest) {
	request = &CreateLivePrivateLineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "CreateLivePrivateLine", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLivePrivateLineResponse creates a response to parse from CreateLivePrivateLine response
func CreateCreateLivePrivateLineResponse() (response *CreateLivePrivateLineResponse) {
	response = &CreateLivePrivateLineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
