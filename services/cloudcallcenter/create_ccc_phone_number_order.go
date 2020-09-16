package cloudcallcenter

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

// CreateCCCPhoneNumberOrder invokes the cloudcallcenter.CreateCCCPhoneNumberOrder API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createcccphonenumberorder.html
func (client *Client) CreateCCCPhoneNumberOrder(request *CreateCCCPhoneNumberOrderRequest) (response *CreateCCCPhoneNumberOrderResponse, err error) {
	response = CreateCreateCCCPhoneNumberOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCCCPhoneNumberOrderWithChan invokes the cloudcallcenter.CreateCCCPhoneNumberOrder API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createcccphonenumberorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCCCPhoneNumberOrderWithChan(request *CreateCCCPhoneNumberOrderRequest) (<-chan *CreateCCCPhoneNumberOrderResponse, <-chan error) {
	responseChan := make(chan *CreateCCCPhoneNumberOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCCCPhoneNumberOrder(request)
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

// CreateCCCPhoneNumberOrderWithCallback invokes the cloudcallcenter.CreateCCCPhoneNumberOrder API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createcccphonenumberorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCCCPhoneNumberOrderWithCallback(request *CreateCCCPhoneNumberOrderRequest, callback func(response *CreateCCCPhoneNumberOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCCCPhoneNumberOrderResponse
		var err error
		defer close(result)
		response, err = client.CreateCCCPhoneNumberOrder(request)
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

// CreateCCCPhoneNumberOrderRequest is the request struct for api CreateCCCPhoneNumberOrder
type CreateCCCPhoneNumberOrderRequest struct {
	*requests.RpcRequest
	SpecId             string           `position:"Query" name:"SpecId"`
	SpecName           string           `position:"Query" name:"SpecName"`
	RealNameInsId      string           `position:"Query" name:"RealNameInsId"`
	RegionNameProvince string           `position:"Query" name:"RegionNameProvince"`
	MonthlyPrice       requests.Integer `position:"Query" name:"MonthlyPrice"`
	Number             *[]string        `position:"Query" name:"Number"  type:"Repeated"`
	CorpName           string           `position:"Query" name:"CorpName"`
	RegionNameCity     string           `position:"Query" name:"RegionNameCity"`
}

// CreateCCCPhoneNumberOrderResponse is the response struct for api CreateCCCPhoneNumberOrder
type CreateCCCPhoneNumberOrderResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	OrderPayUrl    string `json:"OrderPayUrl" xml:"OrderPayUrl"`
}

// CreateCreateCCCPhoneNumberOrderRequest creates a request to invoke CreateCCCPhoneNumberOrder API
func CreateCreateCCCPhoneNumberOrderRequest() (request *CreateCCCPhoneNumberOrderRequest) {
	request = &CreateCCCPhoneNumberOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateCCCPhoneNumberOrder", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateCCCPhoneNumberOrderResponse creates a response to parse from CreateCCCPhoneNumberOrder response
func CreateCreateCCCPhoneNumberOrderResponse() (response *CreateCCCPhoneNumberOrderResponse) {
	response = &CreateCCCPhoneNumberOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}