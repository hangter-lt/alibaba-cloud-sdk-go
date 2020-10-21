package sae

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

// DataItem is a nested struct in sae response
type DataItem struct {
	Id                   int    `json:"Id" xml:"Id"`
	Expired              bool   `json:"Expired" xml:"Expired"`
	EdasContainerVersion string `json:"EdasContainerVersion" xml:"EdasContainerVersion"`
	Memory               int    `json:"Memory" xml:"Memory"`
	Disabled             bool   `json:"Disabled" xml:"Disabled"`
	Version              int    `json:"Version" xml:"Version"`
	Enable               bool   `json:"Enable" xml:"Enable"`
	ComponentKey         string `json:"ComponentKey" xml:"ComponentKey"`
	ComponentDescription string `json:"ComponentDescription" xml:"ComponentDescription"`
	Cpu                  int    `json:"Cpu" xml:"Cpu"`
	SpecInfo             string `json:"SpecInfo" xml:"SpecInfo"`
	Type                 string `json:"Type" xml:"Type"`
}
