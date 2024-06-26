package computenestsupplier

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

// CreateServiceInstance invokes the computenestsupplier.CreateServiceInstance API synchronously
func (client *Client) CreateServiceInstance(request *CreateServiceInstanceRequest) (response *CreateServiceInstanceResponse, err error) {
	response = CreateCreateServiceInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateServiceInstanceWithChan invokes the computenestsupplier.CreateServiceInstance API asynchronously
func (client *Client) CreateServiceInstanceWithChan(request *CreateServiceInstanceRequest) (<-chan *CreateServiceInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateServiceInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateServiceInstance(request)
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

// CreateServiceInstanceWithCallback invokes the computenestsupplier.CreateServiceInstance API asynchronously
func (client *Client) CreateServiceInstanceWithCallback(request *CreateServiceInstanceRequest, callback func(response *CreateServiceInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateServiceInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateServiceInstance(request)
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

// CreateServiceInstanceRequest is the request struct for api CreateServiceInstance
type CreateServiceInstanceRequest struct {
	*requests.RpcRequest
	ClientToken       string                      `position:"Query" name:"ClientToken"`
	UserId            string                      `position:"Query" name:"UserId"`
	ResourceGroupId   string                      `position:"Query" name:"ResourceGroupId"`
	TemplateName      string                      `position:"Query" name:"TemplateName"`
	Tag               *[]CreateServiceInstanceTag `position:"Query" name:"Tag"  type:"Repeated"`
	DryRun            requests.Boolean            `position:"Query" name:"DryRun"`
	EndTime           string                      `position:"Query" name:"EndTime"`
	SpecificationName string                      `position:"Query" name:"SpecificationName"`
	Name              string                      `position:"Query" name:"Name"`
	ServiceVersion    string                      `position:"Query" name:"ServiceVersion"`
	ServiceId         string                      `position:"Query" name:"ServiceId"`
	Parameters        string                      `position:"Query" name:"Parameters"`
}

// CreateServiceInstanceTag is a repeated param struct in CreateServiceInstanceRequest
type CreateServiceInstanceTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateServiceInstanceResponse is the response struct for api CreateServiceInstance
type CreateServiceInstanceResponse struct {
	*responses.BaseResponse
	Status            string `json:"Status" xml:"Status"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
	ServiceInstanceId string `json:"ServiceInstanceId" xml:"ServiceInstanceId"`
}

// CreateCreateServiceInstanceRequest creates a request to invoke CreateServiceInstance API
func CreateCreateServiceInstanceRequest() (request *CreateServiceInstanceRequest) {
	request = &CreateServiceInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "CreateServiceInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateServiceInstanceResponse creates a response to parse from CreateServiceInstance response
func CreateCreateServiceInstanceResponse() (response *CreateServiceInstanceResponse) {
	response = &CreateServiceInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
