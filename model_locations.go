/*
 * PharmaTrace Supply Chain Information Network API
 *
 * The PharmaTrace SCIN API provides network members a resource and process oriented access to the blockchain backed market and inventory information. It represents a layer of abstraction above the Hyperledger network to facilitate a business focused development of clients and integration systems without the need to directly connect to Hyperledger nodes.
 *
 * API version: 0.0.1
 * Contact: api@pharmatrace.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Locations struct {
	// Position in pagination.
	Offset int32 `json:"offset,omitempty"`
	// Number of items to retrieve (100 max).
	Limit int32 `json:"limit,omitempty"`
	// Total number of items available.
	Count int32 `json:"count,omitempty"`
	History []Location `json:"history,omitempty"`
}
