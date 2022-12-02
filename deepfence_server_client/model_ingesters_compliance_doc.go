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

// IngestersComplianceDoc struct for IngestersComplianceDoc
type IngestersComplianceDoc struct {
	Timestamp *string `json:"@timestamp,omitempty"`
	ComplianceCheckType *string `json:"compliance_check_type,omitempty"`
	ComplianceNodeType *string `json:"compliance_node_type,omitempty"`
	Description *string `json:"description,omitempty"`
	DocId *string `json:"doc_id,omitempty"`
	KubernetesClusterId *string `json:"kubernetes_cluster_id,omitempty"`
	KubernetesClusterName *string `json:"kubernetes_cluster_name,omitempty"`
	Masked *string `json:"masked,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	NodeName *string `json:"node_name,omitempty"`
	NodeType *string `json:"node_type,omitempty"`
	RemediationAnsible *string `json:"remediation_ansible,omitempty"`
	RemediationPuppet *string `json:"remediation_puppet,omitempty"`
	RemediationScript *string `json:"remediation_script,omitempty"`
	Resource *string `json:"resource,omitempty"`
	ScanId *string `json:"scan_id,omitempty"`
	Status *string `json:"status,omitempty"`
	TestCategory *string `json:"test_category,omitempty"`
	TestDesc *string `json:"test_desc,omitempty"`
	TestNumber *string `json:"test_number,omitempty"`
	TestRationale *string `json:"test_rationale,omitempty"`
	TestSeverity *string `json:"test_severity,omitempty"`
	TimeStamp *int32 `json:"time_stamp,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewIngestersComplianceDoc instantiates a new IngestersComplianceDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersComplianceDoc() *IngestersComplianceDoc {
	this := IngestersComplianceDoc{}
	return &this
}

// NewIngestersComplianceDocWithDefaults instantiates a new IngestersComplianceDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersComplianceDocWithDefaults() *IngestersComplianceDoc {
	this := IngestersComplianceDoc{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetTimestamp() string {
	if o == nil || isNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetTimestampOk() (*string, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *IngestersComplianceDoc) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetComplianceCheckType returns the ComplianceCheckType field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetComplianceCheckType() string {
	if o == nil || isNil(o.ComplianceCheckType) {
		var ret string
		return ret
	}
	return *o.ComplianceCheckType
}

// GetComplianceCheckTypeOk returns a tuple with the ComplianceCheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetComplianceCheckTypeOk() (*string, bool) {
	if o == nil || isNil(o.ComplianceCheckType) {
    return nil, false
	}
	return o.ComplianceCheckType, true
}

// HasComplianceCheckType returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasComplianceCheckType() bool {
	if o != nil && !isNil(o.ComplianceCheckType) {
		return true
	}

	return false
}

// SetComplianceCheckType gets a reference to the given string and assigns it to the ComplianceCheckType field.
func (o *IngestersComplianceDoc) SetComplianceCheckType(v string) {
	o.ComplianceCheckType = &v
}

// GetComplianceNodeType returns the ComplianceNodeType field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetComplianceNodeType() string {
	if o == nil || isNil(o.ComplianceNodeType) {
		var ret string
		return ret
	}
	return *o.ComplianceNodeType
}

// GetComplianceNodeTypeOk returns a tuple with the ComplianceNodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetComplianceNodeTypeOk() (*string, bool) {
	if o == nil || isNil(o.ComplianceNodeType) {
    return nil, false
	}
	return o.ComplianceNodeType, true
}

// HasComplianceNodeType returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasComplianceNodeType() bool {
	if o != nil && !isNil(o.ComplianceNodeType) {
		return true
	}

	return false
}

// SetComplianceNodeType gets a reference to the given string and assigns it to the ComplianceNodeType field.
func (o *IngestersComplianceDoc) SetComplianceNodeType(v string) {
	o.ComplianceNodeType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IngestersComplianceDoc) SetDescription(v string) {
	o.Description = &v
}

// GetDocId returns the DocId field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetDocId() string {
	if o == nil || isNil(o.DocId) {
		var ret string
		return ret
	}
	return *o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetDocIdOk() (*string, bool) {
	if o == nil || isNil(o.DocId) {
    return nil, false
	}
	return o.DocId, true
}

// HasDocId returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasDocId() bool {
	if o != nil && !isNil(o.DocId) {
		return true
	}

	return false
}

// SetDocId gets a reference to the given string and assigns it to the DocId field.
func (o *IngestersComplianceDoc) SetDocId(v string) {
	o.DocId = &v
}

// GetKubernetesClusterId returns the KubernetesClusterId field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetKubernetesClusterId() string {
	if o == nil || isNil(o.KubernetesClusterId) {
		var ret string
		return ret
	}
	return *o.KubernetesClusterId
}

// GetKubernetesClusterIdOk returns a tuple with the KubernetesClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetKubernetesClusterIdOk() (*string, bool) {
	if o == nil || isNil(o.KubernetesClusterId) {
    return nil, false
	}
	return o.KubernetesClusterId, true
}

// HasKubernetesClusterId returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasKubernetesClusterId() bool {
	if o != nil && !isNil(o.KubernetesClusterId) {
		return true
	}

	return false
}

// SetKubernetesClusterId gets a reference to the given string and assigns it to the KubernetesClusterId field.
func (o *IngestersComplianceDoc) SetKubernetesClusterId(v string) {
	o.KubernetesClusterId = &v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetKubernetesClusterName() string {
	if o == nil || isNil(o.KubernetesClusterName) {
		var ret string
		return ret
	}
	return *o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil || isNil(o.KubernetesClusterName) {
    return nil, false
	}
	return o.KubernetesClusterName, true
}

// HasKubernetesClusterName returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasKubernetesClusterName() bool {
	if o != nil && !isNil(o.KubernetesClusterName) {
		return true
	}

	return false
}

// SetKubernetesClusterName gets a reference to the given string and assigns it to the KubernetesClusterName field.
func (o *IngestersComplianceDoc) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = &v
}

// GetMasked returns the Masked field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetMasked() string {
	if o == nil || isNil(o.Masked) {
		var ret string
		return ret
	}
	return *o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetMaskedOk() (*string, bool) {
	if o == nil || isNil(o.Masked) {
    return nil, false
	}
	return o.Masked, true
}

// HasMasked returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasMasked() bool {
	if o != nil && !isNil(o.Masked) {
		return true
	}

	return false
}

// SetMasked gets a reference to the given string and assigns it to the Masked field.
func (o *IngestersComplianceDoc) SetMasked(v string) {
	o.Masked = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetNodeId() string {
	if o == nil || isNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetNodeIdOk() (*string, bool) {
	if o == nil || isNil(o.NodeId) {
    return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasNodeId() bool {
	if o != nil && !isNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *IngestersComplianceDoc) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetNodeName() string {
	if o == nil || isNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetNodeNameOk() (*string, bool) {
	if o == nil || isNil(o.NodeName) {
    return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasNodeName() bool {
	if o != nil && !isNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *IngestersComplianceDoc) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetNodeType() string {
	if o == nil || isNil(o.NodeType) {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetNodeTypeOk() (*string, bool) {
	if o == nil || isNil(o.NodeType) {
    return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasNodeType() bool {
	if o != nil && !isNil(o.NodeType) {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *IngestersComplianceDoc) SetNodeType(v string) {
	o.NodeType = &v
}

// GetRemediationAnsible returns the RemediationAnsible field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetRemediationAnsible() string {
	if o == nil || isNil(o.RemediationAnsible) {
		var ret string
		return ret
	}
	return *o.RemediationAnsible
}

// GetRemediationAnsibleOk returns a tuple with the RemediationAnsible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetRemediationAnsibleOk() (*string, bool) {
	if o == nil || isNil(o.RemediationAnsible) {
    return nil, false
	}
	return o.RemediationAnsible, true
}

// HasRemediationAnsible returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasRemediationAnsible() bool {
	if o != nil && !isNil(o.RemediationAnsible) {
		return true
	}

	return false
}

// SetRemediationAnsible gets a reference to the given string and assigns it to the RemediationAnsible field.
func (o *IngestersComplianceDoc) SetRemediationAnsible(v string) {
	o.RemediationAnsible = &v
}

// GetRemediationPuppet returns the RemediationPuppet field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetRemediationPuppet() string {
	if o == nil || isNil(o.RemediationPuppet) {
		var ret string
		return ret
	}
	return *o.RemediationPuppet
}

// GetRemediationPuppetOk returns a tuple with the RemediationPuppet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetRemediationPuppetOk() (*string, bool) {
	if o == nil || isNil(o.RemediationPuppet) {
    return nil, false
	}
	return o.RemediationPuppet, true
}

// HasRemediationPuppet returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasRemediationPuppet() bool {
	if o != nil && !isNil(o.RemediationPuppet) {
		return true
	}

	return false
}

// SetRemediationPuppet gets a reference to the given string and assigns it to the RemediationPuppet field.
func (o *IngestersComplianceDoc) SetRemediationPuppet(v string) {
	o.RemediationPuppet = &v
}

// GetRemediationScript returns the RemediationScript field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetRemediationScript() string {
	if o == nil || isNil(o.RemediationScript) {
		var ret string
		return ret
	}
	return *o.RemediationScript
}

// GetRemediationScriptOk returns a tuple with the RemediationScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetRemediationScriptOk() (*string, bool) {
	if o == nil || isNil(o.RemediationScript) {
    return nil, false
	}
	return o.RemediationScript, true
}

// HasRemediationScript returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasRemediationScript() bool {
	if o != nil && !isNil(o.RemediationScript) {
		return true
	}

	return false
}

// SetRemediationScript gets a reference to the given string and assigns it to the RemediationScript field.
func (o *IngestersComplianceDoc) SetRemediationScript(v string) {
	o.RemediationScript = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetResource() string {
	if o == nil || isNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetResourceOk() (*string, bool) {
	if o == nil || isNil(o.Resource) {
    return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasResource() bool {
	if o != nil && !isNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *IngestersComplianceDoc) SetResource(v string) {
	o.Resource = &v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetScanId() string {
	if o == nil || isNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetScanIdOk() (*string, bool) {
	if o == nil || isNil(o.ScanId) {
    return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasScanId() bool {
	if o != nil && !isNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *IngestersComplianceDoc) SetScanId(v string) {
	o.ScanId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IngestersComplianceDoc) SetStatus(v string) {
	o.Status = &v
}

// GetTestCategory returns the TestCategory field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetTestCategory() string {
	if o == nil || isNil(o.TestCategory) {
		var ret string
		return ret
	}
	return *o.TestCategory
}

// GetTestCategoryOk returns a tuple with the TestCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetTestCategoryOk() (*string, bool) {
	if o == nil || isNil(o.TestCategory) {
    return nil, false
	}
	return o.TestCategory, true
}

// HasTestCategory returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasTestCategory() bool {
	if o != nil && !isNil(o.TestCategory) {
		return true
	}

	return false
}

// SetTestCategory gets a reference to the given string and assigns it to the TestCategory field.
func (o *IngestersComplianceDoc) SetTestCategory(v string) {
	o.TestCategory = &v
}

// GetTestDesc returns the TestDesc field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetTestDesc() string {
	if o == nil || isNil(o.TestDesc) {
		var ret string
		return ret
	}
	return *o.TestDesc
}

// GetTestDescOk returns a tuple with the TestDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetTestDescOk() (*string, bool) {
	if o == nil || isNil(o.TestDesc) {
    return nil, false
	}
	return o.TestDesc, true
}

// HasTestDesc returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasTestDesc() bool {
	if o != nil && !isNil(o.TestDesc) {
		return true
	}

	return false
}

// SetTestDesc gets a reference to the given string and assigns it to the TestDesc field.
func (o *IngestersComplianceDoc) SetTestDesc(v string) {
	o.TestDesc = &v
}

// GetTestNumber returns the TestNumber field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetTestNumber() string {
	if o == nil || isNil(o.TestNumber) {
		var ret string
		return ret
	}
	return *o.TestNumber
}

// GetTestNumberOk returns a tuple with the TestNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetTestNumberOk() (*string, bool) {
	if o == nil || isNil(o.TestNumber) {
    return nil, false
	}
	return o.TestNumber, true
}

// HasTestNumber returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasTestNumber() bool {
	if o != nil && !isNil(o.TestNumber) {
		return true
	}

	return false
}

// SetTestNumber gets a reference to the given string and assigns it to the TestNumber field.
func (o *IngestersComplianceDoc) SetTestNumber(v string) {
	o.TestNumber = &v
}

// GetTestRationale returns the TestRationale field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetTestRationale() string {
	if o == nil || isNil(o.TestRationale) {
		var ret string
		return ret
	}
	return *o.TestRationale
}

// GetTestRationaleOk returns a tuple with the TestRationale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetTestRationaleOk() (*string, bool) {
	if o == nil || isNil(o.TestRationale) {
    return nil, false
	}
	return o.TestRationale, true
}

// HasTestRationale returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasTestRationale() bool {
	if o != nil && !isNil(o.TestRationale) {
		return true
	}

	return false
}

// SetTestRationale gets a reference to the given string and assigns it to the TestRationale field.
func (o *IngestersComplianceDoc) SetTestRationale(v string) {
	o.TestRationale = &v
}

// GetTestSeverity returns the TestSeverity field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetTestSeverity() string {
	if o == nil || isNil(o.TestSeverity) {
		var ret string
		return ret
	}
	return *o.TestSeverity
}

// GetTestSeverityOk returns a tuple with the TestSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetTestSeverityOk() (*string, bool) {
	if o == nil || isNil(o.TestSeverity) {
    return nil, false
	}
	return o.TestSeverity, true
}

// HasTestSeverity returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasTestSeverity() bool {
	if o != nil && !isNil(o.TestSeverity) {
		return true
	}

	return false
}

// SetTestSeverity gets a reference to the given string and assigns it to the TestSeverity field.
func (o *IngestersComplianceDoc) SetTestSeverity(v string) {
	o.TestSeverity = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetTimeStamp() int32 {
	if o == nil || isNil(o.TimeStamp) {
		var ret int32
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetTimeStampOk() (*int32, bool) {
	if o == nil || isNil(o.TimeStamp) {
    return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasTimeStamp() bool {
	if o != nil && !isNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given int32 and assigns it to the TimeStamp field.
func (o *IngestersComplianceDoc) SetTimeStamp(v int32) {
	o.TimeStamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IngestersComplianceDoc) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersComplianceDoc) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IngestersComplianceDoc) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IngestersComplianceDoc) SetType(v string) {
	o.Type = &v
}

func (o IngestersComplianceDoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timestamp) {
		toSerialize["@timestamp"] = o.Timestamp
	}
	if !isNil(o.ComplianceCheckType) {
		toSerialize["compliance_check_type"] = o.ComplianceCheckType
	}
	if !isNil(o.ComplianceNodeType) {
		toSerialize["compliance_node_type"] = o.ComplianceNodeType
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DocId) {
		toSerialize["doc_id"] = o.DocId
	}
	if !isNil(o.KubernetesClusterId) {
		toSerialize["kubernetes_cluster_id"] = o.KubernetesClusterId
	}
	if !isNil(o.KubernetesClusterName) {
		toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	}
	if !isNil(o.Masked) {
		toSerialize["masked"] = o.Masked
	}
	if !isNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !isNil(o.NodeName) {
		toSerialize["node_name"] = o.NodeName
	}
	if !isNil(o.NodeType) {
		toSerialize["node_type"] = o.NodeType
	}
	if !isNil(o.RemediationAnsible) {
		toSerialize["remediation_ansible"] = o.RemediationAnsible
	}
	if !isNil(o.RemediationPuppet) {
		toSerialize["remediation_puppet"] = o.RemediationPuppet
	}
	if !isNil(o.RemediationScript) {
		toSerialize["remediation_script"] = o.RemediationScript
	}
	if !isNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !isNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.TestCategory) {
		toSerialize["test_category"] = o.TestCategory
	}
	if !isNil(o.TestDesc) {
		toSerialize["test_desc"] = o.TestDesc
	}
	if !isNil(o.TestNumber) {
		toSerialize["test_number"] = o.TestNumber
	}
	if !isNil(o.TestRationale) {
		toSerialize["test_rationale"] = o.TestRationale
	}
	if !isNil(o.TestSeverity) {
		toSerialize["test_severity"] = o.TestSeverity
	}
	if !isNil(o.TimeStamp) {
		toSerialize["time_stamp"] = o.TimeStamp
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableIngestersComplianceDoc struct {
	value *IngestersComplianceDoc
	isSet bool
}

func (v NullableIngestersComplianceDoc) Get() *IngestersComplianceDoc {
	return v.value
}

func (v *NullableIngestersComplianceDoc) Set(val *IngestersComplianceDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersComplianceDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersComplianceDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersComplianceDoc(val *IngestersComplianceDoc) *NullableIngestersComplianceDoc {
	return &NullableIngestersComplianceDoc{value: val, isSet: true}
}

func (v NullableIngestersComplianceDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersComplianceDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


