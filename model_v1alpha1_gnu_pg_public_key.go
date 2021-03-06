/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1GnuPgPublicKey struct {
	Fingerprint string `json:"fingerprint,omitempty"`
	KeyData string `json:"keyData,omitempty"`
	KeyID string `json:"keyID,omitempty"`
	Owner string `json:"owner,omitempty"`
	SubType string `json:"subType,omitempty"`
	Trust string `json:"trust,omitempty"`
}
