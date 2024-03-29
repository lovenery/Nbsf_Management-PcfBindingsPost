/*
 * Nbsf_Management
 *
 * Binding Support Management Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IpEndPoint struct {

	Ipv4Address string `json:"ipv4Address,omitempty"`

	Ipv6Address string `json:"ipv6Address,omitempty"`

	Transport *TransportProtocol `json:"transport,omitempty"`

	Port int32 `json:"port,omitempty"`
}
