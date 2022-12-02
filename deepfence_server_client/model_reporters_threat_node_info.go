/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ReportersThreatNodeInfo struct for ReportersThreatNodeInfo
type ReportersThreatNodeInfo struct {
	AttackPath [][]string `json:"attack_path,omitempty"`
	ComplianceCount *int32 `json:"compliance_count,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	NodeType *string `json:"node_type,omitempty"`
	Nodes map[string]ReportersNodeInfo `json:"nodes,omitempty"`
	SecretsCount *int32 `json:"secrets_count,omitempty"`
	VulnerabilityCount *int32 `json:"vulnerability_count,omitempty"`
}

// NewReportersThreatNodeInfo instantiates a new ReportersThreatNodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersThreatNodeInfo() *ReportersThreatNodeInfo {
	this := ReportersThreatNodeInfo{}
	return &this
}

// NewReportersThreatNodeInfoWithDefaults instantiates a new ReportersThreatNodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersThreatNodeInfoWithDefaults() *ReportersThreatNodeInfo {
	this := ReportersThreatNodeInfo{}
	return &this
}

// GetAttackPath returns the AttackPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportersThreatNodeInfo) GetAttackPath() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}
	return o.AttackPath
}

// GetAttackPathOk returns a tuple with the AttackPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersThreatNodeInfo) GetAttackPathOk() ([][]string, bool) {
	if o == nil || isNil(o.AttackPath) {
    return nil, false
	}
	return o.AttackPath, true
}

// HasAttackPath returns a boolean if a field has been set.
func (o *ReportersThreatNodeInfo) HasAttackPath() bool {
	if o != nil && isNil(o.AttackPath) {
		return true
	}

	return false
}

// SetAttackPath gets a reference to the given [][]string and assigns it to the AttackPath field.
func (o *ReportersThreatNodeInfo) SetAttackPath(v [][]string) {
	o.AttackPath = v
}

// GetComplianceCount returns the ComplianceCount field value if set, zero value otherwise.
func (o *ReportersThreatNodeInfo) GetComplianceCount() int32 {
	if o == nil || isNil(o.ComplianceCount) {
		var ret int32
		return ret
	}
	return *o.ComplianceCount
}

// GetComplianceCountOk returns a tuple with the ComplianceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersThreatNodeInfo) GetComplianceCountOk() (*int32, bool) {
	if o == nil || isNil(o.ComplianceCount) {
    return nil, false
	}
	return o.ComplianceCount, true
}

// HasComplianceCount returns a boolean if a field has been set.
func (o *ReportersThreatNodeInfo) HasComplianceCount() bool {
	if o != nil && !isNil(o.ComplianceCount) {
		return true
	}

	return false
}

// SetComplianceCount gets a reference to the given int32 and assigns it to the ComplianceCount field.
func (o *ReportersThreatNodeInfo) SetComplianceCount(v int32) {
	o.ComplianceCount = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReportersThreatNodeInfo) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersThreatNodeInfo) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReportersThreatNodeInfo) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ReportersThreatNodeInfo) SetCount(v int32) {
	o.Count = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportersThreatNodeInfo) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersThreatNodeInfo) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportersThreatNodeInfo) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReportersThreatNodeInfo) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ReportersThreatNodeInfo) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersThreatNodeInfo) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ReportersThreatNodeInfo) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ReportersThreatNodeInfo) SetLabel(v string) {
	o.Label = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *ReportersThreatNodeInfo) GetNodeType() string {
	if o == nil || isNil(o.NodeType) {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersThreatNodeInfo) GetNodeTypeOk() (*string, bool) {
	if o == nil || isNil(o.NodeType) {
    return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *ReportersThreatNodeInfo) HasNodeType() bool {
	if o != nil && !isNil(o.NodeType) {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *ReportersThreatNodeInfo) SetNodeType(v string) {
	o.NodeType = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportersThreatNodeInfo) GetNodes() map[string]ReportersNodeInfo {
	if o == nil {
		var ret map[string]ReportersNodeInfo
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersThreatNodeInfo) GetNodesOk() (*map[string]ReportersNodeInfo, bool) {
	if o == nil || isNil(o.Nodes) {
    return nil, false
	}
	return &o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *ReportersThreatNodeInfo) HasNodes() bool {
	if o != nil && isNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given map[string]ReportersNodeInfo and assigns it to the Nodes field.
func (o *ReportersThreatNodeInfo) SetNodes(v map[string]ReportersNodeInfo) {
	o.Nodes = v
}

// GetSecretsCount returns the SecretsCount field value if set, zero value otherwise.
func (o *ReportersThreatNodeInfo) GetSecretsCount() int32 {
	if o == nil || isNil(o.SecretsCount) {
		var ret int32
		return ret
	}
	return *o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersThreatNodeInfo) GetSecretsCountOk() (*int32, bool) {
	if o == nil || isNil(o.SecretsCount) {
    return nil, false
	}
	return o.SecretsCount, true
}

// HasSecretsCount returns a boolean if a field has been set.
func (o *ReportersThreatNodeInfo) HasSecretsCount() bool {
	if o != nil && !isNil(o.SecretsCount) {
		return true
	}

	return false
}

// SetSecretsCount gets a reference to the given int32 and assigns it to the SecretsCount field.
func (o *ReportersThreatNodeInfo) SetSecretsCount(v int32) {
	o.SecretsCount = &v
}

// GetVulnerabilityCount returns the VulnerabilityCount field value if set, zero value otherwise.
func (o *ReportersThreatNodeInfo) GetVulnerabilityCount() int32 {
	if o == nil || isNil(o.VulnerabilityCount) {
		var ret int32
		return ret
	}
	return *o.VulnerabilityCount
}

// GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersThreatNodeInfo) GetVulnerabilityCountOk() (*int32, bool) {
	if o == nil || isNil(o.VulnerabilityCount) {
    return nil, false
	}
	return o.VulnerabilityCount, true
}

// HasVulnerabilityCount returns a boolean if a field has been set.
func (o *ReportersThreatNodeInfo) HasVulnerabilityCount() bool {
	if o != nil && !isNil(o.VulnerabilityCount) {
		return true
	}

	return false
}

// SetVulnerabilityCount gets a reference to the given int32 and assigns it to the VulnerabilityCount field.
func (o *ReportersThreatNodeInfo) SetVulnerabilityCount(v int32) {
	o.VulnerabilityCount = &v
}

func (o ReportersThreatNodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttackPath != nil {
		toSerialize["attack_path"] = o.AttackPath
	}
	if !isNil(o.ComplianceCount) {
		toSerialize["compliance_count"] = o.ComplianceCount
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.NodeType) {
		toSerialize["node_type"] = o.NodeType
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if !isNil(o.SecretsCount) {
		toSerialize["secrets_count"] = o.SecretsCount
	}
	if !isNil(o.VulnerabilityCount) {
		toSerialize["vulnerability_count"] = o.VulnerabilityCount
	}
	return json.Marshal(toSerialize)
}

type NullableReportersThreatNodeInfo struct {
	value *ReportersThreatNodeInfo
	isSet bool
}

func (v NullableReportersThreatNodeInfo) Get() *ReportersThreatNodeInfo {
	return v.value
}

func (v *NullableReportersThreatNodeInfo) Set(val *ReportersThreatNodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersThreatNodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersThreatNodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersThreatNodeInfo(val *ReportersThreatNodeInfo) *NullableReportersThreatNodeInfo {
	return &NullableReportersThreatNodeInfo{value: val, isSet: true}
}

func (v NullableReportersThreatNodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersThreatNodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


