/*
 * Nbsf_Management
 *
 * Binding Support Management Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PcfBinding struct {

	Supi string `json:"supi,omitempty"`

	Gpsi string `json:"gpsi,omitempty"`

	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Prefix string `json:"ipv6Prefix,omitempty"`

	IpDomain string `json:"ipDomain,omitempty"`

	MacAddr48 string `json:"macAddr48,omitempty"`

	Dnn string `json:"dnn"`

	PcfFqdn string `json:"pcfFqdn,omitempty"`

	// IP end points of the PCF or the IP end points of the PCF hosting the Npcf_PolicyAuthorization service. At least one of pcfFqdn or pcfIpEndPoints shall be included if the PCF supports N5 interface. If the pcfIpEndPoints is provided at the PCF level, the transport and port in the pcfIpEndPoints are not required.
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty"`

	PcfDiamHost string `json:"pcfDiamHost,omitempty"`

	PcfDiamRealm string `json:"pcfDiamRealm,omitempty"`

	Snssai *Snssai `json:"snssai"`
}
