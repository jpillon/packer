/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DiskCreate struct for DiskCreate
type DiskCreate struct {
	Service  string            `json:"service"`
	Name     string            `json:"name,omitempty"`
	Size     float32           `json:"size,omitempty"`
	Cloud    string            `json:"cloud,omitempty"`
	Metadata DiskMetadata      `json:"metadata,omitempty"`
	Source   string            `json:"source,omitempty"`
	Tag      map[string]string `json:"tag,omitempty"`
}