package dms_dg

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

// Gateway is a nested struct in dms_dg response
type Gateway struct {
	RegionId               string `json:"RegionId" xml:"RegionId"`
	ExceptionMsg           string `json:"ExceptionMsg" xml:"ExceptionMsg"`
	DgVersion              string `json:"DgVersion" xml:"DgVersion"`
	UserId                 string `json:"UserId" xml:"UserId"`
	GatewayName            string `json:"GatewayName" xml:"GatewayName"`
	CreatorId              string `json:"CreatorId" xml:"CreatorId"`
	NumOfExceptionInstance int    `json:"NumOfExceptionInstance" xml:"NumOfExceptionInstance"`
	GatewayId              string `json:"GatewayId" xml:"GatewayId"`
	NumOfRunningInstance   int    `json:"NumOfRunningInstance" xml:"NumOfRunningInstance"`
	Status                 string `json:"Status" xml:"Status"`
	GatewayDesc            string `json:"GatewayDesc" xml:"GatewayDesc"`
}
