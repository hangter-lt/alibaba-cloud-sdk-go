package dms_enterprise

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

// ConfigDetail is a nested struct in dms_enterprise response
type ConfigDetail struct {
	DetailType            string          `json:"DetailType" xml:"DetailType"`
	FileType              string          `json:"FileType" xml:"FileType"`
	CsvTableName          string          `json:"CsvTableName" xml:"CsvTableName"`
	FileEncoding          string          `json:"FileEncoding" xml:"FileEncoding"`
	Cron                  bool            `json:"Cron" xml:"Cron"`
	CronCallTimes         int             `json:"CronCallTimes" xml:"CronCallTimes"`
	CronFormat            string          `json:"CronFormat" xml:"CronFormat"`
	Duration              int             `json:"Duration" xml:"Duration"`
	CronStatus            string          `json:"CronStatus" xml:"CronStatus"`
	CronLastCallStartTime string          `json:"CronLastCallStartTime" xml:"CronLastCallStartTime"`
	CronNextCallTime      string          `json:"CronNextCallTime" xml:"CronNextCallTime"`
	CurrentTaskId         int64           `json:"CurrentTaskId" xml:"CurrentTaskId"`
	ImportExtConfig       ImportExtConfig `json:"ImportExtConfig" xml:"ImportExtConfig"`
	CronExtConfig         CronExtConfig   `json:"CronExtConfig" xml:"CronExtConfig"`
}
