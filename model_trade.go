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

type Trade struct {
	// GMT, Unix Epoch
	Timestamp int64 `json:"timestamp,omitempty"`
	// unique id hash of the transaction
	TxUid string `json:"tx_uid,omitempty"`
	// id reference to the terms and conditions for this transaction
	TermsId string `json:"terms_id,omitempty"`
	// hyperledger channel the transaction has been executed in
	TransactionChannel string `json:"transaction_channel,omitempty"`
}