package polardbx

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

// LogicInstanceTopology is a nested struct in polardbx response
type LogicInstanceTopology struct {
	DBInstanceName              string                                  `json:"DBInstanceName" xml:"DBInstanceName"`
	DBInstanceCreateTime        string                                  `json:"DBInstanceCreateTime" xml:"DBInstanceCreateTime"`
	MaintainStartTime           string                                  `json:"MaintainStartTime" xml:"MaintainStartTime"`
	MaintainEndTime             string                                  `json:"MaintainEndTime" xml:"MaintainEndTime"`
	LockReason                  string                                  `json:"LockReason" xml:"LockReason"`
	DBInstanceStatus            int                                     `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	LockMode                    int                                     `json:"LockMode" xml:"LockMode"`
	EngineVersion               string                                  `json:"EngineVersion" xml:"EngineVersion"`
	DBInstanceStorage           int                                     `json:"DBInstanceStorage" xml:"DBInstanceStorage"`
	DBInstanceConnType          string                                  `json:"DBInstanceConnType" xml:"DBInstanceConnType"`
	DBInstanceId                string                                  `json:"DBInstanceId" xml:"DBInstanceId"`
	Engine                      string                                  `json:"Engine" xml:"Engine"`
	DBInstanceDescription       string                                  `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
	DBInstanceStatusDescription string                                  `json:"DBInstanceStatusDescription" xml:"DBInstanceStatusDescription"`
	Items                       []ItemsItemInDescribeDBInstanceTopology `json:"Items" xml:"Items"`
	HistoryItems                []HistoryItemsItem                      `json:"HistoryItems" xml:"HistoryItems"`
}
