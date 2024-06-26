package slb

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

// HTTPSListenerConfig is a nested struct in slb response
type HTTPSListenerConfig struct {
	XForwardedForClientCertClientVerify      string              `json:"XForwardedFor_ClientCertClientVerify" xml:"XForwardedFor_ClientCertClientVerify"`
	HealthCheckHttpVersion                   string              `json:"HealthCheckHttpVersion" xml:"HealthCheckHttpVersion"`
	XForwardedForClientSrcPort               string              `json:"XForwardedFor_ClientSrcPort" xml:"XForwardedFor_ClientSrcPort"`
	Cookie                                   string              `json:"Cookie" xml:"Cookie"`
	Gzip                                     string              `json:"Gzip" xml:"Gzip"`
	EnableHttp2                              string              `json:"EnableHttp2" xml:"EnableHttp2"`
	CACertificateId                          string              `json:"CACertificateId" xml:"CACertificateId"`
	XForwardedForClientCertClientVerifyAlias string              `json:"XForwardedFor_ClientCertClientVerifyAlias" xml:"XForwardedFor_ClientCertClientVerifyAlias"`
	HealthCheckConnectPort                   int                 `json:"HealthCheckConnectPort" xml:"HealthCheckConnectPort"`
	HealthCheckTimeout                       int                 `json:"HealthCheckTimeout" xml:"HealthCheckTimeout"`
	HealthCheckType                          string              `json:"HealthCheckType" xml:"HealthCheckType"`
	CookieTimeout                            int                 `json:"CookieTimeout" xml:"CookieTimeout"`
	HealthCheckDomain                        string              `json:"HealthCheckDomain" xml:"HealthCheckDomain"`
	UnhealthyThreshold                       int                 `json:"UnhealthyThreshold" xml:"UnhealthyThreshold"`
	XForwardedForSLBID                       string              `json:"XForwardedFor_SLBID" xml:"XForwardedFor_SLBID"`
	XForwardedForClientCertSubjectDN         string              `json:"XForwardedFor_ClientCertSubjectDN" xml:"XForwardedFor_ClientCertSubjectDN"`
	HealthCheckHttpCode                      string              `json:"HealthCheckHttpCode" xml:"HealthCheckHttpCode"`
	XForwardedForClientCertFingerprintAlias  string              `json:"XForwardedFor_ClientCertFingerprintAlias" xml:"XForwardedFor_ClientCertFingerprintAlias"`
	XForwardedForClientCertSubjectDNAlias    string              `json:"XForwardedFor_ClientCertSubjectDNAlias" xml:"XForwardedFor_ClientCertSubjectDNAlias"`
	XForwardedForClientCertIssuerDNAlias     string              `json:"XForwardedFor_ClientCertIssuerDNAlias" xml:"XForwardedFor_ClientCertIssuerDNAlias"`
	XForwardedForClientCertFingerprint       string              `json:"XForwardedFor_ClientCertFingerprint" xml:"XForwardedFor_ClientCertFingerprint"`
	XForwardedFor                            string              `json:"XForwardedFor" xml:"XForwardedFor"`
	RequestTimeout                           int                 `json:"RequestTimeout" xml:"RequestTimeout"`
	IdleTimeout                              int                 `json:"IdleTimeout" xml:"IdleTimeout"`
	ServerCertificateId                      string              `json:"ServerCertificateId" xml:"ServerCertificateId"`
	HealthCheckInterval                      int                 `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	XForwardedForSLBPORT                     string              `json:"XForwardedFor_SLBPORT" xml:"XForwardedFor_SLBPORT"`
	HealthCheckURI                           string              `json:"HealthCheckURI" xml:"HealthCheckURI"`
	StickySessionType                        string              `json:"StickySessionType" xml:"StickySessionType"`
	XForwardedForClientCertIssuerDN          string              `json:"XForwardedFor_ClientCertIssuerDN" xml:"XForwardedFor_ClientCertIssuerDN"`
	HealthyThreshold                         int                 `json:"HealthyThreshold" xml:"HealthyThreshold"`
	XForwardedForProto                       string              `json:"XForwardedFor_proto" xml:"XForwardedFor_proto"`
	XForwardedForSLBIP                       string              `json:"XForwardedFor_SLBIP" xml:"XForwardedFor_SLBIP"`
	StickySession                            string              `json:"StickySession" xml:"StickySession"`
	HealthCheckMethod                        string              `json:"HealthCheckMethod" xml:"HealthCheckMethod"`
	TLSCipherPolicy                          string              `json:"TLSCipherPolicy" xml:"TLSCipherPolicy"`
	HealthCheck                              string              `json:"HealthCheck" xml:"HealthCheck"`
	MaxConnection                            int                 `json:"MaxConnection" xml:"MaxConnection"`
	ServerCertificates                       []ServerCertificate `json:"ServerCertificates" xml:"ServerCertificates"`
}
