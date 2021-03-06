/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import "encoding/json"

type V1alpha1RevisionHistory struct {
	DeployStartedAt *V1Time                    `json:"deployStartedAt,omitempty"`
	DeployedAt      *V1Time                    `json:"deployedAt,omitempty"`
	Id              json.Number                `json:"id,omitempty"`
	Revision        string                     `json:"revision,omitempty"`
	Source          *V1alpha1ApplicationSource `json:"source,omitempty"`
}
