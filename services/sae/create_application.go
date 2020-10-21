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

// CreateApplication invokes the sae.CreateApplication API synchronously
func (client *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
	response = CreateCreateApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateApplicationWithChan invokes the sae.CreateApplication API asynchronously
func (client *Client) CreateApplicationWithChan(request *CreateApplicationRequest) (<-chan *CreateApplicationResponse, <-chan error) {
	responseChan := make(chan *CreateApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateApplication(request)
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

// CreateApplicationWithCallback invokes the sae.CreateApplication API asynchronously
func (client *Client) CreateApplicationWithCallback(request *CreateApplicationRequest, callback func(response *CreateApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateApplicationResponse
		var err error
		defer close(result)
		response, err = client.CreateApplication(request)
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

// CreateApplicationRequest is the request struct for api CreateApplication
type CreateApplicationRequest struct {
	*requests.RoaRequest
	NasId                         string           `position:"Query" name:"NasId"`
	WebContainer                  string           `position:"Query" name:"WebContainer"`
	JarStartArgs                  string           `position:"Query" name:"JarStartArgs"`
	Memory                        requests.Integer `position:"Query" name:"Memory"`
	SlsConfigs                    string           `position:"Query" name:"SlsConfigs"`
	CommandArgs                   string           `position:"Query" name:"CommandArgs"`
	Readiness                     string           `position:"Query" name:"Readiness"`
	Timezone                      string           `position:"Query" name:"Timezone"`
	MountHost                     string           `position:"Query" name:"MountHost"`
	AutoConfig                    requests.Boolean `position:"Query" name:"AutoConfig"`
	Liveness                      string           `position:"Query" name:"Liveness"`
	SecurityGroupId               string           `position:"Query" name:"SecurityGroupId"`
	Envs                          string           `position:"Query" name:"Envs"`
	PhpArmsConfigLocation         string           `position:"Query" name:"PhpArmsConfigLocation"`
	PackageVersion                string           `position:"Query" name:"PackageVersion"`
	TomcatConfig                  string           `position:"Query" name:"TomcatConfig"`
	CustomHostAlias               string           `position:"Query" name:"CustomHostAlias"`
	Deploy                        requests.Boolean `position:"Query" name:"Deploy"`
	WarStartOptions               string           `position:"Query" name:"WarStartOptions"`
	JarStartOptions               string           `position:"Query" name:"JarStartOptions"`
	EdasContainerVersion          string           `position:"Query" name:"EdasContainerVersion"`
	AppName                       string           `position:"Query" name:"AppName"`
	NamespaceId                   string           `position:"Query" name:"NamespaceId"`
	PackageUrl                    string           `position:"Query" name:"PackageUrl"`
	TerminationGracePeriodSeconds requests.Integer `position:"Query" name:"TerminationGracePeriodSeconds"`
	ConfigMapMountDesc            string           `position:"Body" name:"ConfigMapMountDesc"`
	PhpConfig                     string           `position:"Body" name:"PhpConfig"`
	PreStop                       string           `position:"Query" name:"PreStop"`
	Replicas                      requests.Integer `position:"Query" name:"Replicas"`
	Cpu                           requests.Integer `position:"Query" name:"Cpu"`
	Command                       string           `position:"Query" name:"Command"`
	MountDesc                     string           `position:"Query" name:"MountDesc"`
	VSwitchId                     string           `position:"Query" name:"VSwitchId"`
	Jdk                           string           `position:"Query" name:"Jdk"`
	AppDescription                string           `position:"Query" name:"AppDescription"`
	VpcId                         string           `position:"Query" name:"VpcId"`
	ImageUrl                      string           `position:"Query" name:"ImageUrl"`
	PackageType                   string           `position:"Query" name:"PackageType"`
	PhpConfigLocation             string           `position:"Query" name:"PhpConfigLocation"`
	PostStart                     string           `position:"Query" name:"PostStart"`
}

// CreateApplicationResponse is the response struct for api CreateApplication
type CreateApplicationResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateApplicationRequest creates a request to invoke CreateApplication API
func CreateCreateApplicationRequest() (request *CreateApplicationRequest) {
	request = &CreateApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "CreateApplication", "/pop/v1/sam/app/createApplication", "serverless", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateApplicationResponse creates a response to parse from CreateApplication response
func CreateCreateApplicationResponse() (response *CreateApplicationResponse) {
	response = &CreateApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
