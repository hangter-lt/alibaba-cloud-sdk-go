package domain

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

// ReserveDomain invokes the domain.ReserveDomain API synchronously
func (client *Client) ReserveDomain(request *ReserveDomainRequest) (response *ReserveDomainResponse, err error) {
	response = CreateReserveDomainResponse()
	err = client.DoAction(request, response)
	return
}

// ReserveDomainWithChan invokes the domain.ReserveDomain API asynchronously
func (client *Client) ReserveDomainWithChan(request *ReserveDomainRequest) (<-chan *ReserveDomainResponse, <-chan error) {
	responseChan := make(chan *ReserveDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReserveDomain(request)
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

// ReserveDomainWithCallback invokes the domain.ReserveDomain API asynchronously
func (client *Client) ReserveDomainWithCallback(request *ReserveDomainRequest, callback func(response *ReserveDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReserveDomainResponse
		var err error
		defer close(result)
		response, err = client.ReserveDomain(request)
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

// ReserveDomainRequest is the request struct for api ReserveDomain
type ReserveDomainRequest struct {
	*requests.RpcRequest
	DomainName string    `position:"Body" name:"DomainName"`
	Channels   *[]string `position:"Body" name:"Channels"  type:"Repeated"`
}

// ReserveDomainResponse is the response struct for api ReserveDomain
type ReserveDomainResponse struct {
	*responses.BaseResponse
	AuctionId string `json:"AuctionId" xml:"AuctionId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReserveDomainRequest creates a request to invoke ReserveDomain API
func CreateReserveDomainRequest() (request *ReserveDomainRequest) {
	request = &ReserveDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "ReserveDomain", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReserveDomainResponse creates a response to parse from ReserveDomain response
func CreateReserveDomainResponse() (response *ReserveDomainResponse) {
	response = &ReserveDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
