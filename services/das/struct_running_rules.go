package das

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

// RunningRules is a nested struct in das response
type RunningRules struct {
	ItemId                 int64  `json:"ItemId" xml:"ItemId"`
	SqlType                string `json:"SqlType" xml:"SqlType"`
	InstanceId             string `json:"InstanceId" xml:"InstanceId"`
	SqlKeywords            string `json:"SqlKeywords" xml:"SqlKeywords"`
	StartTime              int64  `json:"StartTime" xml:"StartTime"`
	KeywordsHash           string `json:"KeywordsHash" xml:"KeywordsHash"`
	ConcurrencyControlTime int64  `json:"ConcurrencyControlTime" xml:"ConcurrencyControlTime"`
	UserId                 string `json:"UserId" xml:"UserId"`
	MaxConcurrency         string `json:"MaxConcurrency" xml:"MaxConcurrency"`
	Status                 string `json:"Status" xml:"Status"`
}