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

// QueryBookingDomainInfo invokes the domain.QueryBookingDomainInfo API synchronously
func (client *Client) QueryBookingDomainInfo(request *QueryBookingDomainInfoRequest) (response *QueryBookingDomainInfoResponse, err error) {
	response = CreateQueryBookingDomainInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBookingDomainInfoWithChan invokes the domain.QueryBookingDomainInfo API asynchronously
func (client *Client) QueryBookingDomainInfoWithChan(request *QueryBookingDomainInfoRequest) (<-chan *QueryBookingDomainInfoResponse, <-chan error) {
	responseChan := make(chan *QueryBookingDomainInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryBookingDomainInfo(request)
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

// QueryBookingDomainInfoWithCallback invokes the domain.QueryBookingDomainInfo API asynchronously
func (client *Client) QueryBookingDomainInfoWithCallback(request *QueryBookingDomainInfoRequest, callback func(response *QueryBookingDomainInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryBookingDomainInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryBookingDomainInfo(request)
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

// QueryBookingDomainInfoRequest is the request struct for api QueryBookingDomainInfo
type QueryBookingDomainInfoRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Body" name:"DomainName"`
}

// QueryBookingDomainInfoResponse is the response struct for api QueryBookingDomainInfo
type QueryBookingDomainInfoResponse struct {
	*responses.BaseResponse
	BookEndTime     int64   `json:"BookEndTime" xml:"BookEndTime"`
	RequestId       string  `json:"RequestId" xml:"RequestId"`
	MaxBid          float64 `json:"MaxBid" xml:"MaxBid"`
	TransferInPrice float64 `json:"TransferInPrice" xml:"TransferInPrice"`
	AuctionId       int     `json:"AuctionId" xml:"AuctionId"`
	Currency        string  `json:"Currency" xml:"Currency"`
	PartnerType     string  `json:"PartnerType" xml:"PartnerType"`
	SnatchNo        string  `json:"SnatchNo" xml:"SnatchNo"`
}

// CreateQueryBookingDomainInfoRequest creates a request to invoke QueryBookingDomainInfo API
func CreateQueryBookingDomainInfoRequest() (request *QueryBookingDomainInfoRequest) {
	request = &QueryBookingDomainInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "QueryBookingDomainInfo", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryBookingDomainInfoResponse creates a response to parse from QueryBookingDomainInfo response
func CreateQueryBookingDomainInfoResponse() (response *QueryBookingDomainInfoResponse) {
	response = &QueryBookingDomainInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
