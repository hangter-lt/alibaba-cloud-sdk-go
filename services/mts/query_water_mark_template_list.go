package mts

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

func (client *Client) QueryWaterMarkTemplateList(request *QueryWaterMarkTemplateListRequest) (response *QueryWaterMarkTemplateListResponse, err error) {
	response = CreateQueryWaterMarkTemplateListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryWaterMarkTemplateListWithChan(request *QueryWaterMarkTemplateListRequest) (<-chan *QueryWaterMarkTemplateListResponse, <-chan error) {
	responseChan := make(chan *QueryWaterMarkTemplateListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryWaterMarkTemplateList(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) QueryWaterMarkTemplateListWithCallback(request *QueryWaterMarkTemplateListRequest, callback func(response *QueryWaterMarkTemplateListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryWaterMarkTemplateListResponse
		var err error
		defer close(result)
		response, err = client.QueryWaterMarkTemplateList(request)
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

type QueryWaterMarkTemplateListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	WaterMarkTemplateIds string           `position:"Query" name:"WaterMarkTemplateIds"`
}

type QueryWaterMarkTemplateListResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	NonExistWids struct {
		String []string `json:"String" xml:"String"`
	} `json:"NonExistWids" xml:"NonExistWids"`
	WaterMarkTemplateList struct {
		WaterMarkTemplate []struct {
			Id       string `json:"Id" xml:"Id"`
			Name     string `json:"Name" xml:"Name"`
			Width    string `json:"Width" xml:"Width"`
			Height   string `json:"Height" xml:"Height"`
			Dx       string `json:"Dx" xml:"Dx"`
			Dy       string `json:"Dy" xml:"Dy"`
			ReferPos string `json:"ReferPos" xml:"ReferPos"`
			Type     string `json:"Type" xml:"Type"`
			State    string `json:"State" xml:"State"`
			Timeline struct {
				Start    string `json:"Start" xml:"Start"`
				Duration string `json:"Duration" xml:"Duration"`
			} `json:"Timeline" xml:"Timeline"`
			RatioRefer struct {
				Dx     string `json:"Dx" xml:"Dx"`
				Dy     string `json:"Dy" xml:"Dy"`
				Width  string `json:"Width" xml:"Width"`
				Height string `json:"Height" xml:"Height"`
			} `json:"RatioRefer" xml:"RatioRefer"`
		} `json:"WaterMarkTemplate" xml:"WaterMarkTemplate"`
	} `json:"WaterMarkTemplateList" xml:"WaterMarkTemplateList"`
}

func CreateQueryWaterMarkTemplateListRequest() (request *QueryWaterMarkTemplateListRequest) {
	request = &QueryWaterMarkTemplateListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryWaterMarkTemplateList", "mts", "openAPI")
	return
}

func CreateQueryWaterMarkTemplateListResponse() (response *QueryWaterMarkTemplateListResponse) {
	response = &QueryWaterMarkTemplateListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
