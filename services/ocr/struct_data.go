package ocr

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

// Data is a nested struct in ocr response
type Data struct {
	ExpiryDate          string               `json:"ExpiryDate" xml:"ExpiryDate"`
	Date                string               `json:"Date" xml:"Date"`
	ExpiryDay           string               `json:"ExpiryDay" xml:"ExpiryDay"`
	Seat                string               `json:"Seat" xml:"Seat"`
	EstablishDate       string               `json:"EstablishDate" xml:"EstablishDate"`
	FileContent         string               `json:"FileContent" xml:"FileContent"`
	ValidPeriod         string               `json:"ValidPeriod" xml:"ValidPeriod"`
	Type                string               `json:"Type" xml:"Type"`
	NameChineseRaw      string               `json:"NameChineseRaw" xml:"NameChineseRaw"`
	NativePlace         string               `json:"NativePlace" xml:"NativePlace"`
	LegalPerson         string               `json:"LegalPerson" xml:"LegalPerson"`
	RegisterNumber      string               `json:"RegisterNumber" xml:"RegisterNumber"`
	IssuePlaceRaw       string               `json:"IssuePlaceRaw" xml:"IssuePlaceRaw"`
	Price               float64              `json:"Price" xml:"Price"`
	Destination         string               `json:"Destination" xml:"Destination"`
	Relation            string               `json:"Relation" xml:"Relation"`
	Height              int64                `json:"Height" xml:"Height"`
	BirthPlaceRaw       string               `json:"BirthPlaceRaw" xml:"BirthPlaceRaw"`
	DepartureStation    string               `json:"DepartureStation" xml:"DepartureStation"`
	Gender              string               `json:"Gender" xml:"Gender"`
	BirthDay            string               `json:"BirthDay" xml:"BirthDay"`
	Success             bool                 `json:"Success" xml:"Success"`
	PassportNo          string               `json:"PassportNo" xml:"PassportNo"`
	IssueDate           string               `json:"IssueDate" xml:"IssueDate"`
	Status              string               `json:"Status" xml:"Status"`
	Name                string               `json:"Name" xml:"Name"`
	BankName            string               `json:"BankName" xml:"BankName"`
	IDNumber            string               `json:"IDNumber" xml:"IDNumber"`
	ErrorMessage        string               `json:"ErrorMessage" xml:"ErrorMessage"`
	Sex                 string               `json:"Sex" xml:"Sex"`
	Capital             string               `json:"Capital" xml:"Capital"`
	LineOne             string               `json:"LineOne" xml:"LineOne"`
	IsBlur              bool                 `json:"IsBlur" xml:"IsBlur"`
	Authority           string               `json:"Authority" xml:"Authority"`
	CardNumber          string               `json:"CardNumber" xml:"CardNumber"`
	VinCode             string               `json:"VinCode" xml:"VinCode"`
	ErrorCode           string               `json:"ErrorCode" xml:"ErrorCode"`
	Angle               float64              `json:"Angle" xml:"Angle"`
	Business            string               `json:"Business" xml:"Business"`
	Country             string               `json:"Country" xml:"Country"`
	JobId               string               `json:"JobId" xml:"JobId"`
	BirthPlace          string               `json:"BirthPlace" xml:"BirthPlace"`
	Result              string               `json:"Result" xml:"Result"`
	Nationality         string               `json:"Nationality" xml:"Nationality"`
	ValidDate           string               `json:"ValidDate" xml:"ValidDate"`
	Level               string               `json:"Level" xml:"Level"`
	PersonId            string               `json:"PersonId" xml:"PersonId"`
	SourceCountry       string               `json:"SourceCountry" xml:"SourceCountry"`
	NameChinese         string               `json:"NameChinese" xml:"NameChinese"`
	Address             string               `json:"Address" xml:"Address"`
	IsCard              bool                 `json:"IsCard" xml:"IsCard"`
	Number              string               `json:"Number" xml:"Number"`
	InputFile           string               `json:"InputFile" xml:"InputFile"`
	Width               int64                `json:"Width" xml:"Width"`
	BirthDate           string               `json:"BirthDate" xml:"BirthDate"`
	LineZero            string               `json:"LineZero" xml:"LineZero"`
	IssuePlace          string               `json:"IssuePlace" xml:"IssuePlace"`
	Departments         []string             `json:"Departments" xml:"Departments"`
	Companies           []string             `json:"Companies" xml:"Companies"`
	Emails              []string             `json:"Emails" xml:"Emails"`
	CellPhoneNumbers    []string             `json:"CellPhoneNumbers" xml:"CellPhoneNumbers"`
	OfficePhoneNumbers  []string             `json:"OfficePhoneNumbers" xml:"OfficePhoneNumbers"`
	Titles              []string             `json:"Titles" xml:"Titles"`
	Addresses           []string             `json:"Addresses" xml:"Addresses"`
	BackResult          BackResult           `json:"BackResult" xml:"BackResult"`
	SpoofResult         SpoofResult          `json:"SpoofResult" xml:"SpoofResult"`
	FaceResult          FaceResult           `json:"FaceResult" xml:"FaceResult"`
	Summary             Summary              `json:"Summary" xml:"Summary"`
	Box                 Box                  `json:"Box" xml:"Box"`
	Content             Content              `json:"Content" xml:"Content"`
	TitleArea           TitleArea            `json:"TitleArea" xml:"TitleArea"`
	Stamp               Stamp                `json:"Stamp" xml:"Stamp"`
	Emblem              Emblem               `json:"Emblem" xml:"Emblem"`
	FrontResult         FrontResult          `json:"FrontResult" xml:"FrontResult"`
	Title               Title                `json:"Title" xml:"Title"`
	QRCode              QRCode               `json:"QRCode" xml:"QRCode"`
	Plates              []Plate              `json:"Plates" xml:"Plates"`
	Elements            []Element            `json:"Elements" xml:"Elements"`
	Frames              []Frame              `json:"Frames" xml:"Frames"`
	InvalidStampAreas   []InvalidStampArea   `json:"InvalidStampAreas" xml:"InvalidStampAreas"`
	Regions             []Region             `json:"Regions" xml:"Regions"`
	Invoices            []Invoice            `json:"Invoices" xml:"Invoices"`
	Results             []Result             `json:"Results" xml:"Results"`
	UndertakeStampAreas []UndertakeStampArea `json:"UndertakeStampAreas" xml:"UndertakeStampAreas"`
	RegisterStampAreas  []RegisterStampArea  `json:"RegisterStampAreas" xml:"RegisterStampAreas"`
	OtherStampAreas     []OtherStampArea     `json:"OtherStampAreas" xml:"OtherStampAreas"`
	Signboards          []SignboardsItem     `json:"Signboards" xml:"Signboards"`
	Tables              []Table              `json:"Tables" xml:"Tables"`
}
