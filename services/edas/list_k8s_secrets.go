package edas

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

// ListK8sSecrets invokes the edas.ListK8sSecrets API synchronously
func (client *Client) ListK8sSecrets(request *ListK8sSecretsRequest) (response *ListK8sSecretsResponse, err error) {
	response = CreateListK8sSecretsResponse()
	err = client.DoAction(request, response)
	return
}

// ListK8sSecretsWithChan invokes the edas.ListK8sSecrets API asynchronously
func (client *Client) ListK8sSecretsWithChan(request *ListK8sSecretsRequest) (<-chan *ListK8sSecretsResponse, <-chan error) {
	responseChan := make(chan *ListK8sSecretsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListK8sSecrets(request)
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

// ListK8sSecretsWithCallback invokes the edas.ListK8sSecrets API asynchronously
func (client *Client) ListK8sSecretsWithCallback(request *ListK8sSecretsRequest, callback func(response *ListK8sSecretsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListK8sSecretsResponse
		var err error
		defer close(result)
		response, err = client.ListK8sSecrets(request)
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

// ListK8sSecretsRequest is the request struct for api ListK8sSecrets
type ListK8sSecretsRequest struct {
	*requests.RoaRequest
	Condition       string           `position:"Query" name:"Condition"`
	PageNo          requests.Integer `position:"Query" name:"PageNo"`
	Namespace       string           `position:"Query" name:"Namespace"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	ShowRelatedApps requests.Boolean `position:"Query" name:"ShowRelatedApps"`
}

// ListK8sSecretsResponse is the response struct for api ListK8sSecrets
type ListK8sSecretsResponse struct {
	*responses.BaseResponse
	Code       int              `json:"Code" xml:"Code"`
	Message    string           `json:"Message" xml:"Message"`
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	K8sSecrets []K8sSecretsItem `json:"K8sSecrets" xml:"K8sSecrets"`
	Result     []ResultItem     `json:"Result" xml:"Result"`
}

// CreateListK8sSecretsRequest creates a request to invoke ListK8sSecrets API
func CreateListK8sSecretsRequest() (request *ListK8sSecretsRequest) {
	request = &ListK8sSecretsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListK8sSecrets", "/pop/v5/k8s/acs/k8s_secret", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListK8sSecretsResponse creates a response to parse from ListK8sSecrets response
func CreateListK8sSecretsResponse() (response *ListK8sSecretsResponse) {
	response = &ListK8sSecretsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}