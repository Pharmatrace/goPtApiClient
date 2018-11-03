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

type Wallet struct {
	// current balance
	Balance float32 `json:"balance,omitempty"`
	// wallet uuid
	WalletId string `json:"wallet_id,omitempty"`
	// descriptive name of the wallet so user can identify it if s/he has multiple wallets
	DisplayName string `json:"display_name,omitempty"`
	// id of the wallet owner (normally an organization
	OwnerStakeholderId string `json:"owner_stakeholder_id,omitempty"`
}
