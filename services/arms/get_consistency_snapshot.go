package arms

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

// GetConsistencySnapshot invokes the arms.GetConsistencySnapshot API synchronously
func (client *Client) GetConsistencySnapshot(request *GetConsistencySnapshotRequest) (response *GetConsistencySnapshotResponse, err error) {
	response = CreateGetConsistencySnapshotResponse()
	err = client.DoAction(request, response)
	return
}

// GetConsistencySnapshotWithChan invokes the arms.GetConsistencySnapshot API asynchronously
func (client *Client) GetConsistencySnapshotWithChan(request *GetConsistencySnapshotRequest) (<-chan *GetConsistencySnapshotResponse, <-chan error) {
	responseChan := make(chan *GetConsistencySnapshotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetConsistencySnapshot(request)
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

// GetConsistencySnapshotWithCallback invokes the arms.GetConsistencySnapshot API asynchronously
func (client *Client) GetConsistencySnapshotWithCallback(request *GetConsistencySnapshotRequest, callback func(response *GetConsistencySnapshotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetConsistencySnapshotResponse
		var err error
		defer close(result)
		response, err = client.GetConsistencySnapshot(request)
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

// GetConsistencySnapshotRequest is the request struct for api GetConsistencySnapshot
type GetConsistencySnapshotRequest struct {
	*requests.RpcRequest
	CurrentTimestamp requests.Integer `position:"Query" name:"CurrentTimestamp"`
	AppType          string           `position:"Query" name:"AppType"`
	Pid              string           `position:"Query" name:"Pid"`
	ProxyUserId      string           `position:"Query" name:"ProxyUserId"`
}

// GetConsistencySnapshotResponse is the response struct for api GetConsistencySnapshot
type GetConsistencySnapshotResponse struct {
	*responses.BaseResponse
	ConsistencyResultJsonStr string `json:"ConsistencyResultJsonStr" xml:"ConsistencyResultJsonStr"`
	RequestId                string `json:"RequestId" xml:"RequestId"`
}

// CreateGetConsistencySnapshotRequest creates a request to invoke GetConsistencySnapshot API
func CreateGetConsistencySnapshotRequest() (request *GetConsistencySnapshotRequest) {
	request = &GetConsistencySnapshotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2021-05-19", "GetConsistencySnapshot", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetConsistencySnapshotResponse creates a response to parse from GetConsistencySnapshot response
func CreateGetConsistencySnapshotResponse() (response *GetConsistencySnapshotResponse) {
	response = &GetConsistencySnapshotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
