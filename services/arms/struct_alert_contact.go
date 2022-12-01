package arms

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

// AlertContact is a nested struct in arms response
type AlertContact struct {
	ContactId         float64 `json:"ContactId" xml:"ContactId"`
	ContactName       string  `json:"ContactName" xml:"ContactName"`
	Phone             string  `json:"Phone" xml:"Phone"`
	Email             string  `json:"Email" xml:"Email"`
	IsVerify          bool    `json:"IsVerify" xml:"IsVerify"`
	ReissueSendNotice int64   `json:"ReissueSendNotice" xml:"ReissueSendNotice"`
	IsEmailVerify     bool    `json:"isEmailVerify" xml:"isEmailVerify"`
}
