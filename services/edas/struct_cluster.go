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

// Cluster is a nested struct in edas response
type Cluster struct {
	SubNetCidr          string `json:"SubNetCidr" xml:"SubNetCidr"`
	ClusterId           string `json:"ClusterId" xml:"ClusterId"`
	NodeNum             int    `json:"NodeNum" xml:"NodeNum"`
	CsClusterId         string `json:"CsClusterId" xml:"CsClusterId"`
	UpdateTime          int64  `json:"UpdateTime" xml:"UpdateTime"`
	Mem                 int    `json:"Mem" xml:"Mem"`
	CsClusterStatus     string `json:"CsClusterStatus" xml:"CsClusterStatus"`
	CreateTime          int64  `json:"CreateTime" xml:"CreateTime"`
	OversoldFactor      int    `json:"OversoldFactor" xml:"OversoldFactor"`
	MemUsed             int    `json:"MemUsed" xml:"MemUsed"`
	NetworkMode         int    `json:"NetworkMode" xml:"NetworkMode"`
	ClusterImportStatus int    `json:"ClusterImportStatus" xml:"ClusterImportStatus"`
	ClusterStatus       int    `json:"ClusterStatus" xml:"ClusterStatus"`
	RegionId            string `json:"RegionId" xml:"RegionId"`
	ResourceGroupId     string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	VswitchId           string `json:"VswitchId" xml:"VswitchId"`
	ClusterName         string `json:"ClusterName" xml:"ClusterName"`
	CpuUsed             int    `json:"CpuUsed" xml:"CpuUsed"`
	IaasProvider        string `json:"IaasProvider" xml:"IaasProvider"`
	ClusterType         int    `json:"ClusterType" xml:"ClusterType"`
	VpcId               string `json:"VpcId" xml:"VpcId"`
	Cpu                 int    `json:"Cpu" xml:"Cpu"`
	Description         string `json:"Description" xml:"Description"`
}
