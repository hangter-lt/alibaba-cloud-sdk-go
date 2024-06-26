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

// AIGCFaceVerify invokes the cloudauth.AIGCFaceVerify API synchronously
func (client *Client) AIGCFaceVerify(request *AIGCFaceVerifyRequest) (response *AIGCFaceVerifyResponse, err error) {
	response = CreateAIGCFaceVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// AIGCFaceVerifyWithChan invokes the cloudauth.AIGCFaceVerify API asynchronously
func (client *Client) AIGCFaceVerifyWithChan(request *AIGCFaceVerifyRequest) (<-chan *AIGCFaceVerifyResponse, <-chan error) {
	responseChan := make(chan *AIGCFaceVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AIGCFaceVerify(request)
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

// AIGCFaceVerifyWithCallback invokes the cloudauth.AIGCFaceVerify API asynchronously
func (client *Client) AIGCFaceVerifyWithCallback(request *AIGCFaceVerifyRequest, callback func(response *AIGCFaceVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AIGCFaceVerifyResponse
		var err error
		defer close(result)
		response, err = client.AIGCFaceVerify(request)
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

// AIGCFaceVerifyRequest is the request struct for api AIGCFaceVerify
type AIGCFaceVerifyRequest struct {
	*requests.RpcRequest
	ProductCode            string           `position:"Query" name:"ProductCode"`
	OssObjectName          string           `position:"Query" name:"OssObjectName"`
	FaceContrastPicture    string           `position:"Body" name:"FaceContrastPicture"`
	OuterOrderNo           string           `position:"Query" name:"OuterOrderNo"`
	FaceContrastPictureUrl string           `position:"Query" name:"FaceContrastPictureUrl"`
	SceneId                requests.Integer `position:"Query" name:"SceneId"`
	OssBucketName          string           `position:"Query" name:"OssBucketName"`
}

// AIGCFaceVerifyResponse is the response struct for api AIGCFaceVerify
type AIGCFaceVerifyResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Message      string       `json:"Message" xml:"Message"`
	Code         string       `json:"Code" xml:"Code"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateAIGCFaceVerifyRequest creates a request to invoke AIGCFaceVerify API
func CreateAIGCFaceVerifyRequest() (request *AIGCFaceVerifyRequest) {
	request = &AIGCFaceVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "AIGCFaceVerify", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAIGCFaceVerifyResponse creates a response to parse from AIGCFaceVerify response
func CreateAIGCFaceVerifyResponse() (response *AIGCFaceVerifyResponse) {
	response = &AIGCFaceVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
