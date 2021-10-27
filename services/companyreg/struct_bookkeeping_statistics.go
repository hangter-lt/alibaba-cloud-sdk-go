package companyreg

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

// BookkeepingStatistics is a nested struct in companyreg response
type BookkeepingStatistics struct {
	ProduceBizId  string      `json:"ProduceBizId" xml:"ProduceBizId"`
	Year          int         `json:"Year" xml:"Year"`
	Month         int         `json:"Month" xml:"Month"`
	Income        float64     `json:"Income" xml:"Income"`
	Expend        float64     `json:"Expend" xml:"Expend"`
	Profit        float64     `json:"Profit" xml:"Profit"`
	TaxAmount     float64     `json:"TaxAmount" xml:"TaxAmount"`
	TaxZzs        float64     `json:"TaxZzs" xml:"TaxZzs"`
	TaxFjs        float64     `json:"TaxFjs" xml:"TaxFjs"`
	TaxQysds      float64     `json:"TaxQysds" xml:"TaxQysds"`
	TaxYhs        float64     `json:"TaxYhs" xml:"TaxYhs"`
	TaxGhjf       float64     `json:"TaxGhjf" xml:"TaxGhjf"`
	TaxSljj       float64     `json:"TaxSljj" xml:"TaxSljj"`
	TaxCjrbzj     float64     `json:"TaxCjrbzj" xml:"TaxCjrbzj"`
	TaxOther      float64     `json:"TaxOther" xml:"TaxOther"`
	TaxAmountNote string      `json:"TaxAmountNote" xml:"TaxAmountNote"`
	TaxZzsNote    string      `json:"TaxZzsNote" xml:"TaxZzsNote"`
	TaxFjsNote    string      `json:"TaxFjsNote" xml:"TaxFjsNote"`
	TaxQysdsNote  string      `json:"TaxQysdsNote" xml:"TaxQysdsNote"`
	TaxYhsNote    string      `json:"TaxYhsNote" xml:"TaxYhsNote"`
	TaxGhjfNote   string      `json:"TaxGhjfNote" xml:"TaxGhjfNote"`
	TaxSljjNote   string      `json:"TaxSljjNote" xml:"TaxSljjNote"`
	TaxCjrbzjNote string      `json:"TaxCjrbzjNote" xml:"TaxCjrbzjNote"`
	TaxOtherNote  string      `json:"TaxOtherNote" xml:"TaxOtherNote"`
	VoucherCount  int         `json:"VoucherCount" xml:"VoucherCount"`
	SubjectCount  int         `json:"SubjectCount" xml:"SubjectCount"`
	TaxDetails    []TaxDetail `json:"TaxDetails" xml:"TaxDetails"`
}