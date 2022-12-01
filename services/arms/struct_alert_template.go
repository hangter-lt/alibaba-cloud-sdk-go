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

// AlertTemplate is a nested struct in arms response
type AlertTemplate struct {
	Annotations              map[string]interface{}   `json:"Annotations" xml:"Annotations"`
	Name                     string                   `json:"Name" xml:"Name"`
	Id                       int                      `json:"Id" xml:"Id"`
	AlertProvider            string                   `json:"AlertProvider" xml:"AlertProvider"`
	Status                   bool                     `json:"Status" xml:"Status"`
	Message                  string                   `json:"Message" xml:"Message"`
	ParentId                 int                      `json:"ParentId" xml:"ParentId"`
	Public                   bool                     `json:"Public" xml:"Public"`
	Labels                   map[string]interface{}   `json:"Labels" xml:"Labels"`
	UserId                   string                   `json:"UserId" xml:"UserId"`
	Rule                     string                   `json:"Rule" xml:"Rule"`
	Type                     string                   `json:"Type" xml:"Type"`
	TemplateProvider         string                   `json:"TemplateProvider" xml:"TemplateProvider"`
	LabelMatchExpressionGrid LabelMatchExpressionGrid `json:"LabelMatchExpressionGrid" xml:"LabelMatchExpressionGrid"`
}
