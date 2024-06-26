package live

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

// VideoLayer is a nested struct in live response
type VideoLayer struct {
	FillMode            string                                     `json:"FillMode" xml:"FillMode"`
	FixedDelayDuration  int                                        `json:"FixedDelayDuration" xml:"FixedDelayDuration"`
	HeightNormalized    float64                                    `json:"HeightNormalized" xml:"HeightNormalized"`
	PositionRefer       string                                     `json:"PositionRefer" xml:"PositionRefer"`
	WidthNormalized     float64                                    `json:"WidthNormalized" xml:"WidthNormalized"`
	PositionNormalizeds PositionNormalizedsInDescribeCasterLayouts `json:"PositionNormalizeds" xml:"PositionNormalizeds"`
}
