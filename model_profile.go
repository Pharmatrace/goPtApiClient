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

type Profile struct {
	// First name of the PT network user.
	FirstName string `json:"first_name,omitempty"`
	// Last name of the PT network user.
	LastName string `json:"last_name,omitempty"`
	// Email address
	Email string `json:"email,omitempty"`
	// Twitter name
	TwitterName string `json:"twitter_name,omitempty"`
	// LinkedIn name
	LinkedinName string `json:"linkedin_name,omitempty"`
	// Facebook name
	FacebookName string `json:"facebook_name,omitempty"`
	// Image URL
	Picture string `json:"picture,omitempty"`
	// the PT network member that the user/profile is assigned to.
	StakeholderId string `json:"stakeholder_id,omitempty"`
	// true if the user is currently active. Inactive users will be anonymized but not deleted to maintain integrity of transaction history
	Active bool `json:"active,omitempty"`
}
