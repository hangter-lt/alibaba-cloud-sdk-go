package ddoscoo

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

// ElasticQpsItem is a nested struct in ddoscoo response
type ElasticQpsItem struct {
	MaxNormalQps int64 `json:"MaxNormalQps" xml:"MaxNormalQps"`
	Index        int64 `json:"Index" xml:"Index"`
	MaxQps       int64 `json:"MaxQps" xml:"MaxQps"`
	Pv           int64 `json:"Pv" xml:"Pv"`
	Ups          int64 `json:"Ups" xml:"Ups"`
	Status2      int64 `json:"Status2" xml:"Status2"`
	Status3      int64 `json:"Status3" xml:"Status3"`
	Status4      int64 `json:"Status4" xml:"Status4"`
	Status5      int64 `json:"Status5" xml:"Status5"`
}
