package clickhouse

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

// DescribeDBClusterConfigInXML invokes the clickhouse.DescribeDBClusterConfigInXML API synchronously
func (client *Client) DescribeDBClusterConfigInXML(request *DescribeDBClusterConfigInXMLRequest) (response *DescribeDBClusterConfigInXMLResponse, err error) {
	response = CreateDescribeDBClusterConfigInXMLResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterConfigInXMLWithChan invokes the clickhouse.DescribeDBClusterConfigInXML API asynchronously
func (client *Client) DescribeDBClusterConfigInXMLWithChan(request *DescribeDBClusterConfigInXMLRequest) (<-chan *DescribeDBClusterConfigInXMLResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterConfigInXMLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterConfigInXML(request)
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

// DescribeDBClusterConfigInXMLWithCallback invokes the clickhouse.DescribeDBClusterConfigInXML API asynchronously
func (client *Client) DescribeDBClusterConfigInXMLWithCallback(request *DescribeDBClusterConfigInXMLRequest, callback func(response *DescribeDBClusterConfigInXMLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterConfigInXMLResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterConfigInXML(request)
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

// DescribeDBClusterConfigInXMLRequest is the request struct for api DescribeDBClusterConfigInXML
type DescribeDBClusterConfigInXMLRequest struct {
	*requests.RpcRequest
	DBClusterId string `position:"Query" name:"DBClusterId"`
}

// DescribeDBClusterConfigInXMLResponse is the response struct for api DescribeDBClusterConfigInXML
type DescribeDBClusterConfigInXMLResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Config    string `json:"Config" xml:"Config"`
}

// CreateDescribeDBClusterConfigInXMLRequest creates a request to invoke DescribeDBClusterConfigInXML API
func CreateDescribeDBClusterConfigInXMLRequest() (request *DescribeDBClusterConfigInXMLRequest) {
	request = &DescribeDBClusterConfigInXMLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "DescribeDBClusterConfigInXML", "service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterConfigInXMLResponse creates a response to parse from DescribeDBClusterConfigInXML response
func CreateDescribeDBClusterConfigInXMLResponse() (response *DescribeDBClusterConfigInXMLResponse) {
	response = &DescribeDBClusterConfigInXMLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
