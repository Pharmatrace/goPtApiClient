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

type Location struct {
	// GMT, C time format.
	Timestamp int64 `json:"timestamp,omitempty"`
	// Unique identifier for the location.
	Uuid string `json:"uuid,omitempty"`
	// longitude.
	Longitude float32 `json:"longitude,omitempty"`
	// latitude.
	Latitude float32 `json:"latitude,omitempty"`
	// type of container location, e.g. VESSEL, AIR, STREET, WAREHOUSE, PROCESSING, STORE, WITH_PATIENT, ...
	LocationType string `json:"location_type,omitempty"`
}
