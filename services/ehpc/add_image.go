package ehpc

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

// AddImage invokes the ehpc.AddImage API synchronously
func (client *Client) AddImage(request *AddImageRequest) (response *AddImageResponse, err error) {
	response = CreateAddImageResponse()
	err = client.DoAction(request, response)
	return
}

// AddImageWithChan invokes the ehpc.AddImage API asynchronously
func (client *Client) AddImageWithChan(request *AddImageRequest) (<-chan *AddImageResponse, <-chan error) {
	responseChan := make(chan *AddImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddImage(request)
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

// AddImageWithCallback invokes the ehpc.AddImage API asynchronously
func (client *Client) AddImageWithCallback(request *AddImageRequest, callback func(response *AddImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddImageResponse
		var err error
		defer close(result)
		response, err = client.AddImage(request)
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

// AddImageRequest is the request struct for api AddImage
type AddImageRequest struct {
	*requests.RpcRequest
	ImageType          string                     `position:"Query" name:"ImageType"`
	Description        string                     `position:"Query" name:"Description"`
	Version            string                     `position:"Query" name:"Version"`
	Name               string                     `position:"Query" name:"Name"`
	ContainerImageSpec AddImageContainerImageSpec `position:"Query" name:"ContainerImageSpec"  type:"Struct"`
	VMImageSpec        AddImageVMImageSpec        `position:"Query" name:"VMImageSpec"  type:"Struct"`
}

// AddImageContainerImageSpec is a repeated param struct in AddImageRequest
type AddImageContainerImageSpec struct {
	IsACREnterprise    string                                       `name:"IsACREnterprise"`
	RegistryUrl        string                                       `name:"RegistryUrl"`
	RegistryCredential AddImageContainerImageSpecRegistryCredential `name:"RegistryCredential" type:"Struct"`
	RegistryCriId      string                                       `name:"RegistryCriId"`
	IsACRRegistry      string                                       `name:"IsACRRegistry"`
}

// AddImageVMImageSpec is a repeated param struct in AddImageRequest
type AddImageVMImageSpec struct {
	ImageId string `name:"ImageId"`
}

// AddImageContainerImageSpecRegistryCredential is a repeated param struct in AddImageRequest
type AddImageContainerImageSpecRegistryCredential struct {
	Server   string `name:"Server"`
	Password string `name:"Password"`
	UserName string `name:"UserName"`
}

// AddImageResponse is the response struct for api AddImage
type AddImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ImageId   string `json:"ImageId" xml:"ImageId"`
	AppId     string `json:"AppId" xml:"AppId"`
}

// CreateAddImageRequest creates a request to invoke AddImage API
func CreateAddImageRequest() (request *AddImageRequest) {
	request = &AddImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2023-07-01", "AddImage", "ehs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddImageResponse creates a response to parse from AddImage response
func CreateAddImageResponse() (response *AddImageResponse) {
	response = &AddImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
