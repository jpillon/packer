/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SnapshotCreate struct for SnapshotCreate
type SnapshotCreate struct {
	Vault   string            `json:"vault"`
	Service string            `json:"service,omitempty"`
	Name    string            `json:"name"`
	Tag     map[string]string `json:"tag,omitempty"`
}