/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Event is a report of an event somewhere in the cluster.  Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type V1Event struct {
	Action string `json:"action,omitempty"`
	Count int32 `json:"count,omitempty"`
	EventTime *V1MicroTime `json:"eventTime,omitempty"`
	FirstTimestamp *V1Time `json:"firstTimestamp,omitempty"`
	InvolvedObject *V1ObjectReference `json:"involvedObject,omitempty"`
	LastTimestamp *V1Time `json:"lastTimestamp,omitempty"`
	Message string `json:"message,omitempty"`
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Reason string `json:"reason,omitempty"`
	Related *V1ObjectReference `json:"related,omitempty"`
	ReportingComponent string `json:"reportingComponent,omitempty"`
	ReportingInstance string `json:"reportingInstance,omitempty"`
	Series *V1EventSeries `json:"series,omitempty"`
	Source *V1EventSource `json:"source,omitempty"`
	Type_ string `json:"type,omitempty"`
}
